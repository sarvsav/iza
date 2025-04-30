package models

type OptionsCatFunc func(c *CatOptions) error

type CatOptions struct {
	Args []string
}

type OptionsDuFunc func(c *DuOptions) error

type DuOptions struct {
	Args []string
}

type OptionsLsFunc func(c *LsOptions) error

type LsOptions struct {
	LongListing bool
	Color       bool
	Args        []string
}

type OptionsTouchFunc func(c *TouchOptions) error

type TouchOptions struct {
	Args []string
}

type OptionsWhoAmIFunc func(c *WhoAmIOptions) error

type WhoAmIOptions struct {
	Args []string
}
