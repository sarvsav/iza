package datastore

import (
	"github.com/sarvsav/iza/models"
)

// Touch is a function that creates a new collection in the database
// and returns the name of the collection created.
func (ds *DataStoreService) Touch(touchOptions ...models.OptionsTouchFunc) (models.DatabaseTouchResponseData, error) {
	result, err := ds.dataStore.Touch(touchOptions...)
	if err != nil {
		return models.DatabaseTouchResponseData{}, err
	}

	resultData, err := result.GetTouchResult()
	if err != nil {
		return models.DatabaseTouchResponseData{}, err
	}

	return resultData, nil
}
