package model

// PlaybackStore defines the persistent interface for model objects.
type PlaybackStore interface {
	Open() error

	Clips() []MediaClip
	NewClip() (*MediaClip, error)
	UpdateClip(clip MediaClip) error
	DeleteClip(clip MediaClip) error

	PlaybackStates() []PlaybackState
	NewPlaybackState() (*PlaybackState, error)
	UpdatePlaybackState(state PlaybackState) error
	DeletePlaybackClip(state PlaybackState) error
}
