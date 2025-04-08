package datastore

import "github.com/sarvsav/iza/models"

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) Cat(catOptions ...models.OptionsCatFunc) error {
	err := ds.dataStore.Cat(catOptions...)
	if err != nil {
		return err
	}

	return nil
}
