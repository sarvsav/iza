package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type DatabaseCatResponseData struct {
	Documents []bson.M
	Count     int64
}

type DatabaseDuResponseData struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Size       int64  `json:"size"`
}

type DatabaseLsResponseData struct {
	DatabaseDatabases   []DatabaseDatabaseData
	DatabaseCollections []DatabaseCollectionData
	DatabaseIndexes     []DatabaseIndexData
}

type DatabaseDatabaseData struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (read/stop/start/create/configure)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last run
}

type DatabaseCollectionData struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (read/stop/start/create/configure)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last run
}

type DatabaseIndexData struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (read/stop/start/create/configure)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last run
}

type DatabaseTouchResponseData struct {
	Name []string
}

type DatabaseWhoAmIResponseData struct {
	Username string
}

type DatabaseCatResponse interface {
	GetCatResult() (DatabaseCatResponseData, error)
}

type DatabaseDuResponse interface {
	GetDuResult() (DatabaseDuResponseData, error)
}

type DatabaseLsResponse interface {
	GetLsResult() (DatabaseLsResponseData, error)
	isDatabaseLsResponse() // marker method
}

type DatabaseTouchResponse interface {
	GetTouchResult() (DatabaseTouchResponseData, error)
}

type DatabaseWhoAmIResponse interface {
	GetWhoAmIResult() (DatabaseWhoAmIResponseData, error)
}
