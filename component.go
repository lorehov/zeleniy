package zeleniy

import "gopkg.in/mgo.v2/bson"

type Component struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	ProjectId bson.ObjectId `bson:"project_id"`
	Title string `bson:"title"`
}

