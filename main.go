package main

import "errors"

var board [3][3][3]string

func initialiseBoard() {
	board = [3][3][3]string{
		{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
		{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
		{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
	}
}

var moves int

var directions = make([][3]int, 0)

func Start() {
	initialiseBoard()
	for x := -1; x < 2; x += 1 {
		for y := -1; y < 2; y += 1 {
			for z := -1; z < 2; z += 1 {
				directions = append(directions, [3]int{z, y, x})
			}
		}
	}
}

func isValidPlayer(player string) bool {
	if player == "x" || player == "o" {
		return true
	}

	return false
}

func isValidCoordinate(coordinate [3]int) bool {
	for i := 0; i < 3; i += 1 {
		if coordinate[i] < 0 || coordinate[i] > 2 {
			return false
		}
	}

	return true
}

func travel(currentCoordinate [3]int, direction [3]int, length int) (bool, [3]int) {
	if !isValidCoordinate(currentCoordinate) {
		return false, currentCoordinate
	}

	finalCoordinate := [3]int{0, 0, 0}
	for i := 0; i < 3; i += 1 {
		finalCoordinate[i] = currentCoordinate[i] + length * direction[i]
	}

	return isValidCoordinate(finalCoordinate), finalCoordinate
}

func PlayMove(coordinate [3]int, player string) (bool, error) {
	if !isValidCoordinate(coordinate) {
		return false, errors.New("not a valid coordinate")
	}
	if !isValidPlayer(player) {
		return false, errors.New("not a valid player")
	}
	if moves > 27 {
		return false, errors.New("no more moves left")
	}
	moves += 1
	board[coordinate[2]][coordinate[1]][coordinate[0]] = player

	if moves == 27 {
		return false, nil
	}

	return true, nil
}

func CheckWin(player string) (bool, [3][3]int, error) {
	var empty [3][3]int
	if !isValidPlayer(player) {
		return false, empty, errors.New("not a valid player")
	}
	for x := 0; x < 3; x += 1 {
		for y := 0; y < 3; y += 1 {
			for z := 0; z < 3; z += 1 {
				if board[z][y][x] != player {
					continue
				}
				for _, direction := range directions {
					isValid, firstTravel := travel([3]int{z, y, x}, direction, 1)
					if !isValid || board[firstTravel[2]][firstTravel[1]][firstTravel[0]] != player {
						continue
					}

					isValid, secondTravel := travel(firstTravel, direction, 1)
					if isValid {
						return true, [3][3]int{[3]int{z, y, x}, firstTravel, secondTravel}, nil
					}
				}
			}
		}
	}

	return false, empty, nil
}

func Reset() {
	moves = 0
	initialiseBoard()
}

func main() {
	return
}
