package logic

import (
	"log"
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

func CalculateListOfPossibleMoves(currentPosition string, side string) (res []string) {
	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
	split_start_position := re.FindStringSubmatch(currentPosition)

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

func GetAllJumpOverPiece(endPosition, startPosition string, boardState map[string][]string) (res []string) {
	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
	match := re.FindStringSubmatch(endPosition)
	if len(match) < 3 {
		return nil
	}

	row := strings.ToUpper(match[1])
	col, _ := strconv.Atoi(match[2])
	RowInIndex := Character_list_Map[row]

	directions := [][]int{
		{-1, -1}, {-1, 1}, // Top Left, Top Right
		{1, -1}, {1, 1},   // Bottom Left, Bottom Right
	}

	for _, dir := range directions {
		midRowIndex := RowInIndex + dir[0]
		midCol := col + dir[1]
		newRowIndex := RowInIndex + dir[0]*2
		newCol := col + dir[1]*2

		if newRowIndex >= 0 && newRowIndex < len(Character_list_Arr) &&
			newCol >= 1 && newCol <= 10 {

			midRow := Character_list_Arr[midRowIndex]
			newRow := Character_list_Arr[newRowIndex]

			// Ensure middle position has an "X" and landing position is free
			if boardState[midRow][midCol-1] == "0" && (boardState[newRow][newCol-1] != "0" || newRow+strconv.Itoa(newCol) == startPosition) {
				res = append(res, newRow+strconv.Itoa(newCol))
			}
		}
	}
	return res
}

func CalculateMoveStack(endPosition string, startPosition string, boardState map[string][]string) (stack []string) {

	stack = append(stack, endPosition)
	var found bool = false

	for !found {
		res := GetAllJumpOverPiece(stack[len(stack)-1], startPosition, boardState)
		log.Println("CalculateMoveStack", res)

		if true {
			found = true
			break
		}
		found = false
	}

	// List of moves made by the piece; if piece moved 3 locations capturing 3 items than it oculd be [F6, H4, F2]
	// Unless only moved to an adjacent block it would be [E5]
	return stack
}
