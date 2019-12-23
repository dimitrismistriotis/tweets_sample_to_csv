package tweetssampletocsv

import (
	"fmt"
	"time"
)

// GetDefaultFilename default filename to use if none supplier
func GetDefaultFilename() string {
	return fmt.Sprintf("tweet_samples-%s.csv", time.Now().Format("2006-01-02-15-04-05"))
}
