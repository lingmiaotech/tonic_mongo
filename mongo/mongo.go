package mongo

import (
	"errors"
	"fmt"

	"github.com/lingmiaotech/tonic_mongo/configs"
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

func InitMongoDB() (err error) {
	isMongoDBSet := configs.IsSet("mongodb")
	if !isMongoDBSet {
		return nil
	}

	username := configs.GetString("mongodb.username")
	password := configs.GetDynamicString("mongodb.password")
	url := configs.GetString("mongodb.url")
	authSource := configs.GetString("mongodb.authSource")
	dbArgs := configs.GetString("mongodb.args")

	dbString := fmt.Sprintf("mongodb://")
	if username != "" && password != "" {
		dbString = fmt.Sprintf("%s%s:%s@", dbString, username, password)
	}

	if url != "" {
		dbString = fmt.Sprintf("%s%s", dbString, url)
	} else {
		return errors.New("mongo_error.empty_url")
	}

	if authSource != "" {
		dbString = fmt.Sprintf("%s/%s", dbString, authSource)
	}

	if dbArgs != "" {
		dbString = fmt.Sprintf("%s?%s", dbString, dbArgs)
	}

	session, err := mgo.Dial(dbString)
	if err != nil {
		return err
	}

	Session = session

	return nil
}
