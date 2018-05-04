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

	err = insert(Point{"ghi789", "picnics"})
	check(err)

	err = searchUuid("ghi789")
	check(err)
}

func searchUuid(uuid string) error {
	res := Point{}
	err = c.Find(bson.M{"uuid": uuid}).One(&res)
	fmt.Println(res)
	return err
}

func insert(p Point) error {
	err = c.Insert(&p)
	return err
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
