package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/users"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := dropbox.Config{
		Token:    os.Getenv("DROPBOX_ACCESS_TOKEN"),
		LogLevel: dropbox.LogInfo, // if needed, set the desired logging level. Default is off
	}
	dbx := users.New(config)
	// start making API calls
	fmt.Println(dbx)
}
