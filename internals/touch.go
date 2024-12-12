package internals

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"strings"

	"github.com/sarvsav/iza/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OptionsTouchFunc func(c *models.TouchOptions) error

// Touch is equivalent to the touch command in Unix-like systems.
// It is used to create an empty collection in the database.
func Touch(touchOptions ...OptionsTouchFunc) error {
	mongodbURI := os.Getenv("MONGODB_URI")
	touchCmd := &models.TouchOptions{
		Args:   []string{},
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	for _, opt := range touchOptions {
		if err := opt(touchCmd); err != nil {
			return err
		}
	}

	touchCmd.Logger.Debug("provided command with options", "args", touchCmd.Args)
	if len(touchCmd.Args) == 0 {
		touchCmd.Logger.Error("Expected format is database/collection", "received", "empty")
		return errors.New("expected format: iza touch database/collection")
	}

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodbURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var dbName, collectionName string

	// Iterate the arguments and create a collection for each
	for _, arg := range touchCmd.Args {
		// Extract dbname and collection name from the argument
		argParts := strings.Split(arg, "/")
		if len(argParts) > 2 {
			touchCmd.Logger.Error("Expected format is database/collection", "received", arg)
		}
		if len(argParts) == 1 {
			touchCmd.Logger.Info("No database provided, creating inside test", "received", arg)
			dbName = "test"
			collectionName = argParts[0]
		}
		if len(argParts) == 2 {
			touchCmd.Logger.Debug("Creating empty collection", "received", arg)
			dbName = argParts[0]
			collectionName = argParts[1]
		}

		if err := client.Database(dbName).CreateCollection(context.TODO(), collectionName); err != nil {
			touchCmd.Logger.Error("Failed to create collection", "error", err)
		}
		touchCmd.Logger.Info("Successfully created empty collection", "dbname", dbName, "collection", collectionName)
	}
	return nil
}
