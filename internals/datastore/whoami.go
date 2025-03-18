package datastore

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) WhoAmI() (string, error) {
	username, err := ds.dataStore.WhoAmI()
	if err != nil {
		return "", err
	}

	return username, nil
}
