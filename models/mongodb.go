package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MongoDBCatResponse struct {
	Documents []bson.M
	Count     int64
}

type MongoDBDuResponse struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Size       int64  `json:"size"`
}

type MongoDBTouchResponse struct {
	Name []string
}

type MongoDBResult struct {
	MongoDBLsResponse     DatabaseLsResponseData
	MongoDBWhoAmIResponse DatabaseWhoAmIResponseData
}

func (mr MongoDBResult) isDatabaseLsResponse() {
	// marker function
}

func (mr MongoDBResult) GetLsResult() (DatabaseLsResponseData, error) {
	return mr.MongoDBLsResponse, nil
}

func (mr MongoDBResult) GetWhoAmIResult() (DatabaseWhoAmIResponseData, error) {
	return mr.MongoDBWhoAmIResponse, nil
}
