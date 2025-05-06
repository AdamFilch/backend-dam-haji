package handlers

import (
	"log"
	"main/src/common"
	"main/src/db"
	"main/src/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type SetGameStruct struct {
	GameID   string              `json:"gameId"`
	NewBoard map[string][]string `json:"new_board"`
	OldBoard map[string][]string `json:"old_board"`
	Players  []string            `json:"current_players"`
}

type UpdateBoardStruct struct {
	UpdatedAt  time.Time           `json:"updated_at"`
	BoardState map[string][]string `json:"board_state"`
}

func SetGameBoard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	gameId := vars["gameid"]
	templateNum := vars["templatenum"]
	temp, _ := strconv.Atoi(templateNum)

	var err error
	var res any

	if temp > len(common.Board_list_map) {
		additionalData := map[string]string{
			"error":   "Oops!",
			"message": "It seems like we do not have a board mapped to that number!",
		}

		utils.Serve(w, additionalData)
		return
	}

	var fetchedGame []common.TableGameStruct
	err = db.SupaClient.DB.From("games").Select("*").Eq("game_id_pk", gameId).Execute(&fetchedGame)
	if err != nil {
		log.Println("Error: SetGameBoard - Fetching from Games_T: ", err)
	}
	if len(fetchedGame) == 0 {
		additionalData := map[string]string{
			"error":      "Oops!",
			"message":    "It seems like this game does not exists within our system!",
			"start-game": "Please create a new game by using: " + r.Host + `/start-game/` + `username-of-your-choice`,
		}

		utils.Serve(w, additionalData)
		return
	}

	updatedGame := UpdateBoardStruct{
		UpdatedAt:  time.Now().UTC(),
		BoardState: common.Board_list_map[temp],
	}

	err = db.SupaClient.DB.From("games").Update(updatedGame).Eq("game_id_pk", gameId).Execute(&res)
	if err != nil {
		log.Println("Error: SetGameBoard - Updating from games_t", err)
	}

	p := SetGameStruct{
		GameID:   gameId,
		NewBoard: common.Board_list_map[temp],
		OldBoard: fetchedGame[0].BoardState,
		Players:  []string{fetchedGame[0].BlackPlayer1Username, fetchedGame[0].WhitePlayer2Username},
	}

	utils.Serve(w, p)
}
