package models

import "time"

type OptionsLsFunc func(c *LsOptions) error

type LsOptions struct {
	LongListing bool
	Color       bool
	Args        []string
}

type LsResponse struct {
	Databases   []Database
	Collections []Collection
	Indexes     []Index
}

type Database struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (r/w/x)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last modified date
}

type Collection struct {
	Name         string
	Size         int64     // In bytes
	Perms        string    // Permissions (r/w/x)
	Owner        string    // Owner
	Group        string    // Group
	LastModified time.Time // Last modified date
}

type Index struct {
	Name string
}
