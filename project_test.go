package zeleniy

import (
	"reflect"
	"testing"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var TEST_DB_NAME = "zeleniy_test_db"

func TestProjectService(t *testing.T) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Fatal(err)
	}
	db := session.DB(TEST_DB_NAME)
	ps := NewProjectService(db)
	project1 := Project{ID: bson.NewObjectId(), Title: "project1"}
	ps.coll().Insert(project1)
	t.Run("GetProjectById", func(t *testing.T) {
		t.Run("Should return nil project if no such project", func(t *testing.T) {
			p, err := ps.GetProjectById(bson.NewObjectId())
			emptyProj := Project{}
			if p != emptyProj || err != mgo.ErrNotFound {
				t.Errorf("Project is %v error is %+q", p, err)
			}
		})
		t.Run("Should return project by given id", func(t *testing.T) {
			p, err := ps.GetProjectById(project1.ID)
			if p.ID != project1.ID || p.Title != project1.Title || err != nil {
				t.Errorf("Project is %v error is %+q", p, err)
			}
		})
	})
	t.Run("DeleteProject", func(t *testing.T) {
		project2 := Project{ID: bson.NewObjectId(), Title: "project2"}
		ps.coll().Insert(project2)
		t.Run("Should delete project by given id", func(t *testing.T) {
			err := ps.DeleteProject(project2.ID)
			if err != nil {
				t.Errorf("There are should be no error, but %+q", err)
			}
			emptyProj := Project{}
			err = ps.coll().FindId(project2.ID).One(&emptyProj)
			if err != mgo.ErrNotFound {
				t.Errorf("Project should not be found, but %v", emptyProj)
			}
		})
		t.Run("Should not fail if no such project", func(t *testing.T) {
			err := ps.DeleteProject(bson.NewObjectId())
			if err != mgo.ErrNotFound {
				t.Errorf("There are should be ErrNotFound error, but %+q", err)
			}
		})
	})
	t.Run("CreateOrUpdateProject", func(t *testing.T) {
		t.Run("Should create project if no ID provided", func(t *testing.T) {
			newProj := Project{Title: "foo"}
			id, err := ps.CreateOrUpdateProject(newProj)
			if err != nil {
				t.Errorf("Error should be nil, but %+v", err)
			}
			createdProj := Project{}
			err = ps.coll().FindId(id).One(&createdProj)
			if err != nil {
				t.Errorf("Error should be nil, but %+v", err)
			}
			if newProj.Title != createdProj.Title {
				t.Errorf("Expecting title %v, but got %v", newProj.Title, createdProj.Title)
			}
		})
		t.Run("Should update project if project with given ID exists", func(t *testing.T) {
			project1.Title = "bar"
			id, err := ps.CreateOrUpdateProject(project1)
			if err != nil {
				t.Errorf("Error should be nil, but %+v", err)
			}
			createdProj := Project{}
			err = ps.coll().FindId(id).One(&createdProj)
			if err != nil {
				t.Errorf("Error should be nil, but %+v", err)
			}
			if !reflect.DeepEqual(project1, createdProj) {
				t.Errorf("Expecting %v, but got %v", project1, createdProj)
			}
		})
	})
	ps.coll().RemoveAll(nil)
}
