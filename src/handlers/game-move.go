package handlers

import (
	"log"
	"main/src/common"
	"main/src/db"
	logic "main/src/game-logic"
	"main/src/utils"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type newMoveStruct struct {
	GameID        string    `json:"game_id"`
	Username      string    `json:"username"`
	StartPosition string    `json:"move_from"`
	EndPosition   string    `json:"move_to"`
	CreatedAt     time.Time `json:"created_at"`
	PieceColor    string    `json:"piece_color"`
}

type playerMovePayload struct {
	GameID     string                    `json:"gameId"`
	Players    map[string]BasePlayerProp `json:"players"`
	BoardState map[string][]string       `json:"board_state"`
	Data       map[string]string         `json:"data"`
}

type BasePlayerProp struct {
	Points int     `json:"points"`
	Letter string  `json:"letter"`
	Winner *string `json:"winner,omitempty"`
}

type updateGameWithBoardStruct struct {
	WhitePlayer2Username string              `json:"white_player2_username"`
	UpdatedAt            time.Time           `json:"updated_at"`
	BoardState           map[string][]string `json:"board_state"`
}

func isValidPosition(position string) bool {
	// Define regex pattern: One letter (A-J) followed by one or two digits (1-10)
	pattern := `^[A-Ja-j](10|[1-9])$`
	matched, _ := regexp.MatchString(pattern, position)
	return matched
}

func HandleGameMove(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	gameID := vars["gameid"]
	user := vars["user"]
	start_position := vars["start"]
	end_position := vars["end"]

	var err error
	var res any

	var fetchedGame []common.TableGameStruct
	err = db.SupaClient.DB.From("games").Select("*").Eq("game_id_pk", gameID).Execute(&fetchedGame)
	if err != nil {
		log.Println("Error: HandleGameMove - Unable to fetch from Games_T", err)
	}
	if len(fetchedGame) == 0 {
		additionalData := map[string]string{
			"error":      "Oops!",
			"message":    "It seems like this game does not exists within our system!",
			"start-game": "Please create a new game by using: " + r.Host + `/start-game/` + user,
		}

		utils.Serve(w, additionalData)
		return
	}

	additionalData := map[string]string{

		"how_to_play":  r.Host + `/how-to-play`,
		"your_piece":   "Black",
		"turn":         user,
		"last_updated": utils.FormatTimestamp(time.Now().UTC()),
	}

	var fetchedMoves []common.TableMovesStruct
	err = db.SupaClient.DB.From("moves").Select("*").OrderBy("created_at", "desc").Eq("game_id", gameID).Execute(&fetchedMoves)
	if err != nil {
		log.Println("Error: HandleGameMove - Fetching from Moves_T", err)
	}

	if len(fetchedMoves) == 0 {
		// User was the first to make the move
		additionalData["playback"] = user + " has just made the first move of the game from " + start_position + " to " + end_position
	} else {
		additionalData["playback"] = fetchedMoves[0].Username + " has just moved their piece from " + fetchedMoves[0].StartPosition + " to " + fetchedMoves[0].EndPosition
		additionalData["last_updated"] = utils.FormatTimestamp(fetchedGame[0].UpdatedAt)
	}

	p := playerMovePayload{
		GameID:     gameID,
		BoardState: fetchedGame[0].BoardState,
		Data:       additionalData,
		Players:    make(map[string]BasePlayerProp), // Initialize the prop
	}

	p.Players[user] = BasePlayerProp{
		Points: 0,
		Letter: "Black",
	}
	p.Players[fetchedGame[0].BlackPlayer1Username] = BasePlayerProp{
		Points: 0,
		Letter: "White",
	}

	if !isValidPosition(end_position) {
		additionalData["error"] = "Excuse me, where do you think you're going? hat final position is not valid"
		utils.Serve(w, p)
		return
	}
	if !isValidPosition(start_position) {
		additionalData["error"] = "Ummm, whatever you have just tried to move, was not valid"
		utils.Serve(w, p)
		return
	}

	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)

	split_start_position := re.FindStringSubmatch(start_position) // Creates an array like [A8 A 8] [B9 B 9]
	split_end_position := re.FindStringSubmatch(end_position)

	// Convert the row part (second character) to an integer
	start_row, _ := strconv.Atoi(split_start_position[2])
	end_row, _ := strconv.Atoi(split_end_position[2])

	// Access the board correctly
	if fetchedGame[0].BoardState[strings.ToUpper(split_start_position[1])][start_row-1] == "X" {

		if fetchedGame[0].BoardState[strings.ToUpper(split_end_position[1])][end_row-1] == " " {
			// Move logic here
			adjecent_tiles := logic.CalculateListOfPossibleMoves(start_position, "black")

			log.Println("Does it contain", utils.Contains(adjecent_tiles, end_position))

			if utils.Contains(adjecent_tiles, strings.ToUpper(end_position)) != -1 {
				// Ensure proper 0-based indexing in assignment
				p.BoardState[strings.ToUpper(split_start_position[1])][start_row-1] = " "
				p.BoardState[strings.ToUpper(split_end_position[1])][end_row-1] = "X"
			} else {
				additionalData["error"] = "Unfortunately thats not how to play checkers."
				utils.Serve(w, p)
				return
			}

		}
	}

	if fetchedGame[0].BlackPlayer1Username != "" && fetchedGame[0].WhitePlayer2Username != "" {
		if fetchedGame[0].BlackPlayer1Username != user && fetchedGame[0].WhitePlayer2Username != user {
			additionalData := map[string]string{
				"start-game": "Start your own game by using: " + r.Host + `/start-game/` + user,
				"error":      "Unfortunately this game already has 2 players playing, " + fetchedGame[0].BlackPlayer1Username + " and " + fetchedGame[0].WhitePlayer2Username,
			}

			p := alreadyMatchedPayload{
				GameID:               gameID,
				BlackPlayer2Username: fetchedGame[0].BlackPlayer1Username,
				WhitePlayer2Username: fetchedGame[0].WhitePlayer2Username,
				Data:                 additionalData,
			}
			utils.Serve(w, p)
			return
		}
	}

	newUser := newUserStruct{
		Username:    user,
		TotalPoints: 0,
		GamesWon:    0,
		CreatedAt:   time.Now().UTC(),
	}

	var existingUser []newUserStruct
	err = db.SupaClient.DB.From("users").Select("*").Eq("username", user).Execute(&existingUser)
	if err != nil {
		log.Println("Error: HandleInitGame - Fetching from users_t: ", err)
	}
	if len(existingUser) == 0 {
		err = db.SupaClient.DB.From("users").Insert(newUser).Execute(&res)
		if err != nil {
			log.Println("Error: HandleInitGame - Inserting to users_t: ", err)
		}
	}
	updatedGame := updateGameWithBoardStruct{
		WhitePlayer2Username: user,
		UpdatedAt:            time.Now().UTC(),
		BoardState:           p.BoardState,
		// Add the new board state move here
	}

	// Update the game in Database
	err = db.SupaClient.DB.From("games").Update(updatedGame).Eq("game_id_pk", gameID).Execute(&res)
	if err != nil {
		log.Println("Error: HandleGameGetGame - Updating from games_t", err)
	}

	// If Eerything is okay the move will be made and
	newMove := newMoveStruct{
		GameID:        gameID,
		Username:      user,
		StartPosition: start_position,
		EndPosition:   end_position,
		CreatedAt:     time.Now().UTC(),
		PieceColor:    "black",
	}
	err = db.SupaClient.DB.From("moves").Insert(newMove).Execute(&res)
	if err != nil {
		log.Println("Error: HandleGameMove - Inserting to Moves_T: ", err)
	}

	utils.Serve(w, p)
}
