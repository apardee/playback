package db

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/apardee/playback/model"
)

var fileStoreConfig PlaybackFileStoreConfig

func TestFileStore(t *testing.T) {
	fileStoreConfig = PlaybackFileStoreConfig{DataPath: "testData"}
	fileStore := PlaybackFileStore{}
	defer os.RemoveAll(fileStoreConfig.DataPath)
	err := fileStore.Open(fileStoreConfig)
	if err != nil {
		t.Errorf("failed to create and open the filestore")
	}

	testPlaybackStates(t, &fileStore)
	testClips(t, &fileStore)
	testReload(t)
	testDelete(t, &fileStore)
	testMediaUpload(t, &fileStore)
}

func testClips(t *testing.T, fileStore *PlaybackFileStore) {
	clipID := ""
	fileID := ""

	// Create & Update a clip object with the file
	t.Run("Adding & Updating Clips", func(t *testing.T) {
		// Create and edit a clip.
		clip, err := fileStore.NewClip()
		if err != nil {
			t.Errorf("could not create a new clip from the filestore")
		}

		// Make sure the clip's actually been added to the store.
		if len(fileStore.Clips()) != 1 {
			t.Errorf("adding a clip to the filestore failed")
		}

		clip.FileID, err = model.NewUUID()
		if err != nil {
			t.Errorf("failed to create a uuid for the made up file reference")
		}
		clip.Length = 123456789
		clip.Title = "Clip Test Title"

		clipID = clip.ClipID.String()
		fileID = clip.FileID.String()
		if err := fileStore.UpdateClip(*clip); err != nil {
			t.Errorf("failed to update the clip")
		}

		bogusClip := model.MediaClip{}
		err = fileStore.UpdateClip(bogusClip)
		if err == nil || err.Error() != "Clip not found" {
			t.Errorf("expected the clip update to fail")
		}
	})

	t.Run("Evaluate Clip State", func(t *testing.T) {
		clips := fileStore.Clips()

		// Clip count should still be 1
		if len(fileStore.Clips()) != 1 {
			t.Errorf("adding a clip to the filestore failed")
		}

		clip := clips[0]
		if clip.Length != 123456789 || clip.Title != "Clip Test Title" || clip.ClipID.String() != clipID || clip.FileID.String() != fileID {
			t.Errorf("clip attributes don't match the clip that was saved to the filestore")
		}
	})
}

func testPlaybackStates(t *testing.T, fileStore *PlaybackFileStore) {
	playbackStateID := ""
	clipID := ""
	t.Run("Adding & Updating Playback States", func(t *testing.T) {
		// Create and edit a playback states.
		playback, err := fileStore.NewPlaybackState()
		if err != nil {
			t.Errorf("could not create a new clip from the filestore")
		}

		// Make sure the playback state's actually been added to the store.
		if len(fileStore.PlaybackStates()) != 1 {
			t.Errorf("adding a playback state to the filestore failed")
		}

		playback.ClipID, err = model.NewUUID()
		if err != nil {
			t.Errorf("failed to create a uuid for playback state")
		}
		playback.Location = 123456789

		playbackStateID = playback.PlaybackStateID.String()
		clipID = playback.ClipID.String()
		if err := fileStore.UpdatePlaybackState(*playback); err != nil {
			t.Errorf("failed to update the playback state")
		}

		bogusState := model.PlaybackState{}
		err = fileStore.UpdatePlaybackState(bogusState)
		if err == nil || err.Error() != "PlaybackState not found" {
			t.Errorf("expected the playback state update to fail")
		}
	})
}

func testReload(t *testing.T) {
	fileStore := PlaybackFileStore{}
	err := fileStore.Open(fileStoreConfig)
	if err != nil {
		t.Errorf("failed to create and open the filestore")
	}

	if len(fileStore.Clips()) != 1 {
		t.Errorf("reloaded filestore has the wrong number of clips")
	}

	if len(fileStore.PlaybackStates()) != 1 {
		t.Errorf("reloaded filestore has the wrong number of playback states")
	}
}

func testDelete(t *testing.T, fileStore *PlaybackFileStore) {
	clip := fileStore.Clips()[0]
	if err := fileStore.DeleteClip(clip); err != nil {
		t.Errorf("failed to delete clip")
	}

	playbackState := fileStore.PlaybackStates()[0]
	if err := fileStore.DeletePlaybackState(playbackState); err != nil {
		t.Errorf("failed to delete playback state")
	}

	if len(fileStore.Clips()) != 0 || len(fileStore.PlaybackStates()) != 0 {
		t.Errorf("filestore has the wrong number of clips or playback states after deletion")
	}
}

func testMediaUpload(t *testing.T, fileStore *PlaybackFileStore) {
	byt, err := ioutil.ReadFile("../testData/test.mp3")
	if err != nil {
		t.Errorf("failed to read the test mp3 file")
	}

	if err := fileStore.CommitMediaFile(byt); err != nil {
		t.Errorf("failed to commit the test mp3 - %s", err.Error())
	}
}
