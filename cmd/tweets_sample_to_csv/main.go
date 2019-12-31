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
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file, as not all configuration parameters are passed from envirnment")
		log.Fatal(err)
	}

	defaultFilename := tweetssampletocsv.GetDefaultFilename()
	targetFilename := flag.String("filename", defaultFilename, "File to store samples")
	maxItemsToDownload := flag.Int64("items_to_download", -1, "Max sample size, -1 for infinite")

	consumerKey := flag.String("consumer_key", "", "Twitter's API consumer key, if not provided will try to read from environment or .env file")
	consumerSecret := flag.String("consumer_secret", "", "Twitter's API consumer secret, if not provided will try to read from environment or .env file")
	accessKey := flag.String("access_key", "", "Twitter's API access key, if not provided will try to read from environment or .env file")
	accessSecret := flag.String("access_secret", "", "Twitter's API access secret, if not provided will try to read from environment or .env file")
	flag.Parse()

	fmt.Println("Read API keys from environment")

	apiConfig := tweetssampletocsv.GetConfigurationFromParameters(consumerKey, consumerSecret, accessKey, accessSecret)

	fmt.Printf("targetFilename: %s\n", *targetFilename)
	if *targetFilename == "" {
		log.Fatal("Empty filename provided")
	}
	fmt.Printf("maxItemsToDownload (not used yet): %d\n", *maxItemsToDownload)

	// tweets_sample_to_csv.
	tweetssampletocsv.RetrieveAndStoreToCSV(apiConfig, targetFilename, *maxItemsToDownload)
	fmt.Println("After entry point")
}
