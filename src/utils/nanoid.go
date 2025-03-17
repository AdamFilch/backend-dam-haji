package utils

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateNanoID() string {
	shortID, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 6)
	if err != nil {
		log.Fatal(err)
	}

	return shortID
}
