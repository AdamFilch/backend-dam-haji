package logic

import "main/src/common"

var UserColor *string
var OppColor *string

func WhatAmI(user string, fetchedGame common.TableGameStruct) string {
	var color string = "none"
	var opp string = "none"

	// If black player matches your name you are black
	// If black has no players, you are black
	if (fetchedGame.BlackPlayer1Username == user || fetchedGame.BlackPlayer1Username == "") {
		color = "X"
		opp = "0"
	} else if (fetchedGame.WhitePlayer2Username == user || fetchedGame.WhitePlayer2Username == "") {
		color = "0"
		opp = "X"
	}

	OppColor = &opp
	UserColor = &color
	return color
}