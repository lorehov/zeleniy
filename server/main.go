package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	fmt.Printf(bson.ObjectIdHex("552bb85210c630552aa849f5").String())
}
