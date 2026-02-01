package runner

import (
	"context"
	"os"
	"os/exec"
	"time"

	stlerr "github.com/kkkunny/stl/error"

	"github.com/kkkunny/MDM/dal/xunlei"
)

func init() {
	Runners = append(Runners, RunXunlei)
}

func RunXunlei() (<-chan struct{}, <-chan error) {
	cmder := exec.Command("/bin/xlp")
	cmder.Stdout = os.Stdout
	cmder.Stderr = os.Stderr

	succCh := make(chan struct{})
	go func() {
		defer close(succCh)
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for _ = range t.C {
			_, err := xunlei.Client.ListTasks(context.Background())
			if err != nil {
				continue
			}
			return
		}
	}()

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		errCh <- stlerr.ErrorWrap(cmder.Run())
	}()

	return succCh, errCh
}
