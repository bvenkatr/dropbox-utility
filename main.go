package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/joho/godotenv"
)

func listOfFilesToBeDownloaded(dbx files.Client) []files.IsMetadata {
	// start making API calls
	s := files.NewListFolderArg("")
	s.Recursive = true
	res, err := dbx.ListFolder(s)
	if err != nil {
		fmt.Println(err)
	}

	var entries []files.IsMetadata
	entries = res.Entries
	var n int
	for _, entry := range entries {
		switch entry.(type) {
		case *files.FileMetadata:
			// printFileMetadata(f)
			entries[n] = entry
			n++
		}
	}
	return entries[:n]
}

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
	res := listOfFilesToBeDownloaded(dbx)
	for i, entry := range res {
		value := entry.(*files.FileMetadata)
		fmt.Printf("%v.   %s,  %s\n", i, value.PathLower, value.ContentHash)
	}
}
