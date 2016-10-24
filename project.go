package zeleniy

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/pkg/errors"
)

type Project struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Title string `bson:"title"`
}


type ProjectService struct {
	Db *mgo.Database
}


func NewProjectService(db *mgo.Database) ProjectService {
	return ProjectService{Db: db}
}


func (s *ProjectService) GetProjectById(id bson.ObjectId) (Project, error) {
	p := Project{}
	err := s.coll().FindId(id).One(&p)
	if err != nil {
		if err == mgo.ErrNotFound {
			return p, err
		}
		return p, errors.Wrapf(err, "while retriving project with id %v", id)
	}
	return p, nil
}


func (s *ProjectService) DeleteProject(id bson.ObjectId) error {
	err := s.coll().RemoveId(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			return err
		}
		return errors.Wrapf(err, "")
	}
	return nil
}


func (s *ProjectService) CreateOrUpdateProject(p Project) (bson.ObjectId, error) {
	if p.ID == bson.ObjectId("") {
		p.ID = bson.NewObjectId()
	}
	_, err := s.coll().UpsertId(p.ID, p)
	if err != nil {
		return "", errors.Wrapf(err, "error while inserting/updating project %v", p)
	}
	return p.ID, nil
}


func (s *ProjectService) coll() *mgo.Collection {
	return s.Db.C("projects")
}
