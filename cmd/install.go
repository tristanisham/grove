package cmd

import (
	"log"
	"strings"

	homedir "github.com/mitchellh/go-homedir"

	git "github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// all arguments are dictated in the master init() function

var install *cobra.Command = &cobra.Command{
	Use:   "install",
	Short: "Install a specified package",
	Long: `install takes the specified path, finds the package at the end of it, 
			and installs it according to the provided config variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Error: command install requires 1 argument. 0 arguments were provided")
		}
		err := gitInstall(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}
// on specifies where Grove should plant a new package. It defaults to "localhost" which just means your machine.
var on string

func gitInstall(path string) error {
	// https://github.com/go-git/go-git/blob/master/_examples/clone/main.go
	sep_path := strings.Split(path,"/")
	name := sep_path[len(sep_path)-1]
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	_, err = git.PlainClone(home + "/.grove/plants/" + name, false, &git.CloneOptions{
		URL: path,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	
	return nil
}