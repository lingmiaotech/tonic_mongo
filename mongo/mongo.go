package mongo

import (
	"errors"
	"github.com/lingmiaotech/tonic_mongo/configs"
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

func InitDatabase() (err error) {
	isDatabaseSet := configs.IsSet("database")
	if !isDatabaseSet {
		return nil
	}

	driver := configs.GetString("database.driver")
	if driver != "mongo" {
		return errors.New("tonic_mongo_error.database.driver_not_supported")
	}

	dbstring := configs.GetString("database.dbstring")
	if dbstring == "" {
		return errors.New("tonic_mongo_error.database.empty_dbstring_config")
	}

	session, err := mgo.Dial(dbstring)
	if err != nil {
		return err
	}

	Session = session

	return nil
}
