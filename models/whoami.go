package models

type OptionsWhoAmIFunc func(c *WhoAmIOptions) error

type WhoAmIOptions struct {
	Args []string
}

type WhoAmIResponse struct {
	Username string
}
