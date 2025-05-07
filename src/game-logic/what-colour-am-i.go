package logic

import "main/src/common"

var UserColor *string
var OppColor *string

func WhatAmI(user string, fetchedGame common.TableGameStruct) string {
	var color string = "none"
	var opp string = "none"

	if (fetchedGame.BlackPlayer1Username == user) {
		color = "X"
		opp = "0"
	} else if (fetchedGame.WhitePlayer2Username == user) {
		color = "0"
		opp = "X"
	}

	OppColor = &opp
	UserColor = &color
	return color
}