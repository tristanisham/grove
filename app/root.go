package app


type groveConfig struct {
	GroveURL string `json:"groveURL"`
	//server is all config variables for the grove server. Will only be applied for the server if it's running
	Server groveServerConfig `json:"server"`
}

type groveServerConfig struct {
	PackageDirs []string `json:"packageDirs"`
	BlackListIPs []string `json:"blackListIPs"`
	//AllowedConnnections are all IP addresses the server should expect a connection from. For example, the root server IP if you're running Grove from behind NGINX or Caddy
	AllowedConnections []string `json:"allowedConnections"`

}

func DefaultGroveConfig() groveConfig {
	return groveConfig{
		GroveURL: "https://grove.sbs/",
	}
}