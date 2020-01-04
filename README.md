# Tweets: Sample to CSV

Streams from Twitter's sample API to CSV file. Wanted to create a command line application for ease of use and learning purposes.

All fields are populated from Twitter apart from "LinkToTweet" which is deduced. Names try to match Twitter's API:

* IdStr
* UserScreenName
* ExtendedTweetText
* Hashtags
* Urls
* CreatedAt
* Lang
* Longitude
* Latitude
* Source
* Favorited
* FavoriteCount
* Retweeted
* RetweetCount
* LinkToTweet

## Building

```
go build -v "github.com/dimitrismistriotis/tweets_sample_to_csv/cmd/tweets_sample_to_csv/"
```

### Dependency Management

Used [go modules](https://github.com/golang/go/wiki/Modules).

## Running

Try ```tweets_sample_to_csv -help```. In detail:

* **access_key** (string): Twitter's API access key, if not provided will try to read from environment or .env file
* **access_secret** (string): Twitter's API access secret, if not provided will try to read from environment or .env file
* **consumer_key** (string): Twitter's API consumer key, if not provided will try to read from environment or .env file
* **consumer_secret** (string): Twitter's API consumer secret, if not provided will try to read from environment or .env file
* **filename** (string): File to store samples (default "tweet_samples-2019-12-31-13-53-46.csv")
* **items_to_download** (int): Max sample size, -1 for infinite (default -1)

All are optional. If API access parameters are not provided from running flags, they should be in an environment variable. See ```sample.env``` for details.

## Limitations

As currently is, does not store locations of accompanying media of a Tweet.

##  Code Layout

Followed instructions on [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

## Follow Ups

Did this project to re-learn Go and create a command line utility that was needed. These are some follow ups which might do in the future.

* Use official Twitter's API (Used one used because liked its logging)
* Start writing some tests even for trivial things.
* Check with another coder and rearrange code layout.
* Maybe links to Tweet's media.
* Coordinates observed are zeroed; possible bug.
* Any other bugs or issues from running in the wild.
* Maybe additional fields if any one requests them.
* Check go dep on a fresh install.

## References

Different links used while writing this software:

* <https://github.com/golang-standards/project-layout>
* <https://appdividend.com/2019/11/30/golang-flag-example-how-to-use-command-line-flags/>
* <https://gobyexample.com/command-line-flags>
* <https://developer.twitter.com/en/docs/tweets/sample-realtime/api-reference/get-statuses-sample>
* <https://github.com/ChimeraCoder/anaconda>
* <https://developer.twitter.com/en/docs/developer-utilities/twitter-libraries> (to give it a go later)

## License

See "LICENSE"
