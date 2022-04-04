package app


type groveConfig struct {
	GroveURL string `json:"groveURL"`
}

type groveServerConfig struct {
	PackageDirs []string `json:"packageDirs"`

}

func DefaultGroveConfig() groveConfig {
	return groveConfig{
		GroveURL: "https://grove.sbs/",
	}
}