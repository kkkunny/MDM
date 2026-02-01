package xunlei

import (
	"os"

	stlval "github.com/kkkunny/stl/value"
	"github.com/kkkunny/xunlei"
)

var Client *xunlei.Client

func init() {
	port := stlval.ValueOr(os.Getenv("XL_DASHBOARD_PORT"), "2345")
	Client = xunlei.NewClient("http://localhost:"+port, "")
}
