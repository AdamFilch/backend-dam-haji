package logic

import (
	"regexp"
	"strconv"
	"strings"
)

var Character_list_Map = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
	"I": 8,
	"J": 9,
}

var Character_list_Arr = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
}

func CalculateListOfPossibleMoves(move string, side string) (res []string) {
	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
	split_start_position := re.FindStringSubmatch(move)

	// Convert the row part (second character) to an integer

	column, _ := strconv.Atoi(split_start_position[2])
	row := split_start_position[1]

	res = GetAllAdjecentSpaces(column, strings.ToUpper(row))

	return res
}

func GetAllAdjecentSpaces(column int, row string) (res []string) {

	RowInIndex := Character_list_Map[row]

	if RowInIndex != 0 {
		LeftTop := Character_list_Arr[RowInIndex-1] + strconv.Itoa(column-1)
		RigthTop := Character_list_Arr[RowInIndex-1] + strconv.Itoa(column+1)

		res = append(res, LeftTop, RigthTop)

	}

	if RowInIndex != 9 {
		BottomLeft := Character_list_Arr[RowInIndex+1] + strconv.Itoa(column-1)
		BottomRight := Character_list_Arr[RowInIndex+1] + strconv.Itoa(column+1)
		res = append(res, BottomLeft, BottomRight)

	}

	return res
}
