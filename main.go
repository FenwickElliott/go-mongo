package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Point struct {
	Uuid      string
	Interests string
}

var err error

func main() {
	session, err := mgo.Dial("127.0.0.1")
	check(err)
	defer session.Close()

	c := session.DB("test").C("points")

	// err = c.Insert(&Point{"abc123", "golfing"})
	// check(err)

	// err = c.Insert(&Point{"def546", "do-hicky-ing"})
	// check(err)

	res := Point{}
	err = c.Find(bson.M{"uuid": "abc123"}).One(&res)
	check(err)

	fmt.Println(res)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
