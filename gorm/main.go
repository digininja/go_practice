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

	create := false

	if create {

		order := models.Order{Name: "and another"}
		models.DB.Create(&order)

		item1 := models.Item{OrderID: order.ID, Name: "curry", Price: 2233}
		item2 := models.Item{OrderID: order.ID, Name: "pizza", Price: 90}
		item3 := models.Item{OrderID: order.ID, Name: "frogs", Price: 180}

		models.DB.Create(&item1)
		models.DB.Create(&item2)
		models.DB.Create(&item3)

		controllers.CreateURL("https://digi.ninja")
	}

	log.Printf("Load an order with ID 2 and get its items")
	var order models.Order
	models.DB.First(&order, 2)
	retrievedItems := controllers.GetItems(order.ID)
	log.Printf("Order name: %s", order.Name)

	for _, item := range retrievedItems {
		log.Printf("\tItem: %s - Price: %d\n", item.Name, item.Price)
	}

	log.Printf("Done")

	// From here https://gorm.io/docs/query.html
	log.Printf("Loading all orders ordered by creation date")
	orders := controllers.GetAllOrders()

	for _, order := range orders {
		retrievedItems := controllers.GetItems(order.ID)
		log.Printf("Order name: %s", order.Name)
		for _, item := range retrievedItems {
			log.Printf("\tItem: %s - Price: %d\n", item.Name, item.Price)
		}
	}
	log.Printf("Done")

	numberOfOrders := 2
	log.Printf("Loading %d orders reverse ordered by creation date", numberOfOrders)

	orders = controllers.GetRecentOrders(numberOfOrders)

	for _, order := range orders {
		retrievedItems := controllers.GetItems(order.ID)
		log.Printf("Order name: %s", order.Name)
		for _, item := range retrievedItems {
			log.Printf("\tItem: %s - Price: %d\n", item.Name, item.Price)
		}
	}
	log.Printf("Done")
	return

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
