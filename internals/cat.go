package internals

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/sarvsav/iza/database"
	"github.com/sarvsav/iza/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OptionsCatFunc func(c *models.CatOptions) error

// Cat is equivalent to the cat command in Unix-like systems.
// It is used to display the contents of a document in the collection.
func Cat(catOptions ...OptionsCatFunc) error {

	catCmd := &models.CatOptions{
		Args:   []string{},
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	var dbName, collectionName string

	for _, opt := range catOptions {
		if err := opt(catCmd); err != nil {
			return err
		}
	}

	catCmd.Logger.Debug("provided command with options", "args", catCmd.Args)

	// Create a new client and connect to the server
	client, err := database.GetMongoClient()
	defer func() {
		if err := database.DisconnectMongoClient(client); err != nil {
			catCmd.Logger.Error("Failed to disconnect from MongoDB", "error", err)
		}
	}()

	if err != nil {
		catCmd.Logger.Error("Failed to connect to MongoDB", "error", err)
		return err
	}

	// Iterate the arguments and create a collection for each
	for _, arg := range catCmd.Args {
		// Extract db and collection names from the argument
		argParts := strings.Split(arg, "/")
		if len(argParts) > 2 {
			catCmd.Logger.Error("Expected format is database/collection", "received", arg)
		}
		if len(argParts) == 1 {
			catCmd.Logger.Info("No database provided, reading from test", "received", arg)
			dbName = "test"
			collectionName = argParts[0]
		}
		if len(argParts) == 2 {
			catCmd.Logger.Debug("Reading from collection", "received", arg)
			dbName = argParts[0]
			collectionName = argParts[1]
		}

		coll := client.Database(dbName).Collection(collectionName)
		opts := options.Count().SetHint("_id_")
		count, err := coll.CountDocuments(context.TODO(), bson.D{}, opts)
		if err != nil {
			catCmd.Logger.Error("Failed to count documents", "error", err)
		}
		catCmd.Logger.Info("Total documents in collection", "dbName", dbName, "collection", collectionName, "count", count)

		var results []bson.M

		filter := bson.D{{}}

		// Retrieves documents that match the query filter
		cursor, err := coll.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}

		if err := cursor.All(context.TODO(), &results); err != nil {
			log.Panic(err)
		}
		fmt.Println(results)
	}

	return nil
}
