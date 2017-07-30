package db

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"os"

	"github.com/apardee/playback/audio"
	"github.com/apardee/playback/model"
)

const localFilename string = "clips/clips.gob"

// PlaybackFileStore uses local file storage implementing PlaybackStore
type PlaybackFileStore struct {
	ClipsArr          []model.MediaClip
	PlaybackStatesArr []model.PlaybackState
}

// Open opens the file store.
func (p *PlaybackFileStore) Open() error {
	os.Mkdir("clips", 0777)

	clipBytes, err := ioutil.ReadFile(localFilename)
	if err != nil {
		p.ClipsArr = []model.MediaClip{}
		p.PlaybackStatesArr = []model.PlaybackState{}
		return nil
	}

	clipReader := bytes.NewReader(clipBytes)
	clipDecoder := gob.NewDecoder(clipReader)
	if err = clipDecoder.Decode(p); err != nil {
		return err
	}

	if p.ClipsArr == nil {
		p.ClipsArr = []model.MediaClip{}
	}

	if p.PlaybackStatesArr == nil {
		p.PlaybackStatesArr = []model.PlaybackState{}
	}

	return nil
}

// Clips retrieves all active clip objects from the store.
func (p *PlaybackFileStore) Clips() []model.MediaClip {
	return p.ClipsArr
}

// NewClip creates and returns a new clip object. The new object is immediately saved to the store.
func (p *PlaybackFileStore) NewClip() (*model.MediaClip, error) {
	clip, err := model.NewMediaClip()
	if err != nil {
		return nil, err
	}
	p.ClipsArr = append(p.ClipsArr, *clip)
	return clip, p.saveObjects()
}

// UpdateClip updates the state of a clip.
func (p *PlaybackFileStore) UpdateClip(clip model.MediaClip) error {
	clipIdx, err := p.clipIndex(clip)
	if err != nil {
		return err
	}
	p.ClipsArr[clipIdx] = clip
	return p.saveObjects()
}

// DeleteClip removes a clip from the file store.
func (p *PlaybackFileStore) DeleteClip(clip model.MediaClip) error {
	clipIdx, err := p.clipIndex(clip)
	if err != nil {
		return err
	}
	p.ClipsArr = append(p.ClipsArr[:clipIdx], p.ClipsArr[clipIdx+1:]...)
	return p.saveObjects()
}

// PlaybackStates retrieves all active playback states.
func (p *PlaybackFileStore) PlaybackStates() []model.PlaybackState {
	return p.PlaybackStatesArr
}

// NewPlaybackState creates a new playback state object that's immediately saved to the store.
func (p *PlaybackFileStore) NewPlaybackState() (*model.PlaybackState, error) {
	playback, err := model.NewPlaybackState()
	if err != nil {
		return nil, err
	}
	p.PlaybackStatesArr = append(p.PlaybackStatesArr, *playback)
	return playback, p.saveObjects()
}

// UpdatePlaybackState updates the state of an existing playback state object.
func (p *PlaybackFileStore) UpdatePlaybackState(state model.PlaybackState) error {
	index, err := p.playbackStateIndex(state)
	if err != nil {
		return err
	}
	p.PlaybackStatesArr[index] = state
	return p.saveObjects()
}

// DeletePlaybackState deletes an active playback state object from the store.
func (p *PlaybackFileStore) DeletePlaybackState(state model.PlaybackState) error {
	index, err := p.playbackStateIndex(state)
	if err != nil {
		return err
	}
	p.PlaybackStatesArr = append(p.PlaybackStatesArr[:index], p.PlaybackStatesArr[index+1:]...)
	return p.saveObjects()
}

// CommitMediaFile commits the data for an uploaded media file, creating the associated MediaClip object and saving the data to the filesystem.
func (p *PlaybackFileStore) CommitMediaFile(byt []byte) error {
	// Create the uuid for the media file to be stored
	uuid, err := model.NewUUID()
	if err != nil {
		return err
	}

	clip, err := p.NewClip()
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("clips/"+uuid.String(), byt, 0777); err != nil {
		return err
	}

	length, err := audio.MP3Length(byt)
	if err != nil {
		return err
	}

	clip.Title = "New title..."
	clip.Length = length
	clip.FileID = uuid

	if err := p.UpdateClip(*clip); err != nil {
		return err
	}

	return err
}

func (p *PlaybackFileStore) saveObjects() error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(p); err != nil {
		return err
	}
	if err := ioutil.WriteFile(localFilename, buffer.Bytes(), 0777); err != nil {
		return err
	}
	return nil
}

func (p *PlaybackFileStore) clipIndex(clip model.MediaClip) (int, error) {
	idxOut := -1
	for idx, clp := range p.ClipsArr {
		if clp.ClipID == clip.ClipID {
			idxOut = idx
			break
		}
	}

	if idxOut < 0 {
		return -1, errors.New("Clip not found")
	}
	return idxOut, nil
}

func (p *PlaybackFileStore) playbackStateIndex(playback model.PlaybackState) (int, error) {
	idxOut := -1
	for idx, state := range p.PlaybackStatesArr {
		if playback.PlaybackStateID == state.PlaybackStateID {
			idxOut = idx
			break
		}
	}

	if idxOut < 0 {
		return -1, errors.New("PlaybackState not found")
	}
	return idxOut, nil
}
