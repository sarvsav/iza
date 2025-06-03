package dbstore

import (
	"github.com/sarvsav/iza/models"
)

type Client interface {
	Du(duOptions ...models.OptionsDuFunc) (models.MongoDBDuResponse, error)
	Ls(lsOptions ...models.OptionsLsFunc) (models.DatabaseLsResponse, error)
	Touch(touchOptions ...models.OptionsTouchFunc) (models.DatabaseTouchResponse, error)
	WhoAmI(whoAmIOptions ...models.OptionsWhoAmIFunc) (models.DatabaseWhoAmIResponse, error)
	Cat(catOptions ...models.OptionsCatFunc) (models.MongoDBCatResponse, error)
}
