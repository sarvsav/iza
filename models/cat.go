package models

type OptionsCatFunc func(c *CatOptions) error

type CatOptions struct {
	Args []string
}
