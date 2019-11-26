package database

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
)

const databaseName = "roster"

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
