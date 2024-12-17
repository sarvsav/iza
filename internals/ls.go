package internals

import (
	"context"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/sarvsav/iza/database"
	"github.com/sarvsav/iza/models"
	"go.mongodb.org/mongo-driver/bson"
)

type OptionsLsFunc func(c *models.LsOptions) error

func Ls(lsOptions ...OptionsLsFunc) error {
	lsCmd := &models.LsOptions{
		LongListing: false,
		Color:       false,
		Args:        []string{},
		Logger:      slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	for _, opt := range lsOptions {
		if err := opt(lsCmd); err != nil {
			return err
		}
	}

	lsCmd.Logger.Debug("provided command with options", "longListing", lsCmd.LongListing, "color", lsCmd.Color, "args", lsCmd.Args)
	client, err := database.GetMongoClient()
	defer func() {
		if err := database.DisconnectMongoClient(client); err != nil {
			lsCmd.Logger.Error("Failed to disconnect from MongoDB", "error", err)
		}
	}()

	if err != nil {
		lsCmd.Logger.Error("Failed to connect to MongoDB", "error", err)
		return err
	}

	if len(lsCmd.Args) == 0 {
		// Use a filter to only select non-empty databases.
		dbList, err := client.ListDatabaseNames(
			context.TODO(),
			bson.D{},
		)
		if err != nil {
			log.Panic(err)
		}
		lsCmd.Logger.Info("List of databases", "db", dbList)
	}

	for _, arg := range lsCmd.Args {
		// Extract db and collection names from the argument
		argParts := strings.Split(arg, "/")
		if len(argParts) > 2 {
			lsCmd.Logger.Error("Expected format is database/collection", "received", arg)
		}
		if len(argParts) == 1 {
			dbName := argParts[0]
			database := client.Database(dbName)
			collections, err := database.ListCollectionNames(context.TODO(), bson.D{})
			if err != nil {
				lsCmd.Logger.Error("Failed to list collections", "error", err)
			}
			lsCmd.Logger.Info("List of collections", "db", dbName, "collections", collections)
		}
		if len(argParts) == 2 {
			dbName := argParts[0]
			collectionName := argParts[1]
			collection := client.Database(dbName).Collection(collectionName)
			collectionIndexes, err := collection.Indexes().List(context.TODO())
			if err != nil {
				lsCmd.Logger.Error("Failed to list collection info", "error", err)
			}
			lsCmd.Logger.Info("Collection info", "db", dbName, "collection", collectionName, "indexes", collectionIndexes)
		}
	}
	return nil
}
