package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// all arguments are dictated in the master init() function

var plant *cobra.Command = &cobra.Command{
	Use:   "plant",
	Short: "Install a specified package",
	Long: `Plant takes the specified path, finds the package at the end of it, 
			and installs it according to the provided config variables.`,
	Run: plantCmd,
}
// on specifies where Grove should plant a new package. It defaults to "localhost" which just means your machine.
var on string

func plantCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Work in progress....")
}

