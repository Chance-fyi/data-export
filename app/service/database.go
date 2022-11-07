package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/g"
)

func CreateDatabase(r api.CreateDatabaseRequest) error {
	database := model.Database{
		Name:     r.Name,
		Hostname: r.Hostname,
		Port:     r.Port,
		Username: r.Username,
		Password: r.Password,
		Database: r.Database,
	}
	err := g.DB().Create(&database).Error
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
	}
	err := g.DB().Model(&database).Updates(database).Error

	return err
}
