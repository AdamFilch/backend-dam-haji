package utils

import (
	"main/src/common"
	"regexp"
	"strconv"
	"strings"
)

func IsAdjacent(startPosition string, endPosition string) (res bool){
	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
	match := re.FindStringSubmatch(endPosition)
	startRow := strings.ToUpper(match[1])
	startCol, _ := strconv.Atoi(match[2])
	rowInIndex := common.Character_list_Map[startRow]
	row := common.Number_list_arr[startCol-2:startCol+1]

	if rowInIndex-1 >= 0 {
		char := common.Character_list_Arr[rowInIndex-1]


		for _, num := range row {
			if startPosition == char+strconv.Itoa(num) {
				return true
			}
		}
	}

	char := common.Character_list_Arr[rowInIndex]
	for _, num := range row {
		if startPosition == char+strconv.Itoa(num) {
			return true
		}
	}

	if rowInIndex <=9 {
		char := common.Character_list_Arr[rowInIndex+1]
		for _, num := range row {
			if startPosition == char+strconv.Itoa(num) {
				return true
			}
		}
	}



	return false
}