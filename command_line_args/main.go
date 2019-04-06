package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	reportNamePtr := flag.String("reportName", "report.html", "The name of the report to create")
	wordlistPtr := flag.String("wordlist", "", "The wordlist to process")
	debugPtr := flag.String("debugLevel", "", "Debug options, I = Info, D = Full Debug")
	flag.Parse()

	switch strings.ToUpper(*debugPtr) {
	case "I":
		log.Printf("Debug level: Info\n")
	case "D":
		log.Printf("Debug level: Debug\n")
	default:
		log.Printf("Debug level: Default\n")
	}

	if *wordlistPtr == "" {
		log.Fatal("No wordlist specified")
	}
	log.Printf("Will use the wordlist: %s", *wordlistPtr)

	log.Printf("Will create the output file: %s", *reportNamePtr)

}
