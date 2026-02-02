package xunlei

import (
	"fmt"
	"os"

	stlval "github.com/kkkunny/stl/value"
	"github.com/kkkunny/xunlei"

	"github.com/kkkunny/MDM/config"
)

var Client *xunlei.Client

func init() {
	port := stlval.ValueOr(os.Getenv("XL_DASHBOARD_PORT"), "2345")
	did := os.Getenv("MDM_DID")
	if did == "" {
		panic("unknown xunlei deviceID")
	}
	Client = xunlei.NewClient(stlval.TernaryAction(config.Release, func() string {
		return "http://localhost:" + port
	}, func() string {
		return fmt.Sprintf("http://%s:%s", os.Getenv("XL_DASHBOARD_HOST"), port)
	}), did)
}
