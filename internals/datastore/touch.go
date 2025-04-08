package datastore

import (
	"github.com/sarvsav/iza/models"
)

// Touch is a function that creates a new collection in the database
// and returns the name of the collection created.
func (ds *DataStoreService) Touch(touchOptions ...models.OptionsTouchFunc) (string, error) {
	collectionCreated, err := ds.dataStore.Touch(touchOptions...)
	if err != nil {
		return "", err
	}

	return collectionCreated.Name, nil
}
