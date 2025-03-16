package handlers

import (
	"log"
	"net/http"
)

func HowToUse(w http.ResponseWriter, r *http.Request) {
	log.Println("How TO use this program")
}
