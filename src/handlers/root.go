package handlers

import (
	"main/src/utils"
	"net/http"
)


type RootServePayload struct {

}

func RootHandle(w http.ResponseWriter, r *http.Request) {


	utils.SubstituteServe(w)
}