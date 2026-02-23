package qb

import (
	"os"

	"github.com/autobrr/go-qbittorrent"
)

var Client *qbittorrent.Client

func init() {
	addr := os.Getenv("QB_ADDR")
	user := os.Getenv("QB_USER")
	pass := os.Getenv("QB_PASS")
	if addr == "" || user == "" || pass == "" {
		panic("qbittorrent client not configured")
	}
	Client = qbittorrent.NewClient(qbittorrent.Config{
		Host:     addr,
		Username: user,
		Password: pass,
	})
}
