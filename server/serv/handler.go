package serv

import (
	"github.com/Morzaka/project-layout/client"
	"github.com/Morzaka/project-layout/server/session"
)

func CreateNewRoom(connection client.Connection) {

	// if you allow creating new room than...
	session.NewRoom(connection, "")
}
