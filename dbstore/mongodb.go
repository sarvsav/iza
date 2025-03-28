package dbstore

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/sarvsav/iza/database"
	"github.com/sarvsav/iza/foundation/logger"
	"github.com/sarvsav/iza/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	mc  *mongo.Client
	log *logger.Logger
}

func NewMongoClient(mc *mongo.Client, log *logger.Logger) *mongoClient {
	return &mongoClient{
		mc:  mc,
		log: log,
	}
}

// WhoAmI is equivalent to the whoami command.
// It prints the current logged in user.
func (m *mongoClient) WhoAmI(whoAmIOptions ...OptionsWhoAmIFunc) (models.WhoAmIResponse, error) {

	var username string

	whoAmICmd := &models.WhoAmIOptions{
		Args: []string{},
	}

	for _, opt := range whoAmIOptions {
		if err := opt(whoAmICmd); err != nil {
			return models.WhoAmIResponse{}, err
		}
	}

	m.log.Debug(context.Background(), "provided options", "args", whoAmICmd.Args)

	info := bson.M{}
	if err := m.mc.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "connectionStatus", Value: 1}}).Decode(&info); err != nil {
		m.log.Error(context.Background(), "Failed to get connection status", "error", err)
		return models.WhoAmIResponse{}, err
	}

	// Accessing "authenticatedUsers"
	authInfo, ok := info["authInfo"].(bson.M)
	if !ok {
		m.log.Error(context.Background(), "authInfo is not of type bson.M")
		return models.WhoAmIResponse{}, errors.New("authInfo is not of type bson.M")
	}

	authenticatedUsers, ok := authInfo["authenticatedUsers"].(primitive.A)
	if !ok {
		m.log.Error(context.Background(), "authenticatedUsers is not of type []interface{}", "authenticatedUsers", authInfo["authenticatedUsers"])
		return models.WhoAmIResponse{}, errors.New("authenticatedUsers is not of type []interface{}")
	}

	// Accessing user details
	if len(authenticatedUsers) > 0 {
		user, ok := authenticatedUsers[0].(bson.M)
		if !ok {
			m.log.Error(context.Background(), "first element in authenticatedUsers is not of type bson.M")
			return models.WhoAmIResponse{}, errors.New("first element in authenticatedUsers is not of type bson.M")
		}

		// Extract the "user" field
		username, ok = user["user"].(string)
		if ok {
			m.log.Debug(context.Background(), "Authenticated", "username", username)
		} else {
			m.log.Error(context.Background(), "User field is not a string or does not exist")
		}
	} else {
		m.log.Error(context.Background(), "authenticatedUsers is empty")
	}

	return models.WhoAmIResponse{
		Username: username,
	}, nil
}

func (m mongoClient) Ls(lsOptions ...OptionsLsFunc) ([]string, error) {
	lsCmd := &models.LsOptions{
		LongListing: false,
		Color:       false,
		Args:        []string{},
		Logger:      slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	for _, opt := range lsOptions {
		if err := opt(lsCmd); err != nil {
			return nil, err
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
		return nil, err
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
	return nil, nil
}

// Du is equivalent to the du command in Unix-like systems.
// It is used to calculate the disk usage of database or collection.
func (m mongoClient) Du(duOptions ...OptionsDuFunc) error {

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

// Touch is equivalent to the touch command in Unix-like systems.
// It is used to create an empty collection in the database.
func (m mongoClient) Touch(touchOptions ...OptionsTouchFunc) (string, error) {

	var dbName, collectionName string

	touchCmd := &models.TouchOptions{
		Args:   []string{},
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	for _, opt := range touchOptions {
		if err := opt(touchCmd); err != nil {
			return "", err
		}
	}

	touchCmd.Logger.Debug("provided command with options", "args", touchCmd.Args)
	if len(touchCmd.Args) == 0 {
		touchCmd.Logger.Error("Expected format is database/collection", "received", "empty")
		return "", errors.New("expected format: iza touch database/collection")
	}

	client, err := database.GetMongoClient()
	defer func() {
		if err := database.DisconnectMongoClient(client); err != nil {
			touchCmd.Logger.Error("Failed to disconnect from MongoDB", "error", err)
		}
	}()

	if err != nil {
		touchCmd.Logger.Error("Failed to connect to MongoDB", "error", err)
		return "", err
	}

	// Iterate the arguments and create a collection for each
	for _, arg := range touchCmd.Args {
		// Extract db and collection names from the argument
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
		touchCmd.Logger.Info("Successfully created empty collection", "dbName", dbName, "collection", collectionName)
	}

	return collectionName, nil
}

// Cat is equivalent to the cat command in Unix-like systems.
// It is used to display the contents of a document in the collection.
func (m mongoClient) Cat(catOptions ...OptionsCatFunc) error {

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
