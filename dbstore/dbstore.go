package dbstore

import (
	"github.com/sarvsav/iza/models"
)

type Client interface {
	Cat(catOptions ...models.OptionsCatFunc) (models.DatabaseCatResponse, error)
	Du(duOptions ...models.OptionsDuFunc) (models.DatabaseDuResponse, error)
	Ls(lsOptions ...models.OptionsLsFunc) (models.DatabaseLsResponse, error)
	Touch(touchOptions ...models.OptionsTouchFunc) (models.DatabaseTouchResponse, error)
	WhoAmI(whoAmIOptions ...models.OptionsWhoAmIFunc) (models.DatabaseWhoAmIResponse, error)
}
