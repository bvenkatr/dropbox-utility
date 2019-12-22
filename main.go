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
	res, err := dbx.ListFolder(files.NewListFolderArg(""))
	if err != nil {
		fmt.Println(err)
	}

	var entries []files.IsMetadata
	entries = res.Entries

	for i, entry := range entries {
		switch f := entry.(type) {
		case *files.FileMetadata:
			printFileMetadata(f)
		case *files.FolderMetadata:
			printFolderMetadata(f)
		}
		if i%4 == 0 {
			fmt.Println("i%4 == 0...., nothing ")
		}
	}
}

func printFolderMetadata(e *files.FolderMetadata) {
	fmt.Println("*files.FileMetadata....")
	fmt.Printf("%s\t\n", e.PathDisplay)
}

func printFileMetadata(e *files.FileMetadata) {
	fmt.Println("*files.FolderMetadata....")
	fmt.Printf("%s\t", e.PathDisplay)
}
