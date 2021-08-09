package main

import (
	"Mentor_prg/project-layout/services/service_hand/greeting"
	"fmt"
)

func main() {

	hand := greeting.NewHand("left")
	//hand.Greet()

	if hand.HandLeve == greeting.High && hand.Action == "waving" {
		fmt.Println("Hello, hello ... I wave to you.")
	} else {
		fmt.Println("No greeting")
	}

	//println("Hello new project")
}
