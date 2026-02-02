package xunlei

import (
	"os"

	stlval "github.com/kkkunny/stl/value"
	"github.com/kkkunny/xunlei"
)

var Client *xunlei.Client

func init() {
	port := stlval.ValueOr(os.Getenv("XL_DASHBOARD_PORT"), "2345")
	did := stlval.ValueOr(os.Getenv("MDM_DID"), "")
	Client = xunlei.NewClient("http://localhost:"+port, did)
}
