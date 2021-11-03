package main

import (
	"github.com/Morzaka/project-layout/client"
	"github.com/oleksandr-chornovol/lets-go-chat/pkg/hasher"
)

func main() {
	if hasher.CheckPasswordHash("", "") {
		println("that ok")
	}
	client.NewConnection()
}
