package dbstore

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sarvsav/iza/foundation/logger"
	"github.com/sarvsav/iza/models"
)

type postgresClient struct {
	pc  *pgxpool.Pool
	log *logger.Logger
}

func NewPostgresClient(pc *pgxpool.Pool, log *logger.Logger) *postgresClient {
	return &postgresClient{
		pc:  pc,
		log: log,
	}
}

// Cat is equivalent to the cat command in Unix-like systems.
// It is used to display the contents of a document in the collection.
func (p postgresClient) Cat(catOptions ...models.OptionsCatFunc) (models.DatabaseCatResponse, error) {
	return models.PostgresResult{}, nil
}

func (p postgresClient) Du(duOptions ...models.OptionsLsFunc) (models.DatabaseDuResponse, error) {
	return models.PostgresResult{}, nil
}

func (p postgresClient) Ls(lsOptions ...models.OptionsLsFunc) (models.DatabaseLsResponse, error) {
	return models.PostgresResult{}, nil
}

// Touch is equivalent to the touch command in Unix-like systems.
// It is used to create an empty collection in the database.
func (p postgresClient) Touch(touchOptions ...models.OptionsTouchFunc) (models.DatabaseTouchResponse, error) {
	return models.PostgresResult{}, nil
}

// WhoAmI is equivalent to the whoami command.
// It prints the current logged in user.
func (p postgresClient) WhoAmI(whoAmIOptions ...models.OptionsWhoAmIFunc) (models.DatabaseWhoAmIResponse, error) {
	var userName string
	err := p.pc.QueryRow(context.Background(), "SELECT current_user").Scan(&userName)
	if err != nil {
		p.log.Error(context.Background(), "failed to get current user", "error", err)
		return models.PostgresResult{}, err
	}
	return models.PostgresResult{
		PostgresWhoAmIResponse: models.DatabaseWhoAmIResponseData{
			Username: userName,
		},
	}, nil
}
