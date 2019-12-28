package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dimitrismistriotis/tweets_sample_to_csv/internal/tweetssampletocsv"
)

func main() {
	fmt.Println("Entry point")
	defaultFilename := tweetssampletocsv.GetDefaultFilename()
	targetFilename := flag.String("filename", defaultFilename, "File to store samples")
	maxItemsToDownload := flag.Int64("items_to_download", -1, "Max sample size, -1 (default) for infinite")
	flag.Parse()

	fmt.Printf("targetFilename: %s\n", *targetFilename)
	if *targetFilename == "" {
		log.Fatal("Empty filename provided")
	}
	fmt.Printf("maxItemsToDownload (not used yet): %d\n", *maxItemsToDownload)

	// tweets_sample_to_csv.
	tweetssampletocsv.RetrieveAndStoreToCSV(targetFilename)
	fmt.Println("After entry point")
}
