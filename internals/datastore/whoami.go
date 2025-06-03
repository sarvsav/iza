package datastore

import (
	"github.com/sarvsav/iza/models"
)

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) WhoAmI() (models.DatabaseWhoAmIResponseData, error) {
	result, err := ds.dataStore.WhoAmI()
	if err != nil {
		return models.DatabaseWhoAmIResponseData{}, err
	}

	resultData, err := result.GetWhoAmIResult()
	if err != nil {
		return models.DatabaseWhoAmIResponseData{}, err
	}

	return resultData, nil
}
