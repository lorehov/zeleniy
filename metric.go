package zeleniy

import "gopkg.in/mgo.v2/bson"

type Metric struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	UID           string
	Text          string
	Expandable    int
	Leaf          int
	AllowChildren int
}
