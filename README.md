# go_app
go application for DB


# Go Module Create

go mod init github.com/aungkokoye/go_app   // github_url/username/package_name
go mod tidy


# Go Cobra cmd

go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest

~/go/bin/cobra-cli init
~/go/bin/cobra-cli add <command_name> // add new command

All good and run `go run main.go`

# GO mysql

go get -u github.com/go-sql-driver/mysql