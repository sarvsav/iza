package internals

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"github.com/sarvsav/iza/database"
	"github.com/sarvsav/iza/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OptionsWhoAmIFunc func(c *models.WhoAmIOptions) error

// WhoAmI is equivalent to the whoami command.
// It prints the current logged in user.
func WhoAmI(whoAmIOptions ...OptionsWhoAmIFunc) error {
	whoAmICmd := &models.WhoAmIOptions{
		Args:   []string{},
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	for _, opt := range whoAmIOptions {
		if err := opt(whoAmICmd); err != nil {
			return err
		}
	}

	whoAmICmd.Logger.Debug("provided command with options", "args", whoAmICmd.Args)

	client, err := database.GetMongoClient()
	defer func() {
		if err := database.DisconnectMongoClient(client); err != nil {
			whoAmICmd.Logger.Error("Failed to disconnect from MongoDB", "error", err)
		}
	}()

	if err != nil {
		whoAmICmd.Logger.Error("Failed to connect to MongoDB", "error", err)
		return err
	}

	info := bson.M{}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "connectionStatus", Value: 1}}).Decode(&info); err != nil {
		whoAmICmd.Logger.Error("Failed to get connection status", "error", err)
		return err
	}

	// Accessing "authenticatedUsers"
	authInfo, ok := info["authInfo"].(bson.M)
	if !ok {
		whoAmICmd.Logger.Error("authInfo is not of type bson.M")
		return errors.New("authInfo is not of type bson.M")
	}

	authenticatedUsers, ok := authInfo["authenticatedUsers"].(primitive.A)
	if !ok {
		whoAmICmd.Logger.Error("authenticatedUsers is not of type []interface{}", "authenticatedUsers", authInfo["authenticatedUsers"])
		return errors.New("authenticatedUsers is not of type []interface{}")
	}

	// Accessing user details
	if len(authenticatedUsers) > 0 {
		user, ok := authenticatedUsers[0].(bson.M)
		if !ok {
			whoAmICmd.Logger.Error("first element in authenticatedUsers is not of type bson.M")
			return errors.New("first element in authenticatedUsers is not of type bson.M")
		}

		// Extract the "user" field
		username, ok := user["user"].(string)
		if ok {
			whoAmICmd.Logger.Info("Authenticated", "username", username)
		} else {
			whoAmICmd.Logger.Error("User field is not a string or does not exist")
		}
	} else {
		whoAmICmd.Logger.Error("authenticatedUsers is empty")
	}

	return nil
}
