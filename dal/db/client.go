package db

import (
	"time"

	stlerr "github.com/kkkunny/stl/error"
	"github.com/kkkunny/stl/lazy"
	stlval "github.com/kkkunny/stl/value"
	"gorm.io/gorm/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/kkkunny/MDM/config"
	"github.com/kkkunny/MDM/dal/db/query"
)

var (
	ClientGetter = lazy.Getter(func() (*gorm.DB, error) {
		return stlerr.ErrorWith(gorm.Open(
			sqlite.Open(stlval.Ternary(config.Release, "/config/mdm.db", "mdm.db")),
			&gorm.Config{
				Logger: logger.New(new(customLogger), logger.Config{
					SlowThreshold:             200 * time.Millisecond,
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
				}),
			},
		))
	})
	QueryGetter = lazy.Getter(func() (*query.Query, error) {
		cli, err := ClientGetter()
		if err != nil {
			return nil, err
		}
		query.SetDefault(cli)
		return query.Use(cli), nil
	})
)
