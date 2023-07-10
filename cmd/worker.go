/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/aungkokoye/go_app/cmd/worker"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "A brief description of worker command",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("worker called")
	// },
}

func init() {
	rootCmd.AddCommand(workerCmd)
	workerCmd.AddCommand(worker.UserSyncCmd)
}
