package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Point struct {
	UID       string
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

	err = insert(Point{"abc123", "hullabaloo"})
	check(err)

	err = find("abc123")
	check(err)

	err = remove("abc123")
	check(err)
}

func find(uID string) error {
	res := Point{}
	err = c.Find(bson.M{"uid": uID}).One(&res)
	fmt.Println(res)
	return err
}

func insert(p Point) error {
	err = c.Insert(&p)
	return err
}

func remove(uID string) error {
	err = c.Remove(bson.M{"uid": uID})
	return err
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
