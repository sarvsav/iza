package models

type OptionsDuFunc func(c *DuOptions) error

type DuOptions struct {
	Args []string
}

type DuResponse struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Size       int64  `json:"size"`
}
