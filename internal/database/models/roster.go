package models

import (
	"math"
	"math/rand"
)

// Roster is a roster document structure
type Roster struct {
	ID      int64 `bson:"_id"`
	Players struct {
		Active  []Player
		Benched []Player
	}
	// Version is used internally for CAS
	Version int64
}

// Player is a single player inside a roster
type Player struct {
	ID        int64  `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Alias     string `bson:"alias"`
}

var maxID = int(math.Pow10(18))

func GenerateID() int64 {
	return int64(rand.Intn(maxID))
}
