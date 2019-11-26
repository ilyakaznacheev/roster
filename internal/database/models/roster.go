package models

import (
	"math"
	"math/rand"
)

// Roster is a roster document structure
type Roster struct {
	ID      int64 `json:"_id"`
	Players struct {
		Active  []Player
		Benched []Player
	}
	// Version is used internally for CAS
	Version int64
}

// Player is a single player inside a roster
type Player struct {
	ID        int64
	FirstName string
	LastName  string
	Alias     string
}

var maxID = int(math.Pow10(18))

func GenerateID() int {
	return rand.Intn(maxID)
}
