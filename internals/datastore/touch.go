package datastore

// Touch returns the username of the current user from the datastore.
func (ds *DataStoreService) Touch() (string, error) {
	collectionCreated, err := ds.dataStore.Touch()
	if err != nil {
		return "", err
	}

	return collectionCreated, nil
}
