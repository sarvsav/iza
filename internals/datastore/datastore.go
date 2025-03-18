package datastore

import (
	"github.com/sarvsav/iza/dbstore"
	"github.com/sarvsav/iza/foundation/logger"
)

type DataStoreService struct {
	dataStore dbstore.Client
	logger    *logger.Logger
}

func NewDataStoreService(dataStore dbstore.Client, logger *logger.Logger) *DataStoreService {
	return &DataStoreService{
		dataStore: dataStore,
		logger:    logger,
	}
}
