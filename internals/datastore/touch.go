package datastore

import (
	"github.com/sarvsav/iza/models"
)

// Touch is a function that creates a new collection in the database
// and returns the name of the collection created.
func (ds *DataStoreService) Touch(touchOptions ...models.OptionsTouchFunc) (models.MongoDBTouchResponse, error) {
	result, err := ds.dataStore.Touch(touchOptions...)
	if err != nil {
		return models.MongoDBTouchResponse{}, err
	}

	return result, nil
}
