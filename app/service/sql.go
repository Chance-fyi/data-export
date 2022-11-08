package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/g"
)

func CreateSql(r api.CreateSqlRequest) error {
	sql := model.Sql{
		Sql: r.Sql,
	}
	err := g.DB().Create(&sql).Error
	return err
}

func SqlList(r api.SqlListRequest) (sqls []api.SqlListItem, count int64) {
	Db := g.DB().Model(&model.Sql{})

	Db.Order("id DESC")
	Db.Count(&count)
	Db.Offset((r.Current - 1) * r.PageSize).Limit(r.PageSize).Find(&sqls)

	return
}

func GetSql(id int) (sql api.GetSqlResponse) {
	g.DB().Model(&model.Sql{}).First(&sql, id)
	return
}

func EditSql(r api.EditSqlRequest) error {
	sql := model.Sql{
		Id:  r.Id,
		Sql: r.Sql,
	}
	err := g.DB().Model(&sql).Updates(sql).Error

	return err
}
