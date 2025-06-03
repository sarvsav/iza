package datastore

import "github.com/sarvsav/iza/models"

// Du returns the disk usage of the specified database or collection.
func (ds *DataStoreService) Du(duOptions ...models.OptionsDuFunc) (models.DatabaseDuResponseData, error) {
	result, err := ds.dataStore.Du(duOptions...)
	if err != nil {
		return models.DatabaseDuResponseData{}, err
	}

	resultData, _ := result.GetDuResult()

	return resultData, nil
}
