package main

import (
	"os"
	"path/filepath"
	"strings"

	stlerr "github.com/kkkunny/stl/error"
	"gorm.io/gen"

	"github.com/kkkunny/MDM/dal/db"
)

func removeGenFile(dir string) error {
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, fileInfo := range fileInfos {
		if !strings.HasSuffix(fileInfo.Name(), ".gen.go") {
			continue
		}
		err = os.Remove(filepath.Join(dir, fileInfo.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := removeGenFile("dal/db/query")
	if err != nil {
		panic(err)
	}
	err = removeGenFile("dal/db/po")
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:           "dal/db/query",
		ModelPkgPath:      "po",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	g.UseDB(stlerr.MustWith(db.ClientGetter()))
	models := g.GenerateAllTable()
	g.ApplyBasic(models...)
	g.Execute()
}
