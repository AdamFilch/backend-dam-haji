package logic

import (
	"regexp"
)

func MakeMeKing(currentPosition string, userColor string) bool {

	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
	split_start_position := re.FindStringSubmatch(currentPosition)


	row := split_start_position[1]

	if row == "A" && userColor == "0" {
		return true
	} else if row == "J" && userColor == "X" {
		return true
	}

	return false
}