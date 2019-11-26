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
	mgd *mgo.Database
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
		mgd: session.DB(databaseName),
	}
	return &mh, nil

}

func (h *MongoHandler) GetAllRosters() ([]models.Roster, error) {
	var res []models.Roster
	err := h.mgd.C(entityRoster).Find(bson.M{}).All(&res)
	return res, err
}
func (h *MongoHandler) GetRoster(id int) (*models.Roster, error) {
	var res *models.Roster
	err := h.mgd.C(entityRoster).Find(bson.M{"id": id}).One(res)
	if err == mgo.ErrNotFound {
		return res, NewNotFoundError(err)
	}
	return res, err
}
func (h *MongoHandler) UpdateRoster(r models.Roster) error {
	var res *models.Roster
	err := h.mgd.C(entityRoster).Find(bson.M{"id": r.ID}).One(res)
	if err == mgo.ErrNotFound {
		return NewNotFoundError(err)
	} else if err != nil {
		return err
	}

	h.mgd.C(entityRoster).
		h.mgd.C(entityRoster).UpsertId(r.ID)

	return nil
}
