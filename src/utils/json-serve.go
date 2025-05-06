package utils

import (
	"encoding/json"
	"net/http"
)

func Serve(w http.ResponseWriter, p any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(p)
}

type SubstituteServeResponse struct {
	Title              string            `json:"title"`
	CreatedBy          []string          `json:"created-by"`
	About              map[string]string `json:"about"`
}

func SubstituteServe(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	aboutSection := map[string]string{
		"1": "This is a backend server made by Adam and Nazmi to play Cherkers using REST",
		"2": "We created this mainly to learn how to use Golang and to learn backend development",
		"3": "This is not a project that contributes to anyone in particular,",
		"4": "It was instead our own passion of wanting to learn Golang and Backend Dev",
	}

	p := SubstituteServeResponse{
		About:              aboutSection,
		Title:              "Welcome to Backend Checkers also known as Dam Haji (Checkers in malay)",
		CreatedBy:          []string{"Adam", "Nazmi"},
	}

	json.NewEncoder(w).Encode(p)
}

