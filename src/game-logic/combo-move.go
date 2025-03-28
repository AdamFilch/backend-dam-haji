package logic

import (
	"log"
	"main/src/utils"
)

func ComboMoveAndJump(moveList []string, boardState map[string][]string) (newBoard map[string][]string) {

	newBoard = boardState


	for len(moveList) > 1 {
		startPos := moveList[len(moveList)-1]
		endPos := moveList[len(moveList)-2]
		midPos := utils.GetMiddlePiece(startPos, endPos)
		log.Println(startPos, endPos, midPos)

		moveList = moveList[:len(moveList) - 1]

	}

	return newBoard
}