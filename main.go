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

var (
	err error
	c   *mgo.Collection
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	check(err)
	defer session.Close()

	c = session.DB("test").C("points")

	err = searchUuid("def546")
	check(err)
}

func searchUuid(uuid string) error {
	res := Point{}
	err = c.Find(bson.M{"uuid": uuid}).One(&res)
	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
