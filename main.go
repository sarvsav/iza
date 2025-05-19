/*
Copyright © 2024 sarvsav

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	_ "embed"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"github.com/joho/godotenv"
	"github.com/sarvsav/iza/artifactory"
	"github.com/sarvsav/iza/client"
	"github.com/sarvsav/iza/cmd"
	"github.com/sarvsav/iza/dbstore"
	"github.com/sarvsav/iza/devops"
	"github.com/sarvsav/iza/foundation/logger"
	"github.com/sarvsav/iza/internals/app"
	"github.com/sarvsav/iza/internals/artifactstore"
	"github.com/sarvsav/iza/internals/cicd"
	"github.com/sarvsav/iza/internals/datastore"
	"go.mongodb.org/mongo-driver/mongo"
)

// for cue configuration
var requiredKeys = []string{"database", "artifactory", "ci-tools"}

//go:embed cue/dev/schema.cue
var DevSchema string

//go:embed cue/prod/schema.cue
var ProdSchema string

type NamedMongoClient struct {
	Name   string
	Client *mongo.Client
}

type NamedPostgresClient struct {
	Name   string
	Client *sql.DB
}

type NamedHTTPClient struct {
	Name   string
	Client *http.Client
}

type ClientRegistry struct {
	Mongo    []NamedMongoClient
	Postgres []NamedPostgresClient
	JFrog    []NamedHTTPClient
	Jenkins  []NamedHTTPClient
}

func main() {

	// -------------------------------------------------------------------------
	// Setup logger
	var log *logger.Logger
	clientRegistry := ClientRegistry{}

	// Warn and Error are custom events and called when the log calls Warn or Error
	// For example, if log.Warn is called, then the function defined in Warn will be called
	// This is useful for sending alerts or warnings to external systems
	events := logger.Events{
		Warn:  func(ctx context.Context, r logger.Record) { log.Info(ctx, "******* Warning ******") },
		Error: func(ctx context.Context, r logger.Record) { log.Info(ctx, "******* SEND PAGER ******") },
	}

	traceIDFn := func(ctx context.Context) string {
		return "00000000-0000-0000-0000-000000000000"
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelDebug, "IZA", traceIDFn, events)

	// -------------------------------------------------------------------------
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Error(context.Background(), "Error loading .env file", "error", err)
	}

	// -------------------------------------------------------------------------
	// Load CUE configuration
	ctx := cuecontext.New()
	instances := load.Instances([]string{"./cue/dev"}, nil)

	if len(instances) == 0 {
		log.Error(context.Background(), "no cue instances found", "location", "./cue/dev")
	}

	instance := instances[0] // If there are multiple files then need to iterate

	if instance.Err != nil {
		log.Error(context.Background(), "Error building cue value", "error", instance.Err)
	}

	value := ctx.BuildInstance(instance)
	if value.Err() != nil {
		log.Error(context.Background(), "Error building cue value", "error", value.Err())
	}
	// Extract top-level "config" field
	config := value.LookupPath(cue.ParsePath("config"))
	if !config.Exists() {
		log.Error(context.Background(), "No 'config' field found in CUE file")
	}

	// Validate required keys
	for _, key := range requiredKeys {
		if !config.LookupPath(cue.ParsePath(key)).Exists() {
			log.Info(context.Background(), "missing in the cue config", "section", key)
		} else {
			log.Debug(context.Background(), "section: "+key+" found")
		}
	}

	// Extract and print database entries
	databases := config.LookupPath(cue.ParsePath("database"))
	if databases.Exists() {
		iter, _ := databases.Fields()
		for iter.Next() {
			selector := iter.Selector()
			log.Debug(context.Background(), "database entry", "selector", selector.String(), "value", iter.Value())
			kindVal := iter.Value().LookupPath(cue.ParsePath("type"))
			if kindVal.Exists() {
				kind, _ := kindVal.String()
				switch kind {
				case "mongodb":
					clientName := selector.String()
					// Get mongo client
					mc, err := client.GetMongoClient()
					defer func() {
						if err := client.DisconnectMongoClient(mc); err != nil {
							log.Error(context.Background(), "Failed to disconnect from MongoDB", "error", err)
						}
					}()

					if err != nil {
						log.Error(context.Background(), "Failed to connect to MongoDB", "error", err)
						return
					}
					clientRegistry.Mongo = append(clientRegistry.Mongo, NamedMongoClient{Name: clientName, Client: mc})
				case "postgres":
					clientName := selector.String()
					clientRegistry.Postgres = append(clientRegistry.Postgres, NamedPostgresClient{Name: clientName, Client: nil})
				default:
					log.Debug(context.Background(), "Unknown database type", "type", kind)
				}
			}
		}
	}

	// Extract and print Artifactory entries
	artifactoryTools := config.LookupPath(cue.ParsePath("artifactory"))
	if artifactoryTools.Exists() {
		iter, _ := artifactoryTools.Fields()
		for iter.Next() {
			selector := iter.Selector()
			log.Debug(context.Background(), "artifactory entry", "selector", selector.String(), "value", iter.Value())
			kindVal := iter.Value().LookupPath(cue.ParsePath("type"))
			if kindVal.Exists() {
				kind, _ := kindVal.String()
				switch kind {
				case "jfrog":
					clientName := selector.String()
					// Get jfrog client
					jc, _ := client.GetJFrogClient()
					clientRegistry.JFrog = append(clientRegistry.JFrog, NamedHTTPClient{Name: clientName, Client: jc})
				default:
					log.Debug(context.Background(), "Unknown artifactory type", "type", kind)
				}
			}
		}
	}

	// Extract and print CI/CD tool entries
	ciTools := config.LookupPath(cue.ParsePath("ci_tools"))
	if ciTools.Exists() {
		iter, _ := ciTools.Fields()
		for iter.Next() {
			selector := iter.Selector()
			log.Debug(context.Background(), "ci_tools entry", "selector", selector.String(), "value", iter.Value())
			kindVal := iter.Value().LookupPath(cue.ParsePath("type"))
			if kindVal.Exists() {
				kind, _ := kindVal.String()
				switch kind {
				case "jenkins":
					clientName := selector.String()
					// Get jenkins client
					jc, _ := client.GetJenkinsClient()
					clientRegistry.Jenkins = append(clientRegistry.Jenkins, NamedHTTPClient{Name: clientName, Client: jc})
				case "gh-actions":
					// Not implemented yet
				default:
					log.Debug(context.Background(), "Unknown cicd type", "type", kind)
				}
			}
		}
	}

	// Extract the full JSON configuration for debugging
	// jsonData, err := config.MarshalJSON()
	// if err != nil {
	// 	log.Error(context.Background(), "Error marshalling CUE to JSON", "error", err)
	// }

	// log.Debug(context.Background(), "CUE configuration", "json", string(jsonData))

	// Validate the CUE data (optional but recommended).
	if err := value.Validate(); err != nil {
		log.Error(context.Background(), "CUE validation error", "error", err)
	}

	// -------------------------------------------------------------------------
	// Setup application
	jFrogClient := artifactory.NewJFrogClient(clientRegistry.JFrog[0].Client, log)
	jenkinsClient := devops.NewJenkinsClient(clientRegistry.Jenkins[0].Client, log)
	mongoClient := dbstore.NewMongoClient(clientRegistry.Mongo[0].Client, log)
	app := &app.Application{
		ArtifactoryService: artifactstore.NewArtifactoryService(jFrogClient, log),
		CiCdService:        cicd.NewCiCdService(jenkinsClient, log),
		DataStoreService:   datastore.NewDataStoreService(mongoClient, log),
		Logger:             log,
	}

	// -------------------------------------------------------------------------
	// Execute the command
	cmd.Execute(app)
}
