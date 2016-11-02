package main

import (
	"github.com/lorehov/zeleniy/httpapi"
	"github.com/lorehov/zeleniy"
	"gopkg.in/mgo.v2"
	"github.com/Sirupsen/logrus"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		logrus.Fatal(err.Error())
	}
	app := zeleniy.NewApplication(session, "zeleniy")
	httpapi.StartApi(app)
}
