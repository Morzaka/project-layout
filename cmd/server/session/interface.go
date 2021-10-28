package session

type SessionMaker interface {
	Join()
	Send()
	Leave()
	Broadcast()
}
