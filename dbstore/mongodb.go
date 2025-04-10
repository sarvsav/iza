package dbstore

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

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
func (m *mongoClient) WhoAmI(whoAmIOptions ...models.OptionsWhoAmIFunc) (models.WhoAmIResponse, error) {

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

func (m mongoClient) Ls(lsOptions ...models.OptionsLsFunc) (models.LsResponse, error) {
	var result models.LsResponse
	lsCmd := &models.LsOptions{
		LongListing: false,
		Color:       false,
		Args:        []string{},
	}

	for _, opt := range lsOptions {
		if err := opt(lsCmd); err != nil {
			return models.LsResponse{}, err
		}
	}

	m.log.Debug(context.Background(), "provided command with options",
		"longListing", lsCmd.LongListing,
		"color", lsCmd.Color,
		"args", lsCmd.Args)

	if len(lsCmd.Args) == 0 {
		// Use a filter to only select non-empty databases.
		dbList, err := m.mc.ListDatabaseNames(
			context.TODO(),
			bson.D{},
		)
		if err != nil {
			log.Panic(err)
		}
		m.log.Debug(context.Background(), "List of databases", "db", dbList)
		var database models.Database
		for _, dbName := range dbList {
			database.Name = dbName
			database.Perms = "rw-rw-rw-"       // Placeholder, as MongoDB does not provide permissions in ListDatabaseNames
			database.Owner = "root"            // Placeholder, as MongoDB does not provide owner/group in ListDatabaseNames
			database.Group = "root"            // Placeholder, as MongoDB does not provide owner/group in ListDatabaseNames
			database.Size = 0                  // Placeholder, as MongoDB does not provide size in ListDatabaseNames
			database.LastModified = time.Now() // Placeholder, as MongoDB does not provide last modified in ListDatabaseNames
			result.Databases = append(result.Databases, database)
		}
	}

	for _, arg := range lsCmd.Args {
		// Extract db and collection names from the argument
		argParts := strings.Split(arg, "/")
		if len(argParts) > 2 {
			m.log.Error(context.Background(), "Expected format is database/collection", "received", arg)
		}
		if len(argParts) == 1 {
			dbName := argParts[0]
			database := m.mc.Database(dbName)
			collections, err := database.ListCollectionNames(context.TODO(), bson.D{})
			if err != nil {
				m.log.Error(context.Background(), "Failed to list collections", "error", err)
			}
			m.log.Debug(context.Background(), "List of collections", "db", dbName, "collections", collections)
			result.Databases = []models.Database{
				{
					Name:         dbName,
					Perms:        "rw-rw-rw-", // Placeholder, as MongoDB does not provide permissions in ListDatabaseNames
					Owner:        "root",      // Placeholder, as MongoDB does not provide owner/group in ListDatabaseNames
					Group:        "root",      // Placeholder, as MongoDB does not provide owner/group in ListDatabaseNames
					Size:         0,           // Placeholder, as MongoDB does not provide size in ListDatabaseNames
					LastModified: time.Now(),  // Placeholder, as MongoDB does not provide last modified in ListDatabaseNames
				},
			}
			var collection models.Collection
			for _, collectionName := range collections {
				collection.Name = collectionName
				collection.Perms = "rw-rw-rw-"       // Placeholder, as MongoDB does not provide permissions in ListCollectionNames
				collection.Owner = "root"            // Placeholder, as MongoDB does not provide owner/group in ListCollectionNames
				collection.Group = "root"            // Placeholder, as MongoDB does not provide owner/group in ListCollectionNames
				collection.Size = 0                  // Placeholder, as MongoDB does not provide size in ListCollectionNames
				collection.LastModified = time.Now() // Placeholder, as MongoDB does not provide last modified in ListCollectionNames
				result.Collections = append(result.Collections, collection)
			}
		}
		if len(argParts) == 2 {
			dbName := argParts[0]
			collectionName := argParts[1]
			collection := m.mc.Database(dbName).Collection(collectionName)
			collectionIndexes, err := collection.Indexes().List(context.TODO())
			if err != nil {
				m.log.Error(context.Background(), "Failed to list collection info", "error", err)
			}
			result.Databases = []models.Database{
				{
					Name:         dbName,
					Perms:        "rw-rw-rw-", // Placeholder, as MongoDB does not provide permissions in ListCollectionNames
					Owner:        "root",      // Placeholder, as MongoDB does not provide owner/group in ListCollectionNames
					Group:        "root",      // Placeholder, as MongoDB does not provide owner/group in ListCollectionNames
					Size:         0,           // Placeholder, as MongoDB does not provide size in ListCollectionNames
					LastModified: time.Now(),  // Placeholder, as MongoDB does not provide last modified in ListCollectionNames
				},
			}
			result.Collections = []models.Collection{
				{
					Name:         collectionName,
					Perms:        "rw-rw-rw-", // Placeholder, as MongoDB does not provide permissions in ListCollectionNames
					Owner:        "root",      // Placeholder, as MongoDB does not provide owner/group in ListCollectionNames
					Group:        "root",      // Placeholder, as MongoDB does not provide owner/group in ListCollectionNames
					Size:         0,           // Placeholder, as MongoDB does not provide size in ListCollectionNames
					LastModified: time.Now(),  // Placeholder, as MongoDB does not provide last modified in ListCollectionNames
				},
			}
			var indexes []bson.M
			if err := collectionIndexes.All(context.TODO(), &indexes); err != nil {
				m.log.Error(context.Background(), "Failed to decode collection indexes", "error", err)
			}
			for _, index := range indexes {
				indexName, ok := index["name"].(string)
				if !ok {
					m.log.Error(context.Background(), "Index name is not a string", "index", index)
					continue
				}
				result.Indexes = append(result.Indexes, models.Index{Name: indexName})
			}
			m.log.Debug(context.Background(), "List of indexes", "db", dbName, "collection", collectionName, "indexes", result.Indexes)
		}
	}
	return result, nil
}

// Du is equivalent to the du command in Unix-like systems.
// It is used to calculate the disk usage of database or collection.
func (m mongoClient) Du(duOptions ...models.OptionsDuFunc) (models.DuResponse, error) {
	var result models.DuResponse

	var dbName, collectionName string

	duCmd := &models.DuOptions{
		Args: []string{},
	}

	for _, opt := range duOptions {
		if err := opt(duCmd); err != nil {
			return models.DuResponse{}, err
		}
	}

	m.log.Debug(context.Background(), "provided command with options", "args", duCmd.Args)

	// Find db and collection name
	for _, arg := range duCmd.Args {
		// Extract db and collection names from the argument
		argParts := strings.Split(arg, "/")
		if len(argParts) > 2 {
			m.log.Error(context.Background(), "Expected format is database/collection", "received", arg)
		}
		if len(argParts) == 1 {
			m.log.Debug(context.Background(), "No collection name provided", "received", arg)
			dbName = argParts[0]
			collectionName = ""
		}
		if len(argParts) == 2 {
			m.log.Debug(context.Background(), "Calculating collection size inside db", "received", arg)
			dbName = argParts[0]
			collectionName = argParts[1]
		}
		if collectionName != "" {
			stats := bson.M{}
			err := m.mc.Database(dbName).RunCommand(context.TODO(), bson.D{{Key: "collStats", Value: collectionName}}).Decode(&stats)
			if err != nil {
				m.log.Error(context.Background(), "Failed to get collection stats", collectionName, err)
			}
			m.log.Debug(context.Background(), "Collection size in bytes", collectionName, stats["size"])
			result = models.DuResponse{
				Database:   dbName,
				Collection: collectionName,
				Size:       int64(stats["size"].(int32)),
			}
		} else {
			stats := bson.M{}
			err := m.mc.Database(dbName).RunCommand(context.TODO(), bson.D{{Key: "dbStats", Value: 1}}).Decode(&stats)
			if err != nil {
				m.log.Error(context.Background(), "Failed to get database stats", dbName, err)
			}
			m.log.Debug(context.Background(), "Database size: in bytes", dbName, stats["dataSize"])
			result = models.DuResponse{
				Database:   dbName,
				Collection: "",
				Size:       stats["dataSize"].(int64),
			}
		}
	}
	return result, nil
}

// Touch is equivalent to the touch command in Unix-like systems.
// It is used to create an empty collection in the database.
func (m mongoClient) Touch(touchOptions ...models.OptionsTouchFunc) (models.TouchResponse, error) {

	var dbName, collectionName string
	var result models.TouchResponse

	touchCmd := &models.TouchOptions{
		Args: []string{},
	}

	for _, opt := range touchOptions {
		if err := opt(touchCmd); err != nil {
			return models.TouchResponse{}, err
		}
	}

	m.log.Debug(context.Background(), "provided command with options", "args", touchCmd.Args)
	if len(touchCmd.Args) == 0 {
		m.log.Error(context.Background(), "Expected format is database/collection", "received", "empty")
		return models.TouchResponse{}, errors.New("expected format: iza touch database/collection")
	}

	// Iterate the arguments and create a collection for each
	for _, arg := range touchCmd.Args {
		// Extract db and collection names from the argument
		argParts := strings.Split(arg, "/")
		if len(argParts) > 2 {
			m.log.Error(context.Background(), "Expected format is database/collection", "received", arg)
		}
		if len(argParts) == 1 {
			m.log.Info(context.Background(), "No database provided, creating inside test", "received", arg)
			dbName = "test"
			collectionName = argParts[0]
		}
		if len(argParts) == 2 {
			m.log.Debug(context.Background(), "Creating empty collection", "received", arg)
			dbName = argParts[0]
			collectionName = argParts[1]
		}

		if err := m.mc.Database(dbName).CreateCollection(context.TODO(), collectionName); err != nil {
			m.log.Error(context.Background(), "Failed to create collection", "error", err)
		}
		m.log.Debug(context.Background(), "Successfully created empty collection", "dbName", dbName, "collection", collectionName)
		result.Name = append(result.Name, dbName+"/"+collectionName)
	}

	return result, nil
}

// Cat is equivalent to the cat command in Unix-like systems.
// It is used to display the contents of a document in the collection.
func (m mongoClient) Cat(catOptions ...models.OptionsCatFunc) (models.CatResponse, error) {

	var result models.CatResponse
	catCmd := &models.CatOptions{
		Args: []string{},
	}

	var dbName, collectionName string

	for _, opt := range catOptions {
		if err := opt(catCmd); err != nil {
			return models.CatResponse{}, err
		}
	}

	m.log.Debug(context.Background(), "provided command with options", "args", catCmd.Args)

	// Iterate the arguments and create a collection for each
	for _, arg := range catCmd.Args {
		// Extract db and collection names from the argument
		argParts := strings.Split(arg, "/")
		if len(argParts) > 2 {
			m.log.Error(context.Background(), "Expected format is database/collection", "received", arg)
		}
		if len(argParts) == 1 {
			m.log.Info(context.Background(), "No database provided, reading from test", "received", arg)
			dbName = "test"
			collectionName = argParts[0]
		}
		if len(argParts) == 2 {
			m.log.Debug(context.Background(), "Reading from collection", "received", arg)
			dbName = argParts[0]
			collectionName = argParts[1]
		}

		coll := m.mc.Database(dbName).Collection(collectionName)
		opts := options.Count().SetHint("_id_")
		count, err := coll.CountDocuments(context.TODO(), bson.D{}, opts)
		if err != nil {
			m.log.Error(context.Background(), "Failed to count documents", "error", err)
		}
		m.log.Debug(context.Background(), "Total documents in collection", "dbName", dbName, "collection", collectionName, "count", count)
		result.Count = count
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
		result.Documents = results
	}

	return result, nil
}
