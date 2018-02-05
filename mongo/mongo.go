package mongo

import (
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

	dbString := fmt.Sprintf("mongodb://%s:%s@%s/%s?%s", username, password, url, authSource, dbArgs)

	session, err := mgo.Dial(dbString)
	if err != nil {
		return err
	}

	Session = session

	return nil
}
