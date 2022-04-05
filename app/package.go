package app

import (
	"github.com/robertkrimen/otto"
)

type PackageScript struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	License     string `json:"License"`
}

func (p PackageScript) liveUpdate(call otto.FunctionCall) otto.Value {
	return otto.Value{}
}

func LoadPackageScript(js string) error {
	vm := otto.New()
	pk := &PackageScript{}
	vm.Run(js)
	// Get the variable values into the PackageScript object
	for _, i := range []string{"Name", "Description", "Homepage", "License", "sha256"} {
		if err := loadPSvalue(i, pk, vm); err != nil {
			return err
		}
	}

	vm.Set("liveUpdate", pk.liveUpdate)
	return nil
}

//loadPSvalue loads a variable from the VM into a PackageScript
func loadPSvalue(val string, pk *PackageScript, vm *otto.Otto) error {
	if value, err := vm.Get(val); err == nil {
		if vs, err := value.ToString(); err == nil {
			pk.Name = vs
		} else {
			return err
		}
	} else {
		return err
	}
	return nil
}