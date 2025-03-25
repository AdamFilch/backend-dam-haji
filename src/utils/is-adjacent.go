package utils

import (
	"main/src/common"
	"regexp"
	"strconv"
	"strings"
)

func IsAdjacent(startPosition string, endPosition string) (res bool){
	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
	match := re.FindStringSubmatch(startPosition)
	startRow := strings.ToUpper(match[1])
	startCol, _ := strconv.Atoi(match[2])
	startRowIndex := common.Character_list_Map[startRow]

	// row := common.Number_list_arr[startCol-2:startCol+1]

	possibleMoves := [][]int{
		{startRowIndex - 1, startCol - 1}, {startRowIndex - 1, startCol}, {startRowIndex - 1, startCol + 1}, // Top row
		{startRowIndex, startCol - 1}, /* Current Row */ {startRowIndex, startCol + 1}, // Left & Right
		{startRowIndex + 1, startCol - 1}, {startRowIndex + 1, startCol}, {startRowIndex + 1, startCol + 1}, // Bottom row
	}

	for _, move := range possibleMoves {
		rowIndex, col := move[0], move[1]

		if rowIndex >= 0 && rowIndex < 10 && col >=1 && col <= 10 {
			pos := common.Character_list_Arr[rowIndex] + strconv.Itoa(col)
			if endPosition ==  pos{
				return true
			}
		}
	}


	// if rowInIndex-1 >= 0 {
	// 	char := common.Character_list_Arr[rowInIndex-1]

	// 	for _, num := range row {
	// 		if startPosition == char+strconv.Itoa(num) {
	// 			return true
	// 		}
	// 	}
	// }

	// char := common.Character_list_Arr[rowInIndex]
	// for _, num := range row {
	// 	if startPosition == char+strconv.Itoa(num) {
	// 		return true
	// 	}
	// }

	// if rowInIndex <=9 {
	// 	char := common.Character_list_Arr[rowInIndex+1]
	// 	for _, num := range row {
	// 		if startPosition == char+strconv.Itoa(num) {
	// 			return true
	// 		}
	// 	}
	// }
	return false
}