package dbhandler

import (
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestCheckIfExist(t *testing.T) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	CheckIfExist(*session,"test.com","link")

}
