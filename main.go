package main

import (
	"fmt"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
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
		d := &files.DownloadArg{Path: value.PathLower}
		f, data, err := dbx.Download(d)
		if err != nil {
			log.Fatal(err)
		}
		outFile, err := os.Create("/tmp/" + f.Name)
		if err != nil {
			log.Fatal(err)
		}
		// handle err
		defer outFile.Close()
		_, err = io.Copy(outFile, data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
