package tweetssampletocsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// RetrieveAndStoreToCSV wraps retrieve and store
func RetrieveAndStoreToCSV(apiConfig *APIConfig, targetFilename *string, maxItemsToDownload int64) {
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
	RetrieveAndStore(apiConfig, csvwriter, maxItemsToDownload)
}

func getCsvHeaders() []string {
	return []string{"IdStr", "UserScreenName", "ExtendedTweetText",
		"Hashtags", "Urls", "CreatedAt",
		"Lang", "Longitude", "Latitude", "Source", "Favorited", "FavoriteCount",
		"Retweeted", "RetweetCount", "LinkToTweet"}
}

// RetrieveAndStore exposing main functionality of the package
//
func RetrieveAndStore(apiConfig *APIConfig, writer *csv.Writer, itemsToDownload int64) {
	fmt.Printf("Into Retrieve and Store, items to download: %d", itemsToDownload)

	log.Println("start")

	// api := anaconda.NewTwitterApiWithCredentials(accessToken, accessSecret, consumerKey, consumerSecret)
	api := anaconda.NewTwitterApiWithCredentials(apiConfig.AccessKey, apiConfig.AccessSecret, apiConfig.ConsumerKey, apiConfig.ConsumerSecret)
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
			if len(v.Entities.Hashtags) > 0 {
				for _, h := range v.Entities.Hashtags {
					hashTags += h.Text
				}
			}
			urls := ""
			if len(v.Entities.Urls) > 0 {
				for _, u := range v.Entities.Urls {
					urls += u.Display_url + " "
				}
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
				fmt.Sprintf("%v", urls),
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
