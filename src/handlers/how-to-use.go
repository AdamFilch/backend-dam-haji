package handlers

import (
	"main/src/utils"
	"net/http"
)

type HowToUsePayload struct {
	Title              string            `json:"title"`
	CreatedBy          []string          `json:"created_by"`
	WhereToNext        string            `json:"where_to_next"`
	About              map[string]string `json:"about"`
	AvailableEndpoints map[string]string `json:"available_endpoints"`
}

func HowToUse(w http.ResponseWriter, r *http.Request) {

	aboutSection := map[string]string{
		"1": "This is a backend server made by Adam and Nazmi to play Cherkers using REST",
		"2": "We created this mainly to learn how to use Golang and to learn backend development",
		"3": "This is not a project that contributes to anyone in particular,",
		"4": "It was instead our own passion of wanting to learn Golang and Backend Dev",
	}

	availableEndpointsSection := map[string]string{
		"1": "/how-to-use",
		"2": "/learn-checkers",
		"3": "/start-game/{user}",
		"4": "/{gameid}",
		"5": "/{gameid}/{user}",
		"6": "/{gameid}/{user}/move/{start}/to/{end}",
		"7": "/history/{gameid}",
		"8": "/leaderboard",
	}

	p := HowToUsePayload{
		AvailableEndpoints: availableEndpointsSection,
		About:              aboutSection,
		Title:              "Welcome to Backend Checkers also known as Dam Haji (Checkers in malay)",
		WhereToNext:        r.Host + `/learn-checkers`,
		CreatedBy:          []string{"Adam", "Nazmi"},
	}


	utils.Serve(w, p)
}
