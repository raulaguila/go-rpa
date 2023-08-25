package main

import (
	"os"

	"github.com/go-rod/rod/lib/utils"
	"github.com/joho/godotenv"
	"github.com/raulaguila/go-rpa/internal/rpa/twitter"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func main() {
	myRpa := twitter.RPATwitter{}
	myRpa.Init(true)
	defer myRpa.CloseAll()

	if err := myRpa.Login(os.Getenv("TWITTER_USER"), os.Getenv("TWITTER_PASS")); err != nil {
		panic(err)
	}

	myRpa.ListTweets()
	utils.Pause()
}
