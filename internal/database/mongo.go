package database

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/ilyakaznacheev/roster/internal/database/models"
)

const databaseName = "roster"
const entityRoster = "roster"

type MongoHandler struct {
	db *mgo.Database
}

// NewMongoHandler connects to a MongoDB and creates database handler
func NewMongoHandler(conURI string) (*MongoHandler, error) {
	var (
		session *mgo.Session
		err     error
	)

	// try to connect to MongoDB
	log.Printf("connecting to MongoDB...")
	session, err = mgo.Dial(conURI)
	if err != nil {
		return nil, fmt.Errorf("MongoDB connection error: %w", err)
	}
	log.Printf("MongoDB connection established")

	mh := MongoHandler{
		db: session.DB(databaseName),
	}
	return &mh, nil

}

func (h *MongoHandler) GetAllRosters() ([]models.Roster, error) {
	var res []models.Roster
	err := h.db.C(entityRoster).Find(bson.M{}).All(&res)
	return res, err
}
func (h *MongoHandler) GetRoster(id int64) (*models.Roster, error) {
	var res models.Roster
	err := h.db.C(entityRoster).FindId(id).One(&res)
	if err == mgo.ErrNotFound {
		return nil, NewNotFoundError(err)
	}
	return &res, err
}
func (h *MongoHandler) UpdateRoster(r models.Roster) error {
	id := r.ID
	version := r.Version
	r.Version++

	err := h.db.C(entityRoster).Update(
		bson.M{"_id": id, "version": version},
		r,
	)
	if err == mgo.ErrNotFound {
		return NewNotFoundError(err)
	}
	return err
}

func (h *MongoHandler) PushPlayer(id int64, p models.Player) error {

	err := h.db.C(entityRoster).Update(
		bson.M{"_id": id},
		bson.M{
			"$inc":  bson.M{"version": 1},
			"$push": bson.M{"players.benched": p},
		},
	)
	if err == mgo.ErrNotFound {
		return NewNotFoundError(err)
	}
	return err
}
