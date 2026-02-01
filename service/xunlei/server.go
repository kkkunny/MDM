package xunlei

import (
	"os"
	"os/exec"

	stlerr "github.com/kkkunny/stl/error"
)

func Run() error {
	cmder := exec.Command("/bin/xlp")
	cmder.Stdout = os.Stdout
	cmder.Stderr = os.Stderr
	err := stlerr.ErrorWrap(cmder.Run())
	return err
}
