// +build windows

package spy

import (
	"errors"
	"os/exec"
)

func setProcessGroupID(cmd *exec.Cmd) {
	//not implemented
}

func killByProcessGroupID(cmd *exec.Cmd) error {
	return errors.New("not implemented")
}
