package internals

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/sarvsav/iza/database"
	"github.com/sarvsav/iza/models"
	"go.mongodb.org/mongo-driver/bson"
)

type OptionsDuFunc func(c *models.DuOptions) error

// Du is equivalent to the du command in Unix-like systems.
// It is used to calculate the disk usage of database or collection.
func Du(duOptions ...OptionsDuFunc) error {

	var dbName, collectionName string

	duCmd := &models.DuOptions{
		Args:   []string{},
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	for _, opt := range duOptions {
		if err := opt(duCmd); err != nil {
			return err
		}
	}

	duCmd.Logger.Debug("provided command with options", "args", duCmd.Args)

	client, err := database.GetMongoClient()
	defer func() {
		if err := database.DisconnectMongoClient(client); err != nil {
			duCmd.Logger.Error("Failed to disconnect from MongoDB", "error", err)
		}
	}()

	if err != nil {
		duCmd.Logger.Error("Failed to connect to MongoDB", "error", err)
		return err
	}

	// Find db and collection name
	for _, arg := range duCmd.Args {
		// Extract db and collection names from the argument
		argParts := strings.Split(arg, "/")
		if len(argParts) > 2 {
			duCmd.Logger.Error("Expected format is database/collection", "received", arg)
		}
		if len(argParts) == 1 {
			duCmd.Logger.Debug("No collection name provided", "received", arg)
			dbName = argParts[0]
			collectionName = ""
		}
		if len(argParts) == 2 {
			duCmd.Logger.Debug("Calculating collection size inside db", "received", arg)
			dbName = argParts[0]
			collectionName = argParts[1]
		}
		if collectionName != "" {
			stats := bson.M{}
			err := client.Database(dbName).RunCommand(context.TODO(), bson.D{{Key: "collStats", Value: collectionName}}).Decode(&stats)
			if err != nil {
				duCmd.Logger.Error("Failed to get collection stats", collectionName, err)
			}
			duCmd.Logger.Info("Collection size in bytes", collectionName, stats["size"])
		} else {
			stats := bson.M{}
			err := client.Database(dbName).RunCommand(context.TODO(), bson.D{{Key: "dbStats", Value: 1}}).Decode(&stats)
			if err != nil {
				duCmd.Logger.Error("Failed to get database stats", dbName, err)
			}
			duCmd.Logger.Info("Database size: in bytes", dbName, stats["dataSize"])
		}
	}
	return nil
}
