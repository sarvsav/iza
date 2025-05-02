package dbstore

import (
	"github.com/sarvsav/iza/models"
)

type Client interface {
	Du(duOptions ...models.OptionsDuFunc) (models.MongoDBDuResponse, error)
	Ls(lsOptions ...models.OptionsLsFunc) (models.MongoDBLsResponse, error)
	Ps(psOptions ...models.OptionsPsFunc) (models.MongoDBPsResponse, error)
	Touch(touchOptions ...models.OptionsTouchFunc) (models.MongoDBTouchResponse, error)
	WhoAmI(whoAmIOptions ...models.OptionsWhoAmIFunc) (models.MongoDBWhoAmIResponse, error)
	Cat(catOptions ...models.OptionsCatFunc) (models.MongoDBCatResponse, error)
}
