package model

import (
	"encoding/hex"
	"testing"
)

func TestMediaClip(t *testing.T) {
	_, err := NewMediaClip()
	if err != nil {
		t.Errorf("failed to create a clip")
	}
}
func TestPlaybackState(t *testing.T) {
	_, err := NewPlaybackState()
	if err != nil {
		t.Errorf("failed to create a playback state")
	}
}

func TestUUID(t *testing.T) {
	{
		_, err := NewUUID()
		if err != nil {
			t.Errorf("failed to create a uuid")
		}
	}

	{
		testUUID := "74309f6e-ad6f-424b-5c74-7102c7d7f1c2"
		uuid, _ := NewUUID()
		byt := [16]byte(*uuid)
		t.Errorf(uuid.String())
		t.Errorf(hex.Dump(byt[:]))
		// uuidEncoding := []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
		// var uuid UUID
		// if err := json.Unmarshal(uuidEncoding, &uuid); err != nil {
		// 	t.Errorf("failed to unmarshal the test uuid - %s", err.Error())
		// }
	}
}
