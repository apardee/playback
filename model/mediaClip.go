package model

// MediaClip is the object representation of a media file,
// along with associated metadata.
type MediaClip struct {
	ClipID *UUID  `json:"clip_id"`
	FileID *UUID  `json:"file_id"`
	Title  string `json:"title"`
	Length int64  `json:"length"`
}

// NewMediaClip creates a new, unbound media clip that can be associated with a file and the associated metadata
func NewMediaClip() (*MediaClip, error) {
	clipID, err := NewUUID()
	if err != nil {
		return nil, err
	}
	return &MediaClip{ClipID: clipID, FileID: nil, Title: "", Length: 0}, nil
}
