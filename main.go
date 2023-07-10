package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aungkokoye/go_app/cmd"
)

func main() {

	fmt.Println("Running go_app!")

	err := cmd.Execute()
	if err != nil {
		log.Println(err)

		os.Exit(1)
	}

	os.Exit(0)

}
