package main

import (
	"fmt"
	"github.com/digininja/go_practice/gorm/controllers"
	"github.com/digininja/go_practice/gorm/models"
	"github.com/digininja/go_practice/gorm/views"
	"log"
)

func main() {
	models.ConnectDatabase()

	controllers.CreateURL("https://digi.ninja")

	urls := controllers.GetURLs()

	for _, url := range urls {
		views.DumpURL(url)
	}

	url, err := controllers.FindURLByID(1)
	if err != nil {
		fmt.Printf("There was an error getting the URL\n")
		log.Printf("There was an error: %s", err)
	} else {
		views.DumpURL(url)
	}

	url, err = controllers.FindURLByID(99)
	if err != nil {
		fmt.Printf("There was an error getting the URL\n")
		log.Printf("There was an error: %s", err)
	} else {
		views.DumpURL(url)
	}

	url, err = controllers.FindURLByURL("https://digi.ninja")
	if err != nil {
		fmt.Printf("There was an error getting the URL\n")
		log.Printf("There was an error: %s", err)
	} else {
		views.DumpURL(url)
	}

	controllers.UpdateURL(url.ID, "error")
	controllers.DeleteURL(url.ID)
	controllers.DeleteURL(url.ID)
	controllers.UpdateURL(url.ID+1, "fish")

	url, err = controllers.FindURLByURL("https://blah.com")
	if err != nil {
		fmt.Printf("There was an error getting the URL\n")
		log.Printf("There was an error: %s", err)
	} else {
		views.DumpURL(url)
	}

}
