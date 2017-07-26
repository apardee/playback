package service

import (
	"encoding/json"
	"io"
	"net/http"

	"io/ioutil"

	"github.com/apardee/playback/model"
)

type context struct {
	store model.PlaybackStore
}

func (c context) clipsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		clips := c.store.Clips()
		bytes, err := json.Marshal(clips)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "500 - Failed to prepare clips")
			return
		}
		io.WriteString(w, string(bytes))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "405 - Method not allowed")
		return
	}
}

func (c context) clipHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "TODO: Clip updating")
}

func (c context) playbackStatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		playbacks := c.store.PlaybackStates()
		bytes, err := json.Marshal(playbacks)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "500 - Failed to prepare clips")
			return
		}
		io.WriteString(w, string(bytes))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "405 - Method not allowed")
		return
	}
}

func (c context) playbackStateHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "TODO: Playback state updating")
}

func (c context) uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// byt, err := httputil.DumpRequest(r, true)
	// if err == nil {
	// 	log.Println(string(byt))
	// }
	// fmt.Println(r.Header.Get("Content-Type"))
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "405 - Uploaded files must be posted")
		return
	}

	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "500 - Failed to read file")
		return
	}
	ioutil.WriteFile("test.mp3", byt, 7777)
}

// RunService starts up the http service for serving up Playback objects.
func RunService(store model.PlaybackStore) error {

	ctx := context{store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("/clips", ctx.clipsHandler)
	mux.HandleFunc("/clip/", ctx.clipHandler)
	mux.HandleFunc("/playback_states", ctx.playbackStatesHandler)
	mux.HandleFunc("/playback_state/", ctx.playbackStateHandler)
	mux.HandleFunc("/upload_file", ctx.uploadFileHandler)

	return http.ListenAndServe(":8080", mux)
}
