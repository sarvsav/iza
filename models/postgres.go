package models

type PostgresResult struct {
	PostgresCatResponse    DatabaseCatResponseData
	PostgresDuResponse     DatabaseDuResponseData
	PostgresLsResponse     DatabaseLsResponseData
	PostgresTouchResponse  DatabaseTouchResponseData
	PostgresWhoAmIResponse DatabaseWhoAmIResponseData
}

func (pr PostgresResult) isDatabaseLsResponse() {
	// marker function
}

func (pr PostgresResult) GetCatResult() (DatabaseCatResponseData, error) {
	return pr.PostgresCatResponse, nil
}

func (pr PostgresResult) GetDuResult() (DatabaseDuResponseData, error) {
	return pr.PostgresDuResponse, nil
}

func (pr PostgresResult) GetLsResult() (DatabaseLsResponseData, error) {
	return pr.PostgresLsResponse, nil
}

func (pr PostgresResult) GetTouchResult() (DatabaseTouchResponseData, error) {
	return pr.PostgresTouchResponse, nil
}

func (pr PostgresResult) GetWhoAmIResult() (DatabaseWhoAmIResponseData, error) {
	return pr.PostgresWhoAmIResponse, nil
}
