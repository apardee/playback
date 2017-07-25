package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"log"
	"net/http/httputil"

	"github.com/apardee/playback/model"
)

// RunService starts up the http service for serving up Playback objects.
func RunService(store model.PlaybackStore) error {
	http.HandleFunc("/clips", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			clips := store.Clips()
			bytes, err := json.Marshal(clips)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, "500 - Failed to prepare clips")
			}
			io.WriteString(w, string(bytes))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, "405 - Method not allowed")
		}
	})

	http.HandleFunc("/clip/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "TODO: Clip updating")
	})

	http.HandleFunc("/playback_states", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			playbacks := store.PlaybackStates()
			bytes, err := json.Marshal(playbacks)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, "500 - Failed to prepare clips")
			}
			io.WriteString(w, string(bytes))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, "405 - Method not allowed")
		}
	})

	http.HandleFunc("/playback_state/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "TODO: Playback state updating")
	})

	http.HandleFunc("/upload_file", func(w http.ResponseWriter, r *http.Request) {
		byt, err := httputil.DumpRequest(r, true)
		if err == nil {
			log.Println(string(byt))
		}

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, "405 - Uploaded files must be posted")
		}

		fmt.Println(r.Header.Get("Content-Type"))
	})

	return http.ListenAndServe(":8080", nil)
}
