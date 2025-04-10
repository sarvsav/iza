package datastore

import "github.com/sarvsav/iza/models"

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) Cat(catOptions ...models.OptionsCatFunc) (models.CatResponse, error) {
	result, err := ds.dataStore.Cat(catOptions...)
	if err != nil {
		return models.CatResponse{}, err
	}

	return result, nil
}
