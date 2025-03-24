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

func GetAllJumpOverPiece(endPosition string, startPosition string, boardState map[string][]string) (res []string, found int) {
	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
	splitEndPosition := re.FindStringSubmatch(endPosition)

	column, _ := strconv.Atoi(splitEndPosition[2])
	row := strings.ToUpper(splitEndPosition[1])

	RowInIndex := Character_list_Map[row]

	var mid []string

	if RowInIndex > 1 {
		
		// Top Left
		mid = re.FindStringSubmatch(Character_list_Arr[RowInIndex-1] + strconv.Itoa(column-1))
		LTMColumn, _ := strconv.Atoi(mid[2])

		if boardState[mid[1]][LTMColumn-1] == "X" {

			ft := re.FindStringSubmatch(Character_list_Arr[RowInIndex-2] + strconv.Itoa(column-2))
			col, _ := strconv.Atoi(ft[2])

			if boardState[ft[1]][col-1] != "X" || ft[0] == startPosition {
				concat := ft[0]
				res = append(res, concat)
				if ft[0] == startPosition {
					found = 1
				}
			}

		}


		// Top Right
		mid = re.FindStringSubmatch(Character_list_Arr[RowInIndex-1] + strconv.Itoa(column+1))
		RTMColumn, _ := strconv.Atoi(mid[2])

		if boardState[mid[1]][RTMColumn-1] == "X" { // Is Occupied
			
			ft := re.FindStringSubmatch(Character_list_Arr[RowInIndex-2] + strconv.Itoa(column+2))
			col, _ := strconv.Atoi(ft[2])

			if boardState[ft[1]][col-1] != "X"  || ft[0] == startPosition { // is not occupied or the position is the start position
				concat := ft[0]
				res = append(res, concat)
				if ft[0] == startPosition {
					found = 1
				}
			}
		}
	}

	if RowInIndex < 8 {

		var mid []string


		// Bottom Left
		mid = re.FindStringSubmatch(Character_list_Arr[RowInIndex+1] + strconv.Itoa(column-1))
		LBMColumn, _ := strconv.Atoi(mid[2])

		if boardState[mid[1]][LBMColumn-1] == "X" {
			
			ft := re.FindStringSubmatch(Character_list_Arr[RowInIndex+2] + strconv.Itoa(column-2))
			col, _ := strconv.Atoi(ft[2])

			if boardState[ft[1]][col-1] != "X" || ft[0] == startPosition {
				concat := ft[0]
				res = append(res, concat)
				if ft[0] == startPosition {
					found = 1
				}
			}
		}

		// Bottom Right
		mid = re.FindStringSubmatch(Character_list_Arr[RowInIndex+1] + strconv.Itoa(column+1))
		RBMColumn, _ := strconv.Atoi(mid[2])

		if boardState[mid[1]][RBMColumn-1] == "X" {
			
			ft := re.FindStringSubmatch(Character_list_Arr[RowInIndex+2] + strconv.Itoa(column+2))
			col, _ := strconv.Atoi(ft[2])

			if boardState[ft[1]][col-1] != "X" || ft[0] == startPosition { 
				concat := ft[0]
				res = append(res, concat)
				if ft[0] == startPosition {
					found = 1
				}
			}
		}
	}

	return res, found
}

func CalculateMoveStack(endPosition string, startPosition string, boardState map[string][]string) (stack []string) {

	stack = append(stack, endPosition)
	var found bool = false

	for !found {
		res, foundRes := GetAllJumpOverPiece(stack[len(stack)-1], startPosition, boardState)
		log.Println("CalculateMoveStack", res, foundRes)

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
