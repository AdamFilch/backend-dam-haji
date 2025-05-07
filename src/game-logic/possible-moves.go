package logic

import (
	"main/src/common"
	"regexp"
	"strconv"
	"strings"
)

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

	RowInIndex := common.Character_list_Map[row]

	if RowInIndex != 0 {
		LeftTop := common.Character_list_Arr[RowInIndex-1] + strconv.Itoa(column-1)
		RigthTop := common.Character_list_Arr[RowInIndex-1] + strconv.Itoa(column+1)
		res = append(res, LeftTop, RigthTop)
	}

	if RowInIndex != 9 {
		BottomLeft := common.Character_list_Arr[RowInIndex+1] + strconv.Itoa(column-1)
		BottomRight := common.Character_list_Arr[RowInIndex+1] + strconv.Itoa(column+1)
		res = append(res, BottomLeft, BottomRight)
	}
	return res
}

// Finds all valid jumps
func GetAllJumpOverPiece(endPosition, startPosition string, boardState map[string][]string) (res []string) {
	re := regexp.MustCompile(`([A-Ja-j]+)(\d+)`)
	match := re.FindStringSubmatch(endPosition)
	if len(match) < 3 {
		return nil
	}

	row := strings.ToUpper(match[1])
	col, _ := strconv.Atoi(match[2])
	RowInIndex := common.Character_list_Map[row]

	directions := [][]int{
		{-1, -1}, {-1, 1}, // Top Left, Top Right
		{1, -1}, {1, 1}, // Bottom Left, Bottom Right
	}

	for _, dir := range directions {
		midRowIndex := RowInIndex + dir[0]
		midCol := col + dir[1]
		newRowIndex := RowInIndex + dir[0]*2
		newCol := col + dir[1]*2

		// Ensure new move is within board boundaries
		if newRowIndex >= 0 && newRowIndex < len(common.Character_list_Arr) &&
			newCol >= 1 && newCol <= 10 {

			midRow := common.Character_list_Arr[midRowIndex]
			newRow := common.Character_list_Arr[newRowIndex]

			// Ensure middle position has an "X" and landing position is free
			if boardState[midRow][midCol-1] == *OppColor && (boardState[newRow][newCol-1] == " " || newRow+strconv.Itoa(newCol) == startPosition) {
				res = append(res, newRow+strconv.Itoa(newCol))
			}
		}
	}
	return res
}

// Recursive function to find shortest jump sequence
func FindShortestPath(currentPos, startPos string, boardState map[string][]string, path []string, minPath *[]string, visited map[string]bool) {
	if currentPos == startPos {
		// If the new path is shorter, update minPath
		if len(*minPath) == 0 || len(path) < len(*minPath) {
			*minPath = append([]string{}, path...)
		}
		return
	}

	if visited[currentPos] {
		return
	}

	visited[currentPos] = true

	nextMoves := GetAllJumpOverPiece(currentPos, startPos, boardState)
	for _, move := range nextMoves {
		FindShortestPath(move, startPos, boardState, append(path, move), minPath, visited)
	}

	visited[currentPos] = false
}

// Wrapper function
func CalculateMoveStack(endPosition, startPosition string, boardState map[string][]string) []string {
	var shortestPath []string
	visited := make(map[string]bool)

	FindShortestPath(endPosition, startPosition, boardState, []string{endPosition}, &shortestPath, visited)

	return shortestPath
}
