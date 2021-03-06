package cmd

import (
	"grove/server"
	"log"

	"github.com/spf13/cobra"
)

var serve *cobra.Command = &cobra.Command{
	Use: "serve",
	Short: "starts your own package management server",
	Long: `Running the <Serve> command starts a production grade server the Grove binary can interact with.
			It's identical to the default server Grove pings, and is meant for anyone who wants to host their own packages.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := server.Start(AllowedServerProxies); err != nil {
			log.Panic(err)
		}
	},
}

