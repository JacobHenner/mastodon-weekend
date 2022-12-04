package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/mattn/go-mastodon"
)

type config struct {
	instanceUrl      string
	appClientId      string
	appClientSecret  string
	userId           string
	userPassword     string
	mediaPath        string
	mediaDescription string
}

func parseArgs() config {
	// TODO: Ensure CLI arguments are set appropriately
	var instanceUrl = flag.String("instance-url", "https://socialcoders.org", "Mastodon instance URL")
	var appClientId = flag.String("app-client-id", "", "Application client ID")
	var userId = flag.String("user-id", "", "Mastodon user id (email)")
	var mediaPath = flag.String("media-path", "", "Local path of media to be shared")
	var mediaDescription = flag.String("media-description", "Actor Daniel Craig introduces The Weeknd", "Description of media file")
	flag.Parse()

	var userPassword = os.Getenv("USER_PASSWORD")
	var appClientSecret = os.Getenv("APP_CLIENT_SECRET")

	return config{instanceUrl: *instanceUrl, appClientId: *appClientId, appClientSecret: appClientSecret, userId: *userId, userPassword: userPassword, mediaPath: *mediaPath, mediaDescription: *mediaDescription}
}

func main() {
	config := parseArgs()
	toot(config)
}

func toot(config config) {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       config.instanceUrl,
		ClientID:     config.appClientId,
		ClientSecret: config.appClientSecret,
	})
	err := c.Authenticate(context.Background(), config.userId, config.userPassword)
	if err != nil {
		log.Fatal(err)
	}

	videoFile, err := os.Open(config.mediaPath)
	if err != nil {
		log.Fatal(err)
	}

	video := mastodon.Media{File: videoFile, Description: config.mediaDescription}
	attachment, err := c.UploadMediaFromMedia(context.Background(), &video)
	if err != nil {
		log.Fatal(err)
	}

	toot := mastodon.Toot{MediaIDs: []mastodon.ID{attachment.ID}}
	_, err = c.PostStatus(context.Background(), &toot)
	if err != nil {
		log.Fatal(err)
	}
}
