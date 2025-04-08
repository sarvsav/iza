package datastore

import "github.com/sarvsav/iza/models"

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) Ls(lsOptions ...models.OptionsLsFunc) ([]string, error) {
	result, err := ds.dataStore.Ls(lsOptions...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
