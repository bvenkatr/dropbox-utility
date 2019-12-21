package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tj/go-dropbox"
	"github.com/tj/go-dropy"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := dropy.New(dropbox.New(dropbox.NewConfig(os.Getenv("DROPBOX_ACCESS_TOKEN"))))
	filesList, err := client.ListFiles("/test1/")
	if err != nil {
		fmt.Println("Got error while listing the files ", err)
	}

	fmt.Println(len(filesList))

	// fileNamesList := []string{}
	for v, fileInfo := range filesList {
		// fileNamesList = append(fileNamesList, fileInfo.Name())
		fmt.Println(v)
		fmt.Println(fileInfo.Name())
	}
	// fmt.Println("List of file name are", fileNamesList)
}
