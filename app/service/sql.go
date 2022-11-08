package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/g"
	"data-export/pkg/sqlparse"
	"fmt"
	"github.com/thoas/go-funk"
	"strconv"
)

func CreateSql(r api.CreateSqlRequest) error {
	fields, err := sqlparse.GetColAsName(r.Sql)
	if err != nil {
		return err
	}
	sql := model.Sql{
		Sql:    r.Sql,
		Fields: fields,
	}
	err = g.DB().Create(&sql).Error
	return err
}

func SqlList(r api.SqlListRequest) (sqls []api.SqlListItem, count int64) {
	Db := g.DB().Model(&model.Sql{})

	if r.Fields != "" {
		Db.Where("fields like?", "%"+r.Fields+"%")
	}

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
	fields, err := sqlparse.GetColAsName(r.Sql)
	if err != nil {
		return err
	}
	sql := model.Sql{
		Id:     r.Id,
		Sql:    r.Sql,
		Fields: fields,
	}
	err = g.DB().Model(&sql).Updates(sql).Error

	return err
}

func GetUserSql(id int) (userIds []string) {
	var userSql []model.UserSql
	g.DB().Where("sql_id = ?", id).Find(&userSql)
	for _, u := range userSql {
		userIds = append(userIds, strconv.Itoa(int(u.UserId)))
	}
	return
}

func SetUserSql(r api.SetUserSqlRequest) error {
	userIds := GetUserSql(r.Id)
	delIds, insIds := funk.DifferenceString(userIds, r.UserIds)
	tx := g.DB().Begin()
	err := tx.Where("sql_id = ? and user_id in ?", r.Id, delIds).Delete(&model.UserSql{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	var userSql []model.UserSql
	for _, id := range insIds {
		userId, err := strconv.Atoi(id)
		if err != nil {
			tx.Rollback()
			return err
		}
		userSql = append(userSql, model.UserSql{
			UserId: uint(userId),
			SqlId:  uint(r.Id),
		})
	}
	if len(userSql) > 0 {
		err = tx.Create(&userSql).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func MySqlList(r api.MySqlListRequest, user model.User) (list []api.MySqlListItem, count int64) {
	var (
		us model.UserSql
		s  model.Sql
	)
	Db := g.DB().Debug().Table(fmt.Sprintf("%v us", us.TableName())).
		Joins(fmt.Sprintf("left join %v s on s.ID = us.sql_id", s.TableName()))

	if r.Name != "" {
		Db.Where("us.name like ?", "%"+r.Name+"%")
	}

	if r.Fields != "" {
		Db.Where("s.fields like ?", "%"+r.Fields+"%")
	}

	Db.Where("us.user_id = ?", user.Id)
	Db.Order("us.id DESC")
	Db.Count(&count)
	Db.Offset((r.Current-1)*r.PageSize).Limit(r.PageSize).
		Select("us.id", "us.name", "s.fields", "us.sql_id").
		Scan(&list)

	return
}

func GetUserSqlName(id int) (userSql api.GetUserSqlNameResponse) {
	g.DB().Model(&model.UserSql{}).First(&userSql, id)
	return
}

func SetUserSqlName(r api.SetUserSqlNameRequest) error {
	us := model.UserSql{
		Id:   r.Id,
		Name: r.Name,
	}
	err := g.DB().Model(&us).Updates(us).Error
	return err
}
