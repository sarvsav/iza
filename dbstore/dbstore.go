package dbstore

import (
	"github.com/sarvsav/iza/models"
)

type Client interface {
	Du(duOptions ...models.OptionsDuFunc) error
	Ls(lsOptions ...models.OptionsLsFunc) (models.LsResponse, error)
	Touch(touchOptions ...models.OptionsTouchFunc) (models.TouchResponse, error)
	WhoAmI(whoAmIOptions ...models.OptionsWhoAmIFunc) (models.WhoAmIResponse, error)
	Cat(catOptions ...models.OptionsCatFunc) (models.CatResponse, error)
}
