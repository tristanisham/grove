package app

import "encoding/json"

type PackageScript struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Homepage    string     `json:"homepage"`
	License     string     `json:"license"`
	Sha256      sha256     `json:"sha256"`
	URL         packageURL `json:"url"`
	Version     string     `json:"version"`
}

type sha256 struct {
	X86_Linux string `json:"x86_linux"`
}

type packageURL struct {
	X86_Linux string `json:"x86_linux"`
}

func LoadPackageScript(js []byte) (*PackageScript, error) {
	pk := new(PackageScript)
	if err := json.Unmarshal(js, pk); err != nil {
		return nil, err
	}
	return pk, nil
}

func CreateDefaultPackageScript() PackageScript {
	return PackageScript{
		Name:        "",
		Description: "",
		Homepage:    "",
		License:     "",
		Sha256:      sha256{X86_Linux: ""},
		URL: packageURL{
			X86_Linux: "",
		},
		Version: "",
	}
}
