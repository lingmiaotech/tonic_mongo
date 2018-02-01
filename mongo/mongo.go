package mongo

import (
	"errors"
	"github.com/lingmiaotech/tonic_mongo/configs"
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

func InitMongoDB() (err error) {
	isMongoDBSet := configs.IsSet("mongodb")
	if !isMongoDBSet {
		return nil
	}

	dbstring := configs.GetString("mongodb.dbstring")
	if dbstring == "" {
		return errors.New("tonic_mongo_error.empty_dbstring_config")
	}

	session, err := mgo.Dial(dbstring)
	if err != nil {
		return err
	}

	Session = session

	return nil
}
