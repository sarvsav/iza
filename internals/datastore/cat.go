package datastore

// WhoAmI returns the username of the current user from the datastore.
func (ds *DataStoreService) Cat() error {
	err := ds.dataStore.Cat()
	if err != nil {
		return err
	}

	return nil
}
