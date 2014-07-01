package deisctl

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"github.com/coreos/fleet/unit"
)

// path hierarchy for finding systemd service templates
var rootPaths = []string{"/var/lib/deis/units", "units", "../units"}

// NewUnit takes a component type and returns a Fleet unit
// that includes the relevant systemd service template
func NewUnit(component string) (u *unit.Unit, err error) {
	template, err := readTemplate(component)
	if err != nil {
		return
	}
	u, err = unit.NewUnit(string(template))
	if err != nil {
		return
	}
	return
}

// formatUnitName returns a properly formatted systemd service name
// using the given component type and number
func formatUnitName(component string, num int) (unitName string, err error) {
	unitName = "deis-" + component + "." + strconv.Itoa(num) + ".service"
	return
}

// readTemplate returns the contents of a systemd template for the given component
func readTemplate(component string) (out []byte, err error) {
	templateName := "deis-" + component + ".service"
	var templateFile string
	for _, rootPath := range rootPaths {
		filename := path.Join(rootPath, component, templateName)
		if _, err := os.Stat(filename); err == nil {
			templateFile = filename
			break
		}
	}
	if templateFile == "" {
		return nil, fmt.Errorf("Could not find unit template for %v", component)
	}
	out, err = ioutil.ReadFile(templateFile)
	if err != nil {
		return
	}
	return
}
