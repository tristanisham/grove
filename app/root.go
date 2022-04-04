package app


type groveConfig struct {
	GroveURL string `json:"groveURL"`
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