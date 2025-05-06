package handlers

import (
	"fmt"
	"main/src/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func SetGameBoard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	gameId := vars["gameid"]
	templateNum := vars["templatenum"]

	fmt.Println(gameId, templateNum)

	utils.SubstituteServe(w)
}
