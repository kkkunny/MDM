package main

import (
	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/service/http"
	"github.com/kkkunny/MDM/service/xunlei"
)

func Run() error {
	errChan := make(chan error)
	go func() {
		errChan <- http.Run()
	}()
	go func() {
		errChan <- xunlei.Run()
	}()
	return <-errChan
}

func main() {
	if err := Run(); err != nil {
		config.Logger.Panic(err)
		panic(err)
	}
}
