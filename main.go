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
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"github.com/joho/godotenv"
	"github.com/sarvsav/iza/cmd"
	"github.com/sarvsav/iza/database"
	"github.com/sarvsav/iza/dbstore"
	"github.com/sarvsav/iza/devops"
	"github.com/sarvsav/iza/foundation/logger"
	"github.com/sarvsav/iza/internals/app"
	"github.com/sarvsav/iza/internals/cicd"
	"github.com/sarvsav/iza/internals/datastore"
)

// for cue configuration
var requiredKeys = []string{"database", "artifactory", "ci-tools"}

func main() {

	// -------------------------------------------------------------------------
	// Setup logger
	var log *logger.Logger

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

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "IZA", traceIDFn, events)

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
		}
	}

	// Extract and print Artifactory entries
	artifactory := config.LookupPath(cue.ParsePath("artifactory"))
	if artifactory.Exists() {
		iter, _ := artifactory.Fields()
		for iter.Next() {
			selector := iter.Selector()
			log.Debug(context.Background(), "artifactory entry", "selector", selector.String(), "value", iter.Value())
		}
	}

	// Extract and print CI/CD tool entries
	ciTools := config.LookupPath(cue.ParsePath("ci_tools"))
	if ciTools.Exists() {
		iter, _ := ciTools.Fields()
		for iter.Next() {
			selector := iter.Selector()
			log.Debug(context.Background(), "ci_tools entry", "selector", selector.String(), "value", iter.Value())
		}
	}

	// Extract the full JSON configuration for debugging
	jsonData, err := config.MarshalJSON()
	if err != nil {
		log.Error(context.Background(), "Error marshalling CUE to JSON", "error", err)
	}

	log.Debug(context.Background(), "CUE configuration", "json", string(jsonData))

	// Validate the CUE data (optional but recommended).
	if err := value.Validate(); err != nil {
		log.Error(context.Background(), "CUE validation error", "error", err)
	}

	// -------------------------------------------------------------------------
	// Get mongo client
	client, err := database.GetMongoClient()
	defer func() {
		if err := database.DisconnectMongoClient(client); err != nil {
			log.Error(context.Background(), "Failed to disconnect from MongoDB", "error", err)
		}
	}()

	// -------------------------------------------------------------------------
	// Setup application
	jenkinsClient := devops.NewJenkinsClient("user", "some-api-token")
	mongoClient := dbstore.NewMongoClient(client, log)
	app := &app.Application{
		CiCdService:      cicd.NewCiCdService(jenkinsClient, log),
		DataStoreService: datastore.NewDataStoreService(mongoClient, log),
		Logger:           log,
	}

	// -------------------------------------------------------------------------
	// Execute the command
	cmd.Execute(app)
}
