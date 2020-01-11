package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	/*
		,
			LogLevel: dropbox.LogInfo, // if needed, set the desired logging level. Default is off
	*/
	config := dropbox.Config{
		Token: os.Getenv("DROPBOX_ACCESS_TOKEN"),
	}
	dbx := files.New(config)
	// s := new(files.ListFolderResult)

	// start making API calls
	s := files.NewListFolderArg("")
	s.Recursive = true
	res, err := dbx.ListFolder(s)
	if err != nil {
		fmt.Println(err)
	}

	var entries []files.IsMetadata
	entries = res.Entries

	for _, entry := range entries {
		switch f := entry.(type) {
		case *files.FileMetadata:
			printFileMetadata(f)
		}
	}
}

func printFileMetadata(e *files.FileMetadata) {
	fmt.Printf("%s\n", e.PathDisplay)
}
