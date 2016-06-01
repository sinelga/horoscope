package main

import (
	"flag"
	"fmt"
	"github.com/sinelga/horoscope_libs/dbhandler"
	"gopkg.in/mgo.v2"
	"os"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)

	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	contentsarr := dbhandler.GetAllContents(*session, "test.com")

	file, err := os.OpenFile("/tmp/mcontents.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, contents := range contentsarr {

		if _, err = file.WriteString(contents); err != nil {
			panic(err)
		}

	}

}
