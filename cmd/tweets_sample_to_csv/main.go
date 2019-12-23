package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dimitrismistriotis/tweets_sample_to_csv/internal/tweetssampletocsv"
)

func main() {
	fmt.Println("Entry point")
	targetFilename := flag.String("filename", "",
		"File to store samples, will exit if already exists or cannot write to location")
	flag.Parse()

	fmt.Printf("targetFilename: %s\n", *targetFilename)
	if *targetFilename == "" {
		log.Fatal("Empty filename provided")
	}
	// tweets_sample_to_csv.
	tweetssampletocsv.RetrieveAndStore()
	fmt.Println("After entry point")

	fmt.Println(tweetssampletocsv.GetDefaultFilename())
}
