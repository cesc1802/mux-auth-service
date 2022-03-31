package cmd

import (
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the application",
	Run:   runServeCmd,
}

func init() {
	serveCmd.AddCommand(httpCmd)
}

func runServeCmd(cmd *cobra.Command, args []string) {
	err := cmd.Help()
	if err != nil {
		panic(err)
	}
}
