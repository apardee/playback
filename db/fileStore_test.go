package db

import "testing"

func TestThing(t *testing.T) {
	fileStore := PlaybackFileStore{}
	fileStoreConfig := PlaybackFileStoreConfig{DataPath: "testData"}
	err := fileStore.Open(fileStoreConfig)
	if err != nil {
		t.Errorf("Failed to create and open the filestore")
	}
}
