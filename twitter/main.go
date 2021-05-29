package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	//fmt.Println("Starting")

	flags := struct {
		consumerKey    string
		consumerSecret string
	}{}

	flag.StringVar(&flags.consumerKey, "consumer-key", "", "Twitter Consumer Key")
	flag.StringVar(&flags.consumerSecret, "consumer-secret", "", "Twitter Consumer Secret")
	flag.Parse()
	//flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	if flags.consumerKey == "" || flags.consumerSecret == "" {
		log.Fatal("Application Access Token required")
	}

	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     flags.consumerKey,
		ClientSecret: flags.consumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// search tweets
	// example
	// https://github.com/dghubble/go-twitter/blob/master/twitter/search.go
	searchTweetParams := &twitter.SearchTweetParams{
		Query:     "obama",
		TweetMode: "extended",
		Count:     5,
	}

	search, _, _ := client.Search.Tweets(searchTweetParams)
	fmt.Printf("%+v\n", search.Statuses)
	for _, res := range search.Statuses {
		fmt.Printf("XXXXXXXXXXXXXXXX\n")
		fmt.Printf("%s\n", res.FullText)
		// fmt.Printf("%+v\n", res)
	}
	//fmt.Printf("SEARCH TWEETS:\n%+v\n", search)
	//fmt.Printf("SEARCH METADATA:\n%+v\n", search.Metadata)

	// Search Tweets
	/*
		search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
			Query: "gopher",
		})
		fmt.Printf("search: %s\n", search)
		fmt.Println("Done")
	*/
}
