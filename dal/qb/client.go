package qb

import (
	"os"

	"github.com/autobrr/go-qbittorrent"
)

var Client *qbittorrent.Client

func init() {
	qbAddr := os.Getenv("QB_ADDR")
	qbUser := os.Getenv("QB_USER")
	qbPass := os.Getenv("QB_PASS")
	if qbAddr == "" || qbUser == "" || qbPass == "" {
		panic("qbittorrent client not configured")
	}
	Client = qbittorrent.NewClient(qbittorrent.Config{
		Host:     qbAddr,
		Username: qbUser,
		Password: qbPass,
	})
}
