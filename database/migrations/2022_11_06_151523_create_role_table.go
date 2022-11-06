package migrations

import (
	"data-export/pkg/migrate"
	"gorm.io/gorm"
	"time"
)

func init() {
	type Role struct {
		Id         uint   `gorm:"primarykey"`
		Name       string `gorm:"comment:名称;type:varchar(50);uniqueIndex"`
		CreateTime time.Time
		UpdateTime time.Time
	}

	migrate.Add(migrate.MigrationFile{
		FileName: "2022_11_06_151523_create_role_table",
		Up: func(db *gorm.DB) error {
			return db.Migrator().AutoMigrate(&Role{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&Role{})
		},
	}, "default")
}
