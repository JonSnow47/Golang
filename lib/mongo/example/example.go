/*
 * Revision History:
 *     Initial: 2018/05/30        Chen Yanchen
 */
package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	".."
)

// global variable
var S *mgo.Session

// model global variable
var Con *mongo.Connection

func init() {
	// init with starServer
	S = mongo.Dial("127.0.0.1:27017", "lighten")

	// init in a model
	Con = mongo.NewConnection(S, "lighten", "object")

	Con.C.EnsureIndex(mgo.Index{
		Key:[]string{"_id"},
		Unique:true,
		DropDups:true,
		Background:true,
	})
}

func main() {
	NewObject()
}

//////////////////// model
type Object struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Field string        `bson:"Field"`
}

func NewObject() {
	con := Con.Connect()
	defer con.Disconnect()

	err := con.C.Insert(&Object{Field: "field"})
	if err != nil {
		log.Println(err)
	}
}
