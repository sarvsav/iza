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
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"github.com/joho/godotenv"
	"github.com/sarvsav/iza/cmd"
	"github.com/sarvsav/iza/dbstore"
	"github.com/sarvsav/iza/devops"
	"github.com/sarvsav/iza/foundation/logger"
	"github.com/sarvsav/iza/internals/app"
	"github.com/sarvsav/iza/internals/cicd"
	"github.com/sarvsav/iza/internals/datastore"
)

func main() {

	// -------------------------------------------------------------------------
	// Setup logger
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) { log.Info(ctx, "******* SEND ALERT ******") },
	}

	traceIDFn := func(ctx context.Context) string {
		return "00000000-0000-0000-0000-000000000000"
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "IZA", traceIDFn, events)

	// -------------------------------------------------------------------------
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Error(context.Background(), "Error loading .env file")
	}

	// -------------------------------------------------------------------------
	// Load CUE configuration
	ctx := cuecontext.New()
	instances := load.Instances([]string{"./cue/dev"}, nil)

	log.Error(context.Background(), "instances", instances)

	if len(instances) == 0 {
		log.Error(context.Background(), "no instances found")
	}

	instance := instances[0] // If there are multiple files then need to iterate

	if instance.Err != nil {
		log.Error(context.Background(), "Error building cue value: %v", instance.Err)
	}

	value := ctx.BuildInstance(instance)
	if value.Err() != nil {
		log.Error(context.Background(), "Error building CUE value: %v", value.Err())
	}
	fmt.Printf("%v\n", value)

	port := value.LookupPath(cue.ParsePath("port"))
	databaseName := value.LookupPath(cue.ParsePath("database.name"))
	databaseHost := value.LookupPath(cue.ParsePath("database.host"))

	if port.Err() != nil {
		log.Error(context.Background(), "Error reading port: %v", port.Err())
	}

	if databaseName.Err() != nil {
		log.Error(context.Background(), "Error reading database name: %v", databaseName.Err())
	}

	if databaseHost.Err() != nil {
		log.Error(context.Background(), "Error reading database host: %v", databaseHost.Err())
	}

	portInt, _ := port.Int64()
	databaseNameStr, _ := databaseName.String()
	databaseHostStr, _ := databaseHost.String()

	fmt.Printf("Port: %d\n", portInt)
	fmt.Printf("Database Name: %s\n", databaseNameStr)
	fmt.Printf("Database Host: %s\n", databaseHostStr)

	// 4. Validate the CUE data (optional but recommended).
	if err := value.Validate(); err != nil {
		log.Error(context.Background(), "CUE validation error: %v", err)
	}

	// -------------------------------------------------------------------------
	// Setup application
	jenkinsClient := devops.NewJenkinsClient("user", "some-api-token")
	mongoClient := dbstore.NewMongoClient("user", "some-api-token")
	app := &app.Application{
		CiCdService:      cicd.NewCiCdService(jenkinsClient, log),
		DataStoreService: datastore.NewDataStoreService(mongoClient, log),
		Logger:           log,
	}

	// -------------------------------------------------------------------------
	// Execute the command
	cmd.Execute(app)
}
