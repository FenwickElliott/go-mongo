package main

import (
	"encoding/json"
	"io"
	"net/http"

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

	// http.HandleFunc("/insert", insert)
	http.HandleFunc("/find", find)
	// http.HandleFunc("/remove", remove)
	http.ListenAndServe(":7412", nil)
}

func find(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uID := r.FormValue("uid")
	res := Point{}
	err = c.Find(bson.M{"uid": uID}).One(&res)
	check(err)
	resJSON, err := json.Marshal(res)
	check(err)
	io.WriteString(w, string(resJSON))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
