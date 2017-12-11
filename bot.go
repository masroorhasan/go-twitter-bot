package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"github.com/ChimeraCoder/anaconda"
)

type TwitterConfig struct {
	ConsumerKey			string 	`json:"ConsumerKey"`
	ConsumerSecret		string 	`json:"ConsumerSecret"`
	AccessToken			string 	`json:"AccessToken"`
	AccessTokenSecret	string 	`json:"AccessTokenSecret"`
}

var api *anaconda.TwitterApi

func main() {
	twitterConfig := TwitterConfig{}
	loadTwitterConfig(&twitterConfig)
	anaconda.SetConsumerKey(twitterConfig.ConsumerKey)
	anaconda.SetConsumerSecret(twitterConfig.ConsumerSecret)
	api = anaconda.NewTwitterApi(twitterConfig.AccessToken, twitterConfig.AccessTokenSecret)

	// Read from flags
	queryPtr := flag.String("q", "none", "Flag to specify twitter query.")
	actionPtr := flag.String("a", "none", "Flag to specify twitter action.")
	countPtr := flag.String("c", "0", "Flag to specify count of query items.")
	flag.Parse();

	encodedQuery := url.QueryEscape(*queryPtr)
	switch *actionPtr {
	case "favorite":
		favorite(encodedQuery, *countPtr)
	case "retweet":
		retweet(encodedQuery, *countPtr)
	case "follow":
		follow(encodedQuery, *countPtr)
	}
}

// Favorite all tweets for matching query.
func favorite(query string, count string) {
	searchResult, err := api.GetSearch(query, url.Values {"count": []string {count} })
	if err != nil {
		log.Println("Error querying search API.", err)
	}

	for _, tweet := range searchResult.Statuses {
		rt, rtErr := api.Favorite(tweet.Id)
		if rtErr != nil {
			log.Println("Error while favorting.", rtErr)
		} else {
			log.Println("Favorited: twitter.com/" + rt.User.ScreenName + "/status/" + rt.IdStr)
		}
	}
}

// Re-tweet all tweets for matching query.
func retweet(query string, count string) {
	searchResult, err := api.GetSearch(query, url.Values {"count": []string {count} })
	if err != nil {
		log.Println("Error querying search API.", err)
	}

	for _, tweet := range searchResult.Statuses {
		rt, rtErr := api.Retweet(tweet.Id, false)
		if rtErr != nil {
			log.Println("Error while retweeting.", rtErr)
		} else {
			log.Println("Retweeted: twitter.com/" + rt.User.ScreenName + "/status/" + rt.IdStr)
		}
	}
}

// Follow users for matching query.
func follow (query string, count string) {
	searchResult, err := api.GetSearch(query, url.Values {"count": []string{count} })
	if err != nil {
		log.Println("Error querying search API.", err)
	}

	for _, tweet := range searchResult.Statuses {
		user, userErr := api.FollowUser(tweet.User.ScreenName)
		if userErr != nil {
			log.Println("Error while following.", userErr)
		} else {
			log.Println("Followed: twitter.com/" + user.ScreenName)
		}
	}
}

// Load the twitter config from environment var.
func loadTwitterConfig(config *TwitterConfig) {
	configFilePath := os.Getenv("TWITTER_CONFIG_FILE_PATH")
	file, fileErr := os.Open(configFilePath)
	if fileErr != nil {
		log.Fatal("Error opening file.", fileErr)
	}

	jsonDecoder := json.NewDecoder(file)
	decodeErr := jsonDecoder.Decode(&config)
	if decodeErr != nil {
		log.Fatal("Error loading twitter json configuration.", decodeErr)
	}
}