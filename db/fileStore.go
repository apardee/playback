package db

import (
	"bytes"
	"io/ioutil"

	"encoding/gob"

	"errors"

	"github.com/apardee/playback/model"
)

// PlaybackFileStore uses local file storage implementing PlaybackStore
type PlaybackFileStore struct {
	ClipsArr          []model.MediaClip
	PlaybackStatesArr []model.PlaybackState
}

// Open opens the file store.
func (p *PlaybackFileStore) Open() error {
	clipBytes, err := ioutil.ReadFile("clips.gob")
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

// Clips ...
func (p *PlaybackFileStore) Clips() []model.MediaClip {
	return p.ClipsArr
}

// NewClip ...
func (p *PlaybackFileStore) NewClip() (*model.MediaClip, error) {
	clip, err := model.NewMediaClip()
	if err != nil {
		return nil, err
	}
	p.ClipsArr = append(p.ClipsArr, *clip)
	p.saveClips()
	return clip, nil
}

// UpdateClip ...
func (p *PlaybackFileStore) UpdateClip(clip model.MediaClip) error {
	clipIdx, err := p.clipIndex(clip)
	if err != nil {
		return err
	}
	p.ClipsArr[clipIdx] = clip
	p.saveClips()
	return nil
}

// DeleteClip ...
func (p *PlaybackFileStore) DeleteClip(clip model.MediaClip) error {
	clipIdx, err := p.clipIndex(clip)
	if err != nil {
		return err
	}
	p.ClipsArr = append(p.ClipsArr[:clipIdx], p.ClipsArr[clipIdx+1:]...)
	p.saveClips()
	return nil
}

// PlaybackStates  ...
func (p *PlaybackFileStore) PlaybackStates() []model.PlaybackState {
	return p.PlaybackStatesArr
}

// NewPlaybackState ...
func (p *PlaybackFileStore) NewPlaybackState() (*model.PlaybackState, error) {
	p.saveClips()
	return &model.PlaybackState{}, nil
}

// UpdatePlaybackState ...
func (p *PlaybackFileStore) UpdatePlaybackState(state model.PlaybackState) error {
	p.saveClips()
	return nil
}

// DeletePlaybackClip ...
func (p *PlaybackFileStore) DeletePlaybackClip(state model.PlaybackState) error {
	p.saveClips()
	return nil
}

func (p *PlaybackFileStore) saveClips() error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(p); err != nil {
		return err
	}
	if err := ioutil.WriteFile("clips.gob", buffer.Bytes(), 0777); err != nil {
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
