package tweetssampletocsv

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

// RetrieveAndStore exposing main functionality of the package
//
func RetrieveAndStore() {
	fmt.Println("Into Retrieve and Store")
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

	for t := range stream.C {
		switch v := t.(type) {
		case anaconda.Tweet:
			fmt.Printf("IdStr: %s\n", v.IdStr)
			// fmt.Printf("User: %v\n", v.User)
			fmt.Printf("UserScreenName: %v\n", v.User.ScreenName)
			// fmt.Printf("ExtendedTweet: %v\n", v.ExtendedTweet)
			fmt.Printf("ExtendedTweetText: %v\n", v.FullText)
			fmt.Printf("ExtendedTweetEntitiesHashtags: %v\n", v.ExtendedEntities)
			fmt.Printf("ExtendedTweetEntitiesUrls: %v\n", v.ExtendedEntities.Urls)
			fmt.Printf("CreatedAt: %s\n", v.CreatedAt)
			fmt.Printf("Lang: %s\n", v.Lang)
			// fmt.Printf("Coordinates: %v\n", v.Coordinates)
			longitude, err := v.Longitude()
			if err != nil {
				fmt.Printf("Longitude: %f\n", longitude)
			} else {
				fmt.Printf("Longitude: -\n")
			}
			latitude, err := v.Latitude()
			if err != nil {
				fmt.Printf("Latitude: %f\n", latitude)
			} else {
				fmt.Printf("Latitude: -\n")
			}
			fmt.Printf("Source: %s\n", v.Source)
			fmt.Printf("Favorited: %v\n", v.Favorited)
			fmt.Printf("FavoriteCount: %v\n", v.FavoriteCount)
			fmt.Printf("Retweeted: %v\n", v.Retweeted)
			fmt.Printf("RetweetCount: %v\n", v.RetweetCount)

			// Other data in Tweet struct:
			// fmt.Printf("DisplayTextRange: %v\n", v.DisplayTextRange)
			// fmt.Printf("FilterLevel: %s\n", v.FilterLevel)
			// fmt.Printf("HasExtendedProfile: %v\n", v.HasExtendedProfile)
			// fmt.Printf("InReplyToScreenName: %s\n", v.InReplyToScreenName)
			// fmt.Printf("InReplyToStatusID: %v\n", v.InReplyToStatusID)
			// fmt.Printf("InReplyToStatusIdStr: %s\n", v.InReplyToStatusIdStr)
			// fmt.Printf("InReplyToUserID: %v\n", v.InReplyToUserID)
			// fmt.Printf("InReplyToUserIdStr: %s\n", v.InReplyToUserIdStr)
			// fmt.Printf("IsTranslationEnabled: %v\n", v.IsTranslationEnabled)
			// fmt.Printf("QuotedStatusID: %v\n", v.QuotedStatusID)
			// fmt.Printf("QuotedStatusIdStr: %s\n", v.QuotedStatusIdStr)
			// fmt.Printf("QuotedStatus: %v\n", v.QuotedStatus)
			// fmt.Printf("PossiblySensitive: %v\n", v.PossiblySensitive)
			// fmt.Printf("PossiblySensitiveAppealable: %v\n", v.PossiblySensitiveAppealable)
			// fmt.Printf("RetweetedStatus: %v\n", v.RetweetedStatus)
			// fmt.Printf("Scopes: %s\n", v.Scopes)
			// fmt.Printf("WithheldCopyright: %v\n", v.WithheldCopyright)
			// fmt.Printf("WithheldInCountries: %s\n", v.WithheldInCountries)
			// fmt.Printf("WithheldScope: %s\n", v.WithheldScope)
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

		break // Only one
	}
}
