package main

import (
	"time"

	stlerr "github.com/kkkunny/stl/error"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/runner"
	"github.com/kkkunny/MDM/util"
)

func main() {
	succChs := make([]<-chan struct{}, len(runner.Runners))
	errChs := make([]<-chan error, len(runner.Runners))
	for i, r := range runner.Runners {
		succChs[i], errChs[i] = r()
	}

	succCh := util.MixRChannel(succChs...)
	errCh := util.MixRChannel(errChs...)
	timeoutCh := time.After(time.Second * 10)

	var checkOk bool
	var err error
loop:
	for {
		if !checkOk {
			select {
			case <-succCh:
				checkOk = true
				config.Logger.Keywordf("mdm start")
				continue
			case err = <-errCh:
				break loop
			case <-timeoutCh:
				err = stlerr.Errorf("runner check timeout")
				break loop
			}
		} else {
			err = <-errCh
			break
		}
	}
	if err != nil {
		config.Logger.Panic(err)
		panic(err)
	}
}
