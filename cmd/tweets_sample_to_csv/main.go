package main

import (
	"fmt"

	"github.com/dimitrismistriotis/tweets_sample_to_csv/internal/tweetssampletocsv"
)

func main() {
	fmt.Println("Entry point")
	// tweets_sample_to_csv.
	tweetssampletocsv.RetrieveAndStore()
	fmt.Println("After entry point")
}
