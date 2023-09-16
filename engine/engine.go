package engine

import (
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/shadow"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/standard"
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art/thinkertoy"
	"fmt"
	"os"
)

func Start() {
	var userInput string
	var userTrigger string
	if len(os.Args) > 2 && len(os.Args) < 4 {
		userInput = os.Args[1]
		if len(os.Args) == 2 {
			userTrigger = "standard"
		} else {
			userTrigger = os.Args[2]
		}
		if userTrigger == "shadow" {
			shadow.Shadow(userInput)
		} else if userTrigger == "standard" {
			standard.Standard(userInput)
		} else if userTrigger == "thinkertoy" {
			thinkertoy.Thinkertoy(userInput)
		}
	} else {
		fmt.Println("Please enter an argument with or without a font trigger")
		return
	}
}
