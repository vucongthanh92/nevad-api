package common

import "log"

const (
	DbTypeUser    = 1
	DbTypeTeam    = 2
	DbTypeProxy   = 3
	DbTypeProfile = 4
)

const (
	CurrentUser = "user"
	CurrentTeam = "Team"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
