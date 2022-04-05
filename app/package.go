package app

import (
	"github.com/robertkrimen/otto"
)

type PackageScript struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	License     string `json:"License"`
	Sha256      string `json:"sha256"`
	URL string `json:"url"`
}



func LoadPackageScript(js string) error {
	vm := otto.New()
	pk := &PackageScript{}
	vm.Run(js)
	// Get the variable values into the PackageScript object
	if err := loadPSvalue("Name", vm, func(out string) {
		pk.Name = out
	}); err != nil {
		return err
	}
	if err := loadPSvalue("Description", vm, func(out string) {
		pk.Description = out
	}); err != nil {
		return err
	}
	if err := loadPSvalue("Homepage", vm, func(out string) {
		pk.Homepage = out
	}); err != nil {
		return err
	}
	if err := loadPSvalue("License", vm, func(out string) {
		pk.License = out
	}); err != nil {
		return err
	}
	if err := loadPSvalue("sha256", vm, func(out string) {
		pk.Sha256 = out
	}); err != nil {
		return err
	}
	if err := loadPSvalue("URL", vm, func(out string) {
		pk.URL = out
	}); err != nil {
		return err
	}

	return nil
}

//loadPSvalue loads a variable from the VM into a PackageScript
func loadPSvalue(val string, vm *otto.Otto, set func(out string)) error {
	if value, err := vm.Get(val); err == nil {
		if vs, err := value.ToString(); err == nil {
			set(vs)
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}
