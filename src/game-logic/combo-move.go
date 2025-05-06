package logic

import (
	"main/src/utils"
	"regexp"
	"strconv"
)

func ComboMoveAndJump(moveList []string, boardState map[string][]string) (newBoard map[string][]string) {

	newBoard = boardState

	for len(moveList) > 1 {
		startPos := moveList[len(moveList)-1]
		endPos := moveList[len(moveList)-2]
		midPos := utils.GetMiddlePiece(startPos, endPos)
		re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
		match := re.FindStringSubmatch(midPos)

		colIndx, _ := strconv.Atoi(match[2])
		
		newBoard[match[1]][colIndx- 1] = " "

		moveList = moveList[:len(moveList) - 1]
	}

	return newBoard
}