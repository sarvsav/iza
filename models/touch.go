package models

type OptionsTouchFunc func(c *TouchOptions) error

type TouchOptions struct {
	Args []string
}

type TouchResponse struct {
	Name []string
}
