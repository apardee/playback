package model

import (
	"time"
)

// PlaybackState represents the current state of playback (in seconds)
// of a particular media clip.
type PlaybackState struct {
	PlaybackStateID *UUID         `json:"playback_state_id"`
	ClipID          *UUID         `json:"clip_id"`
	Location        time.Duration `json:"location"`
}

// NewPlaybackState creates a new playback state object associated with a clip and bookmark location
func NewPlaybackState() (*PlaybackState, error) {
	playbackStateID, err := NewUUID()
	if err != nil {
		return nil, err
	}
	return &PlaybackState{PlaybackStateID: playbackStateID, ClipID: nil, Location: 0}, nil
}
