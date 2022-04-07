package app

import (
	"archive/tar"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

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

func Tar(source, target string) error {
    filename := filepath.Base(source)
    target = filepath.Join(target, fmt.Sprintf("%s.tar", filename))
    tarfile, err := os.Create(target)
    if err != nil {
        return err
    }
    defer tarfile.Close()
 
    tarball := tar.NewWriter(tarfile)
    defer tarball.Close()
 
    info, err := os.Stat(source)
    if err != nil {
        return nil
    }
 
    var baseDir string
    if info.IsDir() {
        baseDir = filepath.Base(source)
    }
 
    return filepath.Walk(source,
        func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            header, err := tar.FileInfoHeader(info, info.Name())
            if err != nil {
                return err
            }
 
            if baseDir != "" {
                header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
            }
 
            if err := tarball.WriteHeader(header); err != nil {
                return err
            }
 
            if info.IsDir() {
                return nil
            }
 
            file, err := os.Open(path)
            if err != nil {
                return err
            }
            defer file.Close()
            _, err = io.Copy(tarball, file)
            return err
        })
}