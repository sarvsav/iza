package dbstore

import (
	"github.com/sarvsav/iza/models"
)

type OptionsWhoAmIFunc func(c *models.WhoAmIOptions) error
type OptionsLsFunc func(c *models.LsOptions) error
type OptionsDuFunc func(c *models.DuOptions) error
type OptionsTouchFunc func(c *models.TouchOptions) error
type OptionsCatFunc func(c *models.CatOptions) error

type Client interface {
	Du(duOptions ...OptionsDuFunc) error
	Ls(lsOptions ...OptionsLsFunc) ([]string, error)
	Touch(touchOptions ...OptionsTouchFunc) (string, error)
	WhoAmI(whoAmIOptions ...OptionsWhoAmIFunc) (models.WhoAmIResponse, error)
	Cat(catOptions ...OptionsCatFunc) error
}
