package datastore

import (
	"github.com/sarvsav/iza/models"
)

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) Ps(psOptions ...models.OptionsPsFunc) (models.MongoDBPsResponse, error) {
	result, err := ds.dataStore.Ps(psOptions...)
	if err != nil {
		return models.MongoDBPsResponse{}, err
	}

	return result, nil
}
