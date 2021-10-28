package main

import (
	"fmt"
	_ "github.com/Morzaka/project-layout/cmd/server/model"

	. "github.com/Morzaka/project-layout/cmd/server/serv/service_hand/greeting"
)

func main() {

	hand := NewHand("left")
	//hand.Greet()

	if hand.HandLeve == High && hand.Action == "waving" {
		fmt.Println("Hello, hello ... I wave to you.")
	} else {
		fmt.Println("No greeting")
	}

	//println("Hello new project")
}
