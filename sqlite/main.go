package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func createDatabase() {
	var err error
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS clients (uuid TEXT PRIMARY KEY, hostname TEXT, IP TEXT)")
	statement.Exec()

	if err != nil {
		log.Printf("Error: %s", err.Error())
		log.Fatal("can't create the table")
	}
}

var database *sql.DB

type Client struct {
	uuid     string
	hostname string
	ip       string
}

func getClient(uuid string) Client {
	fmt.Println("getClient")
	rows, err := database.Query("SELECT uuid, hostname, IP FROM clients WHERE uuid = ?", uuid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var client Client
	row_count := 0
	for rows.Next() {
		err := rows.Scan(&client.uuid, &client.hostname, &client.ip)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("uuid %s, hostname %s, ip %s\n", client.uuid, client.hostname, client.ip)
		row_count++
	}
	fmt.Printf("number of rows returned %d\n", row_count)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if row_count > 1 {
		// Should never get here as uuid is a primary key
		log.Fatal("Multiple hits, this shouldn't happen")
	}

	return client

}

func getClients() {
	rows, _ := database.Query("SELECT uuid, hostname, IP FROM clients")
	var uuid string
	var hostname string
	var ip string
	for rows.Next() {
		rows.Scan(&uuid, &hostname, &ip)
		fmt.Printf("uuid %s, hostname %s, ip %s\n", uuid, hostname, ip)
	}
}

func main() {
	var err error
	database, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Printf("Error: %s", err.Error())
		log.Fatal("can't connect to the database")
	}

	createDatabase()
	getClients()
	client := getClient("461f68d2-d89f-489d-aed7-451851977a18")

	if client == (Client{}) {
		fmt.Printf("No records found\n")
	} else {
		fmt.Printf("uuid %s, hostname %s, ip %s\n", client.uuid, client.hostname, client.ip)
	}
	client = getClient("68d2-d89f-489d-aed7-451851977a18")

	if client == (Client{}) {
		fmt.Printf("No records found\n")
	} else {
		fmt.Printf("uuid %s, hostname %s, ip %s\n", client.uuid, client.hostname, client.ip)
	}
}
