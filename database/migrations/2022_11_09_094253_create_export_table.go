package migrations

import (
	"data-export/pkg/migrate"
	"gorm.io/gorm"
	"time"
)

func init() {
	type Export struct {
		Id         uint   `gorm:"primarykey"`
		UserId     uint   `gorm:"comment:用户id;type:int(11)"`
		UserSqlId  uint   `gorm:"comment:用户SQLid;type:int(11)"`
		SqlId      uint   `gorm:"comment:SQLid;type:int(11)"`
		DatabaseId uint   `gorm:"comment:数据库id;type:int(11)"`
		Filename   string `gorm:"comment:文件名;varchar(50)"`
		Sql        string `gorm:"comment:SQL;text"`
		Path       string `gorm:"comment:地址;varchar(255)"`
		Status     uint   `gorm:"comment:状态;tinyint(2)"`
		ErrorMsg   string `gorm:"comment:失败原因;varchar(500)"`
		CreateTime time.Time
		UpdateTime time.Time
	}

	migrate.Add(migrate.MigrationFile{
		FileName: "2022_11_09_094253_create_export_table",
		Up: func(db *gorm.DB) error {
			return db.Migrator().AutoMigrate(&Export{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&Export{})
		},
	}, "default")
}
