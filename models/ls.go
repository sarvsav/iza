package models

type OptionsLsFunc func(c *LsOptions) error

type LsOptions struct {
	LongListing bool
	Color       bool
	Args        []string
}
