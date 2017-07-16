package model

import uuid "github.com/nu7hatch/gouuid"

// MediaClip is the object representation of a media file,
// along with associated metadata.
type MediaClip struct {
	ClipID *uuid.UUID
	FileID *uuid.UUID
	Title  string
	Length int64
}

// NewMediaClip creates a new, unbound media clip that can be associated with a file and the associated metadata
func NewMediaClip() (*MediaClip, error) {
	clipID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &MediaClip{ClipID: clipID, FileID: nil, Title: "", Length: 0}, nil
}
