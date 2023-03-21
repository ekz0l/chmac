package ifconfig

import (
	"fmt"
	"os/exec"
)

var Binary = "ifconfig"

func ExistIF(name string) (ok bool, _ error) {
	cmd := exec.Command(Binary, name)

	if cmdErr := cmd.Run(); cmdErr != nil {
		return false, fmt.Errorf("cannot execute command %q: %s", cmd.String(), cmdErr)
	}

	return true, nil
}
