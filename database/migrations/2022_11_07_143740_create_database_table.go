package migrations

import (
	"data-export/pkg/migrate"
	"gorm.io/gorm"
	"time"
)

func init() {
	type Database struct {
		Id         uint   `gorm:"primarykey"`
		Name       string `gorm:"comment:名称;type:varchar(50);uniqueIndex"`
		Hostname   string `gorm:"comment:数据库地址;varchar(50)"`
		Port       string `gorm:"comment:端口;varchar(10)"`
		Username   string `gorm:"comment:用户名;varchar(50)"`
		Password   string `gorm:"comment:密码;varchar(50)"`
		Database   string `gorm:"comment:数据库名;varchar(50)"`
		CreateTime time.Time
		UpdateTime time.Time
	}

	migrate.Add(migrate.MigrationFile{
		FileName: "2022_11_07_143740_create_database_table",
		Up: func(db *gorm.DB) error {
			return db.Migrator().AutoMigrate(&Database{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&Database{})
		},
	}, "default")
}
