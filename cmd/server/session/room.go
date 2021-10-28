package session

import (
	"github.com/Morzaka/project-layout/client"
	"time"
)

type Room struct {
	Connection   client.Connection
	Name         string
	CreationData time.Time
	Participant  []string
}

func NewRoom(c client.Connection, name string) SessionMaker {
	return &Room{
		Connection:   c,
		Name:         name,
		CreationData: time.Now(),
		Participant:  nil,
	}
}

func (r *Room) Join() {}

func (r *Room) Leave() {}

func (r *Room) Send() {}

func (r *Room) Broadcast() {}
