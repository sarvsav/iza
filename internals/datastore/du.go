package datastore

import "github.com/sarvsav/iza/models"

// Du returns the disk usage of the specified database or collection.
func (ds *DataStoreService) Du(duOptions ...models.OptionsDuFunc) (models.MongoDBDuResponse, error) {
	result, err := ds.dataStore.Du(duOptions...)
	if err != nil {
		return models.MongoDBDuResponse{}, err
	}
	return result, nil
}
