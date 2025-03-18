package datastore

import "github.com/sarvsav/iza/dbstore"

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) Du(duOptions ...dbstore.OptionsDuFunc) error {
	err := ds.dataStore.Du(duOptions...)
	if err != nil {
		return err
	}
	return nil
}
