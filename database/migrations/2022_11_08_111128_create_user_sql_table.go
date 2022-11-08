package migrations

import (
	"data-export/pkg/migrate"
	"gorm.io/gorm"
	"time"
)

func init() {
	type UserSql struct {
		Id         uint   `gorm:"primarykey"`
		UserId     uint   `gorm:"comment:用户id;type:int(11);uniqueindex:user_sql"`
		SqlId      uint   `gorm:"comment:SQLid;type:int(11);uniqueindex:user_sql"`
		Name       string `gorm:"comment:名称;varchar(50)"`
		CreateTime time.Time
		UpdateTime time.Time
	}

	migrate.Add(migrate.MigrationFile{
		FileName: "2022_11_08_111128_create_user_sql_table",
		Up: func(db *gorm.DB) error {
			return db.Migrator().AutoMigrate(&UserSql{})
		},
		Down: func(db *gorm.DB) error {
			return db.Migrator().DropTable(&UserSql{})
		},
	}, "default")
}
