package worker

import (
	"fmt"

	"github.com/aungkokoye/go_app/repositiories"
	"github.com/aungkokoye/go_app/services"
	"github.com/spf13/cobra"
)

var UserSyncCmd = &cobra.Command{
	Use:   "user-sync",
	Short: "Syncs users to the DB provider",
	Run:   main,
}

func init() {
	UserSyncCmd.Flags().String("config", "config.local.yml", "default config file to use for the application")
}

func main(cmd *cobra.Command, args []string) {

	fmt.Println("calling User sync")

	fileName := cmd.Flag("config").Value.String()
	fmt.Println(fileName)

	config, err := services.NewConif(fileName)

	if err != nil {
		fmt.Printf("Error reading config file: %s", err)
	}

	fmt.Printf("loading config: %v \n", config)

	app := services.NewApp(config)

	mysql, err := app.GetDatabase("vivastreet_gb")

	if err != nil {
		fmt.Printf("Error Mysql connection: %s", err)
	}

	repo := repositiories.NewUserRepository(mysql)

	result, err := repo.ContactSync()

	if err != nil {
		fmt.Printf("Error Mysql connection: %s", err)
	}

	for key, value := range result {

		fmt.Printf("[ UserId : %d, UserEmail : %s, Username : %s ]\n", value.UserID, key, value.Username)

	}

	fmt.Println("End")

}
