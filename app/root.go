package app

import (
	"log"

	"github.com/mitchellh/go-homedir"
)


type groveConfig struct {
	GroveURL string `json:"groveURL"`
	//server is all config variables for the grove server. Will only be applied for the server if it's running
	Server groveServerConfig `json:"server"`
}

type groveServerConfig struct {
	AllowedProxies []string `json:"allowed_proxies"`
	PackagesDirectory string `json:"packages_dir"`

}

func DefaultGroveConfig() groveConfig {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	return groveConfig{
		GroveURL: "https://grove.sbs/",
		Server:   groveServerConfig{
			AllowedProxies:    []string{"0.0.0.0"},
			PackagesDirectory: home+"/.grove/packages/",
		},
	}
}