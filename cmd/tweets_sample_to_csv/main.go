package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dimitrismistriotis/tweets_sample_to_csv/internal/tweetssampletocsv"
	"github.com/joho/godotenv"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Println("Entry point")
	defaultFilename := tweetssampletocsv.GetDefaultFilename()
	targetFilename := flag.String("filename", defaultFilename, "File to store samples")
	maxItemsToDownload := flag.Int64("items_to_download", -1, "Max sample size, -1 (default) for infinite")
	flag.Parse()

	fmt.Println("Read API keys from environment")
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
		log.Fatal(err)
	}

	apiConfig := tweetssampletocsv.ApiConfig{}
	apiConfig.ConsumerKey = os.Getenv("CONSUMER_KEY")
	apiConfig.ConsumerSecret = os.Getenv("CONSUMER_SECRET")
	apiConfig.AccessKey = os.Getenv("ACCESS_KEY")
	apiConfig.AccessSecret = os.Getenv("ACCESS_SECRET")

	if apiConfig.ConsumerKey == "" || apiConfig.ConsumerSecret == "" ||
		apiConfig.AccessKey == "" || apiConfig.AccessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	fmt.Printf("targetFilename: %s\n", *targetFilename)
	if *targetFilename == "" {
		log.Fatal("Empty filename provided")
	}
	fmt.Printf("maxItemsToDownload (not used yet): %d\n", *maxItemsToDownload)

	// tweets_sample_to_csv.
	tweetssampletocsv.RetrieveAndStoreToCSV(&apiConfig, targetFilename, *maxItemsToDownload)
	fmt.Println("After entry point")
}
