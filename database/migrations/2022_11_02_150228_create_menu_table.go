package migrations

import (
	"data-export/pkg/migrate"
	"gorm.io/gorm"
	"time"
)

func init() {
	type Menu struct {
		Id         uint   `gorm:"primarykey"`
		Name       string `gorm:"comment:名称;type:varchar(50)"`
		Path       string `gorm:"comment:Path;type:varchar(50)"`
		ParentId   uint   `gorm:"default:0"`
		CreateTime time.Time
		UpdateTime time.Time
	}

	migrate.Add(migrate.MigrationFile{
		FileName: "2022_11_02_150228_create_menu_table",
		Up: func(db *gorm.DB) error {
			return db.Migrator().AutoMigrate(&Menu{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&Menu{})
		},
	}, "default")
}
