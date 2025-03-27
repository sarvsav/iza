package datastore

import (
	"github.com/sarvsav/iza/models"
)

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) WhoAmI() (models.WhoAmIResponse, error) {
	result, err := ds.dataStore.WhoAmI()
	if err != nil {
		return models.WhoAmIResponse{}, err
	}

	return result, nil
}
