package worker

import (
	"fmt"

	"github.com/aungkokoye/go_app/repositiories"
	"github.com/aungkokoye/go_app/services"
	"github.com/rs/zerolog/log"
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
	fmt.Printf("Getting config from %s \n", fileName)

	config, err := services.NewConif(fileName)

	if err != nil {
		log.Error().Msgf("Error reading config file: %s", err)
	}

	fmt.Printf("Loading config: %v", config)

	app := services.NewApp(config)

	mysql, err := app.GetDatabase("vivastreet_gb")

	if err != nil {
		log.Error().Msgf("Error Mysql connection: %s\n", err)
	}

	repo := repositiories.NewUserRepository(mysql)

	result, err := repo.ContactSync()

	if err != nil {
		log.Error().Msgf("Error Mysql connection: %s", err)
	}

	for key, value := range result {

		log.Debug().Msgf("[ UserId : %d, UserEmail : %s, Username : %s ]\n", value.UserID, key, value.Username)

	}

	log.Info().Msg("End")

}
