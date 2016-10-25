package zeleniy

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
)

var ErrorInvalidObjectId = errors.New("Invalid object id")


func ObjectId(id string) (bson.ObjectId, error) {
	if !bson.IsObjectIdHex(id) {
		return bson.ObjectId(""), ErrorInvalidObjectId
	}
	return bson.ObjectIdHex(id), nil
}


func checkErrNotFound(err error, msg string) error {
	if err == mgo.ErrNotFound {
		return nil, err
	}
	return nil, errors.Wrapf(err, msg)
}
