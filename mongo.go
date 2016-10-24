package zeleniy

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
)

var ErrorInvalidObjectId = errors.New("Invalid object id")


func ObjectId(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(id) {
		return bson.ObjectId(""), ErrorInvalidObjectId
	}
	return bson.ObjectIdHex(id), nil
}
