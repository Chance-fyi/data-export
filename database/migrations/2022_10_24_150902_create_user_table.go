package migrations

import (
	"data-export/pkg/migrate"
	"gorm.io/gorm"
	"time"
)

func init() {
	type User struct {
		Id         uint   `gorm:"primarykey"`
		Username   string `gorm:"comment:用户名;type:varchar(50);uniqueIndex"`
		Password   string `gorm:"comment:密码;type:char(60)"`
		CreateTime time.Time
	}

	migrate.Add(migrate.MigrationFile{
		FileName: "2022_10_24_150902_create_user_table",
		Up: func(db *gorm.DB) error {
			return db.Migrator().AutoMigrate(&User{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&User{})
		},
	}, "default")
}
