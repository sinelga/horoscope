package dbhandler

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//	"domains"
	"fmt"
	"log"
)

func CheckIfExist(session mgo.Session, site string, link string) {

	fmt.Println("KSKSKSKSKS")
	
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("horoscope").C("arch")

	count, err := c.Find(bson.M{"Site": site}).Limit(1).Count()

	if err != nil {

		log.Fatal(err)
	}
	if count == 0 {
		fmt.Println("not exists")
	}

}
