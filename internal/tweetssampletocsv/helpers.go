package tweetssampletocsv

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// GetDefaultFilename default filename to use if none supplier
func GetDefaultFilename() string {
	return fmt.Sprintf("tweet_samples-%s.csv", time.Now().Format("2006-01-02-15-04-05"))
}

// printRetrievedTweet helper to retriever, outputs items to stdout
func printRetrievedTweet(v *anaconda.Tweet) {
	fmt.Printf("IdStr: %s\n", v.IdStr)
	// fmt.Printf("User: %v\n", v.User)
	fmt.Printf("UserScreenName: %v\n", v.User.ScreenName)
	// fmt.Printf("ExtendedTweet: %v\n", v.ExtendedTweet)
	fmt.Printf("ExtendedTweetText: %v\n", v.FullText)

	hashTags := ""
	for _, h := range v.Entities.Hashtags {
		hashTags += h.Text + " "
	}

	urls := ""
	for _, u := range v.Entities.Urls {
		urls += u.Display_url + " "
	}

	fmt.Printf("Hashtags: %v\n", hashTags)
	fmt.Printf("Urls: %v\n", urls)
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
}

func GetConfigurationFromParameters(consumerKey *string, consumerSecret *string, accessKey *string, accessSecret *string) *ApiConfig {
	apiConfig := ApiConfig{}

	if *consumerKey != "" {
		apiConfig.ConsumerKey = *consumerKey
	} else {
		apiConfig.ConsumerKey = os.Getenv("CONSUMER_KEY")
	}
	// fmt.Println(apiConfig.ConsumerKey)

	if *consumerSecret != "" {
		apiConfig.ConsumerSecret = *consumerSecret
	} else {
		apiConfig.ConsumerSecret = os.Getenv("CONSUMER_SECRET")
	}
	// fmt.Println(apiConfig.ConsumerSecret)

	if *accessKey != "" {
		apiConfig.AccessKey = *accessKey
	} else {
		apiConfig.AccessKey = os.Getenv("ACCESS_KEY")
	}
	// fmt.Println(apiConfig.AccessKey)

	if *accessSecret != "" {
		apiConfig.AccessSecret = *accessSecret
	} else {
		apiConfig.AccessSecret = os.Getenv("ACCESS_SECRET")
	}
	// fmt.Println(apiConfig.AccessSecret)

	if apiConfig.ConsumerKey == "" || apiConfig.ConsumerSecret == "" ||
		apiConfig.AccessKey == "" || apiConfig.AccessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	return &apiConfig
}
