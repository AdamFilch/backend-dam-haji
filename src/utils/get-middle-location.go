package utils

import (
	"main/src/common"
	"regexp"
	"strconv"
	"strings"
)

func GetMiddlePiece(startPosition string, endPosition string) (mid string) {

	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
	st := re.FindStringSubmatch(startPosition)
	ed := re.FindStringSubmatch(endPosition)

	startColIndex, _ := strconv.Atoi(st[2])
	startRow := strings.ToUpper(st[1])
	startRowIndex := common.Character_list_Map[startRow]

	endColIndex, _ := strconv.Atoi(ed[2])
	endRow := strings.ToUpper(ed[1])
	endRowIndex := common.Character_list_Map[endRow]

	var midRowIndex int
	var midColIndex int

	if startRowIndex < endRowIndex {
		midRowIndex = endRowIndex - 1
	} else {
		midRowIndex = startRowIndex - 1
	}

	if startColIndex < endColIndex {
		midColIndex = endColIndex - 1
	} else {
		midColIndex = startColIndex - 1
	}

	mid = common.Character_list_Arr[midRowIndex] + strconv.Itoa(midColIndex)



	return mid
}


