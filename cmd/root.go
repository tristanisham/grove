package cmd

import (
	"encoding/json"
	"grove/app"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AllowedServerProxies defines which real IPs Grove's webserver expects if you host it behind a proxy.
var AllowedServerProxies []string
var rootCmd = &cobra.Command{
	Use:   "grove",
	Short: "Grove is a modern package manager and software installer",
	Long:  `Grove is a modern package manager and installer built to make setting up your favorite programs as easy as possible`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(install, browse, create, serve)
	install.Flags().StringVar(&on, "on", "localhost", "target for grove to plant a seed")
}

// initConfig creates Grove's app directory in ~/.grove.
func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(home + "/.grove"); os.IsNotExist(err) {
		os.MkdirAll(home+"/.grove", 0775)
		os.MkdirAll(home+"/.grove/packages", 0775)
	}
	//Bug(tristan): I need to redo this whole thing wihtout viper
	viper.AddConfigPath(home + "/.grove")
	viper.SetConfigName("config")
	viper.SetDefault("groveURL", "https://grove.sbs/")
	viper.SetDefault("allowed_proxies", []string{"0.0.0.0"})
	viper.SetDefault("packages_dir", home+"/.grove/packages/")
	if val := viper.GetString("groveURL"); len(val) == 0 {
		os.Setenv("GROVE_REPO", val)
	}
	os.Setenv("GROVE_PKG_DIR", home+"/.grove/packages/")
	AllowedServerProxies = viper.GetStringSlice("allowed_proxies")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired

			cf, err := json.MarshalIndent(app.DefaultGroveConfig(), "", "    ")
			if err != nil {
				log.Panic(err)
			}
			if err := os.WriteFile(home+"/.grove/config.json", cf, 0775); err != nil {
				log.Panic(err)
			}
			initConfig()
		} else {
			// Config file was found but another error was produced
			log.Fatal(err)
		}
	}
}
