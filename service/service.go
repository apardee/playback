package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/apardee/playback/model"
)

type context struct {
	store model.PlaybackStore
}

func (c context) clipsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		clips := c.store.Clips()
		byt, err := json.Marshal(clips)
		if err != nil {
			recordRequestError(w, http.StatusInternalServerError, "Failed to prepare clips")
			return
		}
		w.Write(byt)
	} else {
		recordRequestError(w, http.StatusMethodNotAllowed, "")
		return
	}
}

func (c context) clipHandler(w http.ResponseWriter, r *http.Request) {
	components := strings.Split(r.URL.Path, "/")
	if len(components) < 2 {
		recordRequestError(w, http.StatusBadRequest, "Expected a uuid in the URL")
		return
	}

	uuid := components[1]
	if r.Method == http.MethodGet {
		// Find the clip object with the UUID matching the one provided in the request URL.
		clips := c.store.Clips()
		var clipOut *model.MediaClip
		for _, clip := range clips {
			if clip.ClipID.String() == uuid {
				clipOut = &clip
				break
			}
		}

		if clipOut == nil {
			recordRequestError(w, http.StatusNotFound, "")
			return
		}

		byt, err := json.Marshal(clipOut)
		if err != nil {
			recordRequestError(w, http.StatusMethodNotAllowed, "Failed to prepare clips")
			return
		}
		w.Write(byt)
	} else if r.Method == http.MethodPost {
		// Read the body of the request as a clip.
		byt, err := ioutil.ReadAll(r.Body)
		if err != nil {
			recordRequestError(w, http.StatusInternalServerError, "")
			return
		}

		var clip model.MediaClip
		if err := json.Unmarshal(byt, &clip); err != nil {
			recordRequestError(w, http.StatusBadRequest, "Failed to read clip object")
			return
		}

		if err := c.store.UpdateClip(clip); err != nil {
			recordRequestError(w, http.StatusBadRequest, "Failed to read clip object")
			return
		}
	} else {
		recordRequestError(w, http.StatusMethodNotAllowed, "")
		return
	}
}

func (c context) playbackStatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		playbacks := c.store.PlaybackStates()
		bytes, err := json.Marshal(playbacks)
		if err != nil {
			recordRequestError(w, http.StatusInternalServerError, "Failed to prepare clips")
			return
		}
		io.WriteString(w, string(bytes))
	} else {
		recordRequestError(w, http.StatusMethodNotAllowed, "")
		return
	}
}

func (c context) playbackStateHandler(w http.ResponseWriter, r *http.Request) {
	components := strings.Split(r.URL.Path, "/")
	if len(components) < 2 {
		recordRequestError(w, http.StatusBadRequest, "Expected a uuid in the URL")
		return
	}

	uuid := components[1]
	if r.Method == http.MethodGet {
		// Find the playback state object with the UUID matching the one provided in the request URL.
		states := c.store.PlaybackStates()
		var stateOut *model.PlaybackState
		for _, state := range states {
			if state.PlaybackStateID.String() == uuid {
				stateOut = &state
				break
			}
		}

		if stateOut == nil {
			recordRequestError(w, http.StatusNotFound, "")
			return
		}

		byt, err := json.Marshal(stateOut)
		if err != nil {
			recordRequestError(w, http.StatusMethodNotAllowed, "Failed to prepare playback state")
			return
		}
		w.Write(byt)
	} else if r.Method == http.MethodPost {
		// Read the body of the request as a clip.
		byt, err := ioutil.ReadAll(r.Body)
		if err != nil {
			recordRequestError(w, http.StatusInternalServerError, "")
			return
		}

		var state model.PlaybackState
		if err := json.Unmarshal(byt, &state); err != nil {
			recordRequestError(w, http.StatusBadRequest, "Failed to read playback state object")
			return
		}

		if err := c.store.UpdatePlaybackState(state); err != nil {
			recordRequestError(w, http.StatusBadRequest, "Failed to read playback state object")
			return
		}
	} else {
		recordRequestError(w, http.StatusMethodNotAllowed, "")
		return
	}
}

func (c context) uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		recordRequestError(w, http.StatusMethodNotAllowed, "Uploaded files must be posted")
		return
	}

	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		recordRequestError(w, http.StatusInternalServerError, "Failed to read file")
		return
	}

	if err := c.store.CommitMediaFile(byt); err != nil {
		recordRequestError(w, http.StatusInternalServerError, "Failed to commit file")
		return
	}
}

func recordRequestError(w http.ResponseWriter, status int, message string) {
	if message == "" {
		message = http.StatusText(status)
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, "%d - %s", status, message)
}

func logger(hnd http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hnd(w, r)
		log.Printf(r.RemoteAddr + " - " + r.URL.Path)
	}
}

// RunService starts up the http service for serving up Playback objects.
func RunService(store model.PlaybackStore) error {

	ctx := context{store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("/clips", logger(ctx.clipsHandler))
	mux.HandleFunc("/clip/", logger(ctx.clipHandler))
	mux.HandleFunc("/playback_states", logger(ctx.playbackStatesHandler))
	mux.HandleFunc("/playback_state/", logger(ctx.playbackStateHandler))
	mux.HandleFunc("/upload_file", logger(ctx.uploadFileHandler))

	return http.ListenAndServe(":8080", mux)
}
