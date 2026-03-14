package main

import (
	stlerr "github.com/kkkunny/stl/error"
	stlval "github.com/kkkunny/stl/value"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/kkkunny/MDM/config"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "dal/db/query",
		ModelPkgPath:      "po",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	db := stlerr.MustWith(stlerr.ErrorWith(gorm.Open(sqlite.Open(stlval.Ternary(config.Release, "/config/mdm.db", "mdm.db")))))
	g.UseDB(db)
	models := g.GenerateAllTable()
	g.ApplyBasic(models...)
	g.Execute()
}
