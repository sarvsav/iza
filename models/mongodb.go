package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type MongoDBCatResponse struct {
	Documents []bson.M
	Count     int64
}

type MongoDBDuResponse struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Size       int64  `json:"size"`
}

type MongoDBLsResponse struct {
	Databases   []MongoDBDatabase
	Collections []MongoDBCollection
	Indexes     []MongoDBIndex
}

type MongoDBPsResponse struct {
	Data []any
}

type MongoDBDatabase struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (r/w/x)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last modified date
}

type MongoDBCollection struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (r/w/x)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last modified date
}

type MongoDBIndex struct {
	Name string
}

type MongoDBTouchResponse struct {
	Name []string
}

type MongoDBWhoAmIResponse struct {
	Username string
}
