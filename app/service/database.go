package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/database"
	"data-export/pkg/g"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func CreateDatabase(r api.CreateDatabaseRequest) error {
	database := model.Database{
		Name:     r.Name,
		Hostname: r.Hostname,
		Port:     r.Port,
		Username: r.Username,
		Password: r.Password,
		Database: r.Database,
		Charset:  r.Charset,
	}
	tx := g.DB().Begin()
	err := tx.Create(&database).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = linkDb(database)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func DatabaseList(r api.DatabaseListRequest) (databases []api.DatabaseListItem, count int64) {
	Db := g.DB().Model(&model.Database{})

	if r.Name != "" {
		Db.Where("name like ?", "%"+r.Name+"%")
	}

	Db.Order("id DESC")
	Db.Count(&count)
	Db.Offset((r.Current - 1) * r.PageSize).Limit(r.PageSize).Find(&databases)

	return
}

func GetDatabase(id int) (database api.GetDatabaseResponse) {
	g.DB().Model(&model.Database{}).First(&database, id)
	return
}

func EditDatabase(r api.EditDatabaseRequest) error {
	database := model.Database{
		Id:       r.Id,
		Name:     r.Name,
		Hostname: r.Hostname,
		Port:     r.Port,
		Username: r.Username,
		Password: r.Password,
		Database: r.Database,
		Charset:  r.Charset,
	}

	tx := g.DB().Begin()
	err := tx.Model(&database).Updates(database).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = linkDb(database)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func DatabaseSelectList() (list []api.DatabaseSelectListResponse) {
	g.DB().Model(&model.Database{}).Find(&list)
	return
}

func getDb(id uint) (*gorm.DB, error) {
	name := fmt.Sprintf("connections_%v", id)
	Db := g.DB(name)
	if Db != nil {
		return Db, nil
	}

	var d model.Database
	g.DB().First(&d, id)
	return linkDb(d)
}

func linkDb(d model.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		d.Username,
		d.Password,
		d.Hostname,
		d.Port,
		d.Database,
		d.Charset,
	)
	dialector := mysql.New(mysql.Config{
		DSN: dsn,
	})
	Db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	name := fmt.Sprintf("connections_%v", d.Id)
	database.DB.SetConnections(name, Db)

	return Db, nil
}
