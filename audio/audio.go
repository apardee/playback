package audio

import (
	"bytes"
	"errors"
	"time"

	"github.com/tcolgate/mp3"
)

// MP3Length estimates the length of an mp3 audio file.
func MP3Length(byt []byte) (time.Duration, error) {
	buf := bytes.NewBuffer(byt)
	dec := mp3.NewDecoder(buf)
	if dec == nil {
		return time.Duration(0), errors.New("failed to create the decoder for the audio data passed")
	}

	dur := time.Duration(0)
	var f mp3.Frame
	for {
		skipped := 0
		if err := dec.Decode(&f, &skipped); err != nil {
			break
		}
		dur += f.Duration()
	}
	return dur, nil
}
