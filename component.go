package zeleniy

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/pkg/errors"
)

type Component struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	ProjectId bson.ObjectId `bson:"project_id"`
	Title string `bson:"title"`
}


type ComponentService struct {
	Db *mgo.Database
}


func NewComponentService(db *mgo.Database) ComponentService {
	return ComponentService{Db: db}
}


func (s *ComponentService) GetComponentById(id bson.ObjectId) (Component, error) {
	p := Component{}
	err := s.coll().FindId(id).One(&p)
	if err != nil {
		return p, checkErrNotFound(err, "while retriving project with id ")
	}
	return p, nil
}


func (s *ComponentService) DeleteComponent(id bson.ObjectId) error {
	err := s.coll().RemoveId(id)
	if err != nil {
		return checkErrNotFound(err, "")
	}
	return nil
}


func (s *ComponentService) CreateOrUpdateComponent(p *Component) (bson.ObjectId, error) {
	if p.ID == bson.ObjectId("") {
		p.ID = bson.NewObjectId()
	}
	_, err := s.coll().UpsertId(p.ID, p)
	if err != nil {
		return "", errors.Wrapf(err, "error while inserting/updating project %v", p)
	}
	return p.ID, nil
}


func (s *ComponentService) GetComponentsList() ([]Component, error) {
	projects := []Component{}
	err := s.coll().Find(nil).All(&projects)
	if err != nil {
		return projects, checkErrNotFound(err, "")
	}
	return projects, nil
}


func (s *ComponentService) coll() *mgo.Collection {
	return s.Db.C("metrics")
}
