package models

import (
	"math"
	"math/rand"
)

// Roster is a roster document structure
type Roster struct {
	ID      int64 `bson:"_id"`
	Players RosterPlayers
	// Version is used internally for CAS
	Version int64
}

// RosterPlayers is a two sets of players in the same roster: active and benched
type RosterPlayers struct {
	Active  []Player
	Benched []Player
}

// Player is a single player inside a roster
type Player struct {
	ID        int64      `bson:"id"`
	FirstName string     `bson:"first_name"`
	LastName  string     `bson:"last_name"`
	Alias     string     `bson:"alias"`
	Role      PlayerRole `bson:"role"`
}

var maxID = int(math.Pow10(18))

// GenerateID creates a random 18-digit ID
func GenerateID() int64 {
	return int64(rand.Intn(maxID))
}

// PlayerRole is a player role
type PlayerRole string

const (
	RoleRifler  PlayerRole = "rifler"
	RoleIgl     PlayerRole = "igl"
	RoleSupport PlayerRole = "support"
	RoleAwper   PlayerRole = "awper"
)
