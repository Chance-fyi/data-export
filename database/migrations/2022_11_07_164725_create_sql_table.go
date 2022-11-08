package migrations

import (
	"data-export/pkg/migrate"
	"gorm.io/gorm"
	"time"
)

func init() {
	type Sql struct {
		Id         uint   `gorm:"primarykey"`
		Sql        string `gorm:"comment:SQL;type:text"`
		Fields     string `gorm:"comment:字段;type:varchar(500)"`
		DatabaseId uint   `gorm:"comment:数据库id;type:int(11)"`
		CreateTime time.Time
		UpdateTime time.Time
	}

	migrate.Add(migrate.MigrationFile{
		FileName: "2022_11_07_164725_create_sql_table",
		Up: func(db *gorm.DB) error {
			return db.Migrator().AutoMigrate(&Sql{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&Sql{})
		},
	}, "default")
}
