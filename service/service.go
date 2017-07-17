package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/apardee/playback/model"
)

// RunService starts up the http service for serving up Playback objects.
func RunService(store *model.PlaybackStore) error {

	{
		log.Println("1")
		playback, err := model.NewPlaybackState()
		if err != nil {
			return err
		}

		// log.Println("2")
		// uuid, err := uuid.NewV4()
		// if err != nil {
		// 	return err
		// }

		// playback.ClipID = uuid
		// playback.Location = 123

		log.Println("3")
		bytes, err := json.Marshal(playback)
		if err != nil {
			return err
		}

		fmt.Println("encoded:", string(bytes))

		// uuid, err := uuid.NewV4()
		// if err != nil {
		// 	return err
		// }
	}

	return http.ListenAndServe(":8080", nil)
}
