package tweetssampletocsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

// RetrieveAndStoreToCSV wraps retrieve and store
func RetrieveAndStoreToCSV(targetFilename *string, maxItemsToDownload int64) {
	fmt.Println("Target filename: ", *targetFilename)
	csvfile, err := os.Create(*targetFilename)
	if err != nil {
		log.Print("Could not open CSV file")
		log.Fatal(err)
	}

	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	defer csvwriter.Flush()

	csvwriter.Write(getCsvHeaders())
	RetrieveAndStore(csvwriter, maxItemsToDownload)
}

func getCsvHeaders() []string {
	return []string{"IdStr", "UserScreenName", "ExtendedTweetText",
		"Hashtags", "ExtendedTweetEntitiesUrls", "CreatedAt",
		"Lang", "Longitude", "Latitude", "Source", "Favorited", "FavoriteCount",
		"Retweeted", "RetweetCount", "LinkToTweet"}
}

// RetrieveAndStore exposing main functionality of the package
//
func RetrieveAndStore(writer *csv.Writer, itemsToDownload int64) {
	fmt.Printf("Into Retrieve and Store, items to download: %d", itemsToDownload)
	log.SetOutput(os.Stdout)

	log.Println("start")

	fmt.Println("Preparing for Twitter retrieval, start with API keys from environment")
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
		log.Fatal(err)
	}
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessKey := os.Getenv("ACCESS_KEY")
	accessSecret := os.Getenv("ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessKey == "" || accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	// api := anaconda.NewTwitterApiWithCredentials(accessToken, accessSecret, consumerKey, consumerSecret)
	api := anaconda.NewTwitterApiWithCredentials(accessKey, accessSecret, consumerKey, consumerSecret)
	api.SetLogger(anaconda.BasicLogger)
	fmt.Println(*api.Credentials)

	// searchResult, _ := api.GetSearch("golang", nil)
	// for _, tweet := range searchResult.Statuses {
	// 	fmt.Println(tweet.Text)
	// }

	v := url.Values{}
	v.Set("delimited", "true")
	v.Set("tweet_mode", "extended")
	// v.Set("stall_warnings", "false")
	v.Set("language", "en")
	stream := api.PublicStreamSample(v)
	defer stream.Stop()

	var itemsDownloaded int64
	for t := range stream.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			printRetrievedTweet(&v)

			screenName := v.User.ScreenName
			hashTags := ""
			for _, h := range v.Entities.Hashtags {
				hashTags += h.Text
			}
			longitude, err := v.Longitude()
			longitudeStr := ""
			if err != nil {
				longitudeStr = fmt.Sprintf("%f", longitude)
			}
			latitude, err := v.Latitude()
			latitudeStr := ""
			if err != nil {
				latitudeStr = fmt.Sprintf("%f", latitude)
			}

			linkToTweet := fmt.Sprintf("https://twitter.com/%s/status/%s", screenName, v.IdStr)

			writer.Write([]string{v.IdStr,
				fmt.Sprintf("%v", screenName),
				fmt.Sprintf("%v", v.FullText),
				fmt.Sprintf("%v", hashTags),
				fmt.Sprintf("%v", v.ExtendedEntities.Urls),
				v.CreatedAt,
				v.Lang,
				longitudeStr,
				latitudeStr,
				v.Source,
				fmt.Sprintf("%v", v.Favorited),
				fmt.Sprintf("%v", v.FavoriteCount),
				fmt.Sprintf("%v", v.Retweeted),
				fmt.Sprintf("%v", v.RetweetCount),
				linkToTweet,
			})

			itemsDownloaded++
			log.Printf("Items Downloaded: %d", itemsDownloaded)
			if (itemsToDownload != -1) && (itemsDownloaded == itemsToDownload) {
				return
			}

		case anaconda.EventTweet:
			switch v.Event.Event {
			case "favorite":
				sn := v.Source.ScreenName
				tw := v.TargetObject.Text
				fmt.Printf("Favorited by %-15s: %s\n", sn, tw)
			case "unfavorite":
				sn := v.Source.ScreenName
				tw := v.TargetObject.Text
				fmt.Printf("UnFavorited by %-15s: %s\n", sn, tw)
			}
		}
	}
}
