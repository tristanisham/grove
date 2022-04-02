package cmd

import "github.com/spf13/cobra"

var create *cobra.Command = &cobra.Command{
	Use: "create",
	Short: "Create your own grove package",
	Long: "Create generates the neccessary files to help bundle and distribute your application",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

