package datastore

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) Ls() ([]string, error) {
	result, err := ds.dataStore.Ls()
	if err != nil {
		return nil, err
	}

	return result, nil
}
