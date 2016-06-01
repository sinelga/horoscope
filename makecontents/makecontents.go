package main

import (
	"flag"
	"fmt"
	"github.com/sinelga/horoscope_libs/dbhandler"
	"github.com/sinelga/horoscope_libs/parse_page"
	"gopkg.in/mgo.v2"
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
	tofeedzodiaclinks := dbhandler.ZodiacContents(*session, "test.com")
	
	for _,link :=range tofeedzodiaclinks {
		
		parse_page.Parse(link)
		
	}
	
	
	

}
