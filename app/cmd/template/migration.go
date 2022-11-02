package template

const TemplateMigration = `package migrations

import (
	"{{PackageName}}/pkg/migrate"
	"gorm.io/gorm"
)

func init() {
	type User struct {
		gorm.Model
		Name string
	}

	migrate.Add(migrate.MigrationFile{
		FileName: "{{FileName}}",
		Up: func(db *gorm.DB) error {
			return db.Migrator().AutoMigrate(&User{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&User{})
		},
	}, "{{Connect}}")
}
`
