package models

type MongoDBResult struct {
	MongoDBCatResponse    DatabaseCatResponseData
	MongoDBDuResponse     DatabaseDuResponseData
	MongoDBLsResponse     DatabaseLsResponseData
	MongoDBTouchResponse  DatabaseTouchResponseData
	MongoDBWhoAmIResponse DatabaseWhoAmIResponseData
}

func (mr MongoDBResult) isDatabaseLsResponse() {
	// marker function
}

func (mr MongoDBResult) GetCatResult() (DatabaseCatResponseData, error) {
	return mr.MongoDBCatResponse, nil
}

func (mr MongoDBResult) GetDuResult() (DatabaseDuResponseData, error) {
	return mr.MongoDBDuResponse, nil
}

func (mr MongoDBResult) GetLsResult() (DatabaseLsResponseData, error) {
	return mr.MongoDBLsResponse, nil
}

func (mr MongoDBResult) GetTouchResult() (DatabaseTouchResponseData, error) {
	return mr.MongoDBTouchResponse, nil
}

func (mr MongoDBResult) GetWhoAmIResult() (DatabaseWhoAmIResponseData, error) {
	return mr.MongoDBWhoAmIResponse, nil
}
