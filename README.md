# dropbox-utility
 A dropbox utility for syncing files from remote to local and vice versa written in go

### Setup
* Create a file called .env in root folder of the project
* Add ```DROPBOX_ACCESS_TOKEN=TOKEN_VALUE```

## Initializing the project with go modules
go mod init dropbox-utility

## Build
go build
./dropbox-utility

## Another way to run the project
go run main.go