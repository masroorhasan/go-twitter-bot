# Go Twitter Bot

Experimental twitter bot written as part of my experience in learning Go programming language.

Twitter API (anaconda) is authenticated from environment variable that can be set.
```sh
export TWITTER_CONFIG_FILE_PATH=/path/to/twitter-config.json
```

## CLI ops

Follow last 10 users involved in given query: `./bot -a=follow -q=#golang -c=10`

Re-tweet last 10 tweets involved in given query: `./bot -a=retweet -q=#golang -c=10`

Favorite last 10 tweets involved in given query: `./bot -a=favorite -q=#golang -c=10`