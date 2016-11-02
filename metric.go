package zeleniy

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/pkg/errors"
)

type Metric struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	UID           string
	Text          string
	Expandable    int
	Leaf          int
	AllowChildren int
}


type MetricService struct {
	Db *mgo.Database
}


func NewMetricService(db *mgo.Database) MetricService {
	return MetricService{Db: db}
}


func (s *MetricService) GetMetricById(id bson.ObjectId) (Metric, error) {
	p := Metric{}
	err := s.coll().FindId(id).One(&p)
	if err != nil {
		return p, checkErrNotFound(err, "while retriving project with id ")
	}
	return p, nil
}


func (s *MetricService) DeleteMetric(id bson.ObjectId) error {
	err := s.coll().RemoveId(id)
	if err != nil {
		return checkErrNotFound(err, "")
	}
	return nil
}


func (s *MetricService) CreateOrUpdateMetric(p *Metric) (bson.ObjectId, error) {
	if p.ID == bson.ObjectId("") {
		p.ID = bson.NewObjectId()
	}
	_, err := s.coll().UpsertId(p.ID, p)
	if err != nil {
		return "", errors.Wrapf(err, "error while inserting/updating project %v", p)
	}
	return p.ID, nil
}


func (s *MetricService) GetMetricsList() ([]Metric, error) {
	projects := []Metric{}
	err := s.coll().Find(nil).All(&projects)
	if err != nil {
		return projects, checkErrNotFound(err, "")
	}
	return projects, nil
}


func (s *MetricService) coll() *mgo.Collection {
	return s.Db.C("metrics")
}
