package main

import (
	"log"

	"github.com/apardee/playback/db"
	"github.com/apardee/playback/model"
	"github.com/apardee/playback/service"
)

func main() {
	var fileStore model.PlaybackStore
	fileStore = &db.PlaybackFileStore{}
	if err := fileStore.Open(); err != nil {
		log.Fatal(err)
	}

	{
		clip, err := fileStore.NewClip()
		if err != nil {
			log.Fatal(err)
		}

		clip.Length = 222
		clip.Title = "Hello"
		if err := fileStore.UpdateClip(*clip); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Current Clips:", fileStore.Clips())

	log.Println("Starting web service...")
	if err := service.RunService(&fileStore); err != nil {
		log.Fatal(err)
	}
}
