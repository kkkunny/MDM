package runner

import (
	"context"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/acarl005/stripansi"
	stlerr "github.com/kkkunny/stl/error"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/dal/xunlei"
)

func init() {
	Runners = append(Runners, RunXunlei)
}

func RunXunlei() (<-chan struct{}, <-chan error) {
	cmder := exec.Command("/bin/xlp")
	outputer := new(xunleiOutputer)
	cmder.Stdout = outputer
	cmder.Stderr = outputer

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

type xunleiOutputer struct{}

var xunleiLogRegexp = regexp.MustCompile(`\d+/\d+ \d+:\d+:\d+ (\S+) (\[.+?]) `)

func (_ xunleiOutputer) Write(p []byte) (n int, err error) {
	s := strings.TrimSpace(stripansi.Strip(string(p)))
	matches := xunleiLogRegexp.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return stlerr.ErrorWith(os.Stdout.Write(p))
	}

	xlLevel := matches[0][1]
	group := matches[0][2]
	msg := xunleiLogRegexp.ReplaceAllString(s, "")
	switch xlLevel {
	case "DBG":
		err = config.XLLogger.NewGroup(group).Debugf(msg)
	case "INF":
		err = config.XLLogger.NewGroup(group).Infof(msg)
	case "WRN":
		err = config.XLLogger.NewGroup(group).Warnf(msg)
	case "ERR":
		err = config.XLLogger.NewGroup(group).Errorf(msg)
	default:
		return stlerr.ErrorWith(os.Stdout.Write(p))
	}
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
