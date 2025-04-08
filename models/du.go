package models

type OptionsDuFunc func(c *DuOptions) error

type DuOptions struct {
	Args []string
}
