package zeleniy

import "gopkg.in/mgo.v2"

type Application struct {
	session *mgo.Session
	dbname string
}


func NewApplication(session *mgo.Session, dbname string) *Application {
	return &Application{dbname: dbname, session: session}
}

func (a *Application) Db() *mgo.Database {
	return a.session.DB(a.dbname)
}
