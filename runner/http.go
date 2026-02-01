package runner

import (
	"net/http"
	"time"

	stlerr "github.com/kkkunny/stl/error"
	"github.com/labstack/echo/v5"

	"github.com/kkkunny/MDM/handler"
)

func init() {
	Runners = append(Runners, RunHttp)
}

func RunHttp() (<-chan struct{}, <-chan error) {
	svr := echo.New()

	route(svr.Group(""))

	succCh := make(chan struct{})
	go func() {
		defer close(succCh)
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for _ = range t.C {
			ok := func() bool {
				resp, err := http.Get("http://localhost/ping")
				if err != nil {
					return false
				}
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusOK {
					return false
				}
				return true
			}()
			if !ok {
				continue
			}
			return
		}
	}()

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		errCh <- stlerr.ErrorWrap(svr.Start(":80"))
	}()

	return succCh, errCh
}

func route(root *echo.Group) {
	root.GET("/ping", handler.Ping)
}
