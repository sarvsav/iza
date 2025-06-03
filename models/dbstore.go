package models

import "time"

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

type DatabaseLsResponse interface {
	GetLsResult() (DatabaseLsResponseData, error)
	isDatabaseLsResponse() // marker method
}
