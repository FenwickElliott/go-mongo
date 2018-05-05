package main

import (
	"encoding/json"
	"io"
	"log"
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

	http.HandleFunc("/insert", insert)
	http.HandleFunc("/find", find)
	http.HandleFunc("/remove", remove)
	log.Fatal(http.ListenAndServe(":7412", nil))
}

func insert(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	err := c.Insert(Point{r.FormValue("uid"), r.FormValue("interests")})
	check(err)
	io.WriteString(w, "Successfully inserted "+r.FormValue("uid")+"\n")
}

func find(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	res := Point{}
	err = c.Find(bson.M{"uid": r.FormValue("uid")}).One(&res)
	check(err)
	resJSON, err := json.Marshal(res)
	check(err)
	io.WriteString(w, string(resJSON)+"\n")
}

func remove(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	err = c.Remove(bson.M{"uid": r.FormValue("uid")})
	check(err)
	io.WriteString(w, "Successfully removed "+r.FormValue("uid")+"\n")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
