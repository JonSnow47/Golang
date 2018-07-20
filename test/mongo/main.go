package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	StartServer()
}

func StartServer() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	Model()
	fmt.Println("sucess")
}

type User struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"Name"`
}

type Order struct {
	Id      bson.ObjectId `bson:"_id"`
	User    User
	Created time.Time `bson:"Created"`
}

func Model() {
	session, err := mgo.Dial("localhost:27017/test")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("test")
	c.EnsureIndex(mgo.Index{
		Key:        []string{"-id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})

	user := &User{
		Id:   bson.NewObjectId(),
		Name: "Joe",
	}

	order := &Order{
		Id:      bson.NewObjectId(),
		User:    *user,
		Created: time.Now(),
	}

	err = c.Insert(order)
	log.Println(err)
}
