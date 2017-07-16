package model

import uuid "github.com/nu7hatch/gouuid"

// PlaybackState represents the current state of playback (in seconds)
// of a particular media clip.
type PlaybackState struct {
	PlaybackStateID *uuid.UUID
	ClipID          *uuid.UUID
	Location        int64
}

// NewPlaybackState creates a new playback state object associated with a clip and bookmark location
func NewPlaybackState() (*PlaybackState, error) {
	return nil, nil
}
