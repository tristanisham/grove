package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"

	git "github.com/go-git/go-git/v5"
	_ "github.com/joho/godotenv/autoload"
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
		err := groveInstall(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}

// on specifies where Grove should plant a new package. It defaults to "localhost" which just means your machine.
var on string

// groveInstall talks to a grove server instance and install the requested file if it exists or returns an app.NoPackageFound error
func groveInstall(path string) error {
	if !strings.Contains(path, "/") {
		return errors.New("error: install path requires <name>/<version> format.")
	}
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	name_ver := strings.Split(path, "/")
	if err := os.MkdirAll(fmt.Sprintf("%s/.grove/plants/%s/%s/", home, name_ver[0], name_ver[1]), 0775); err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s/.grove/plants/%s/%s/%s", home, name_ver[0], name_ver[1], name_ver[0]))
	if err != nil {
		return err
	}
	defer file.Close()
	client := http.DefaultClient
	resp, err := client.Get(os.Getenv("GROVE_REPO") + "/plants/" + path)
	if err != nil {
		return fmt.Errorf("%s | %s | %d", err, os.Getenv("GROVE_REPO") + "/plants/" + path, 62)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Planted %s from %s at %s as %d bytes\n", name_ver[0], os.Getenv("GROVE_REPO"), fmt.Sprintf("%s/.grove/plants/%s/%s/%s", home, name_ver[0], name_ver[1], name_ver[0]), size)
	return nil
}

// gitInstall install a git directory to disk.
func gitInstall(path string) error {
	// https://github.com/go-git/go-git/blob/master/_examples/clone/main.go
	sep_path := strings.Split(path, "/")
	name := sep_path[len(sep_path)-1]
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	_, err = git.PlainClone(home+"/.grove/plants/"+name, false, &git.CloneOptions{
		URL:               path,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	return nil
}
