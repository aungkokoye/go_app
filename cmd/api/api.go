package api

import (
	"fmt"

	"github.com/aungkokoye/go_app/services"
	"github.com/spf13/cobra"
)

var ApiSyncCmd = &cobra.Command{
	Use:   "api-run",
	Short: "Running API with Gin",
	Run:   main,
}

func init() {
	ApiSyncCmd.Flags().String("config", "config.local.yml", "default config file to use for the application")
}

func main(cmd *cobra.Command, args []string) {

	fmt.Println("calling Album sync")

	fileName := cmd.Flag("config").Value.String()
	fmt.Println(fileName)

	config, err := services.NewConif(fileName)

	if err != nil {
		fmt.Printf("Error reading config file: %s", err)
	}

	fmt.Printf("loading config: %v \n", config)

	app := services.NewApp(config)

	app.RunGin()
}
