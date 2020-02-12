package main

import "fmt"
import (
	"context"
	"github.com/likexian/doh-go"
	"github.com/likexian/doh-go/dns"
	"time"
)

func main() {
	fmt.Println("start")

	// init a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// init doh client, auto select the fastest provider base on your like
	// you can also use as: c := doh.Use(), it will select from all providers
	c := doh.Use(doh.CloudflareProvider, doh.GoogleProvider)

	// do doh query
	rsp, err := c.Query(ctx, "digi.ninja", dns.TypeTXT)
	if err != nil {
		panic(err)
	}

	// close the client
	c.Close()

	// doh dns answer
	answer := rsp.Answer

	// print all answer
	for _, a := range answer {
		fmt.Printf("%s -> %s\n", a.Name, a.Data)
	}

}
