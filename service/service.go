package service

import (
	"encoding/json"
	"io"
	"net/http"

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

	// http.HandleFunc("/clip/*", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println(r.URL)
	// 	io.WriteString(w, "whoop whoop")
	// })

	// http.HandleFunc("/upload_file", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 		case http.Post
	// 	}
	// })

	return http.ListenAndServe(":8080", nil)
}
