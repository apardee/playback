package main

import "github.com/apardee/playback/db"
import "log"

func main() {
	fileStore := db.PlaybackFileStore{}
	if err := fileStore.Open(); err != nil {
		log.Fatal(err)
	}

	log.Println("clips:", fileStore.Clips())

	clip, err := fileStore.NewClip()
	if err != nil {
		log.Fatal(err)
	}

	clip.Length = 222
	clip.Title = "Hello"
	if err := fileStore.UpdateClip(*clip); err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}
