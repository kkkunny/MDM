package xl

import (
	"os"

	"github.com/kkkunny/xunlei"
)

var Client *xunlei.Client

func init() {
	addr := os.Getenv("XL_ADDR")
	did := os.Getenv("XL_DID")
	if addr == "" || did == "" {
		panic("xunlei client not configured")
	}
	Client = xunlei.NewClient(addr, did)
}
