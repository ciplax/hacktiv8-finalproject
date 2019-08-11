package sessions

import (
	"math/rand"
)

//Sessions stores session, simple session mockup
type Sessions struct {
	SessionID int
	Username  string
}

//Sess public var of Sessions
var Sess Sessions

//GenerateSessions generates new session
func (sess *Sessions) GenerateSessions() {
	sess.SessionID = rand.Intn(100)
}

//ClearSessions returns empty Sessions
func (sess *Sessions) ClearSessions() {
	*sess = Sessions{}
}

//InsertUsernameToSessions put username to sessions
func (sess *Sessions) InsertUsernameToSessions(username string) {
	sess.Username = *&username
}
