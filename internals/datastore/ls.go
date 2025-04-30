package datastore

import (
	"github.com/sarvsav/iza/models"
)

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) Ls(lsOptions ...models.OptionsLsFunc) (models.MongoDBLsResponse, error) {
	result, err := ds.dataStore.Ls(lsOptions...)
	if err != nil {
		return models.MongoDBLsResponse{}, err
	}

	return result, nil
}
