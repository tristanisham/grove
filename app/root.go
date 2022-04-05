package app


type groveConfig struct {
	GroveURL string `json:"groveURL"`
	//server is all config variables for the grove server. Will only be applied for the server if it's running
	Server groveServerConfig `json:"server"`
}

type groveServerConfig struct {
	AllowedProxies []string `json:"allowed_proxies"`

}

func DefaultGroveConfig() groveConfig {
	return groveConfig{
		GroveURL: "https://grove.sbs/",
	}
}