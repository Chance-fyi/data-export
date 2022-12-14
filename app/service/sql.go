package service

import (
	"data-export/app/api"
	"data-export/app/model"
	"data-export/pkg/g"
	"data-export/pkg/sqlparse"
	"fmt"
	"github.com/thoas/go-funk"
	"regexp"
	"strconv"
	"strings"
)

func CreateSql(r api.CreateSqlRequest) error {
	reg, _ := regexp.Compile("[\\{\\}]")
	s := reg.ReplaceAll([]byte(r.Sql), []byte(""))
	fields, err := sqlparse.GetColAsName(string(s))
	if err != nil {
		return err
	}
	databaseId, err := strconv.Atoi(r.DatabaseId)
	if err != nil {
		return err
	}
	sql := model.Sql{
		Name:       r.Name,
		Sql:        r.Sql,
		Fields:     fields,
		DatabaseId: uint(databaseId),
	}
	err = g.DB().Create(&sql).Error
	return err
}

func SqlList(r api.SqlListRequest) (sqls []api.SqlListItem, count int64) {
	var (
		s model.Sql
		d model.Database
	)

	Db := g.DB().Table(fmt.Sprintf("%v s", s.TableName())).
		Joins(fmt.Sprintf("left join %v d on d.id = s.database_id", d.TableName()))

	if r.Fields != "" {
		Db.Where("s.fields like ?", "%"+r.Fields+"%")
	}

	if r.DatabaseId != "" {
		Db.Where("s.database_id = ?", r.DatabaseId)
	}

	Db.Order("s.id DESC")
	Db.Count(&count)
	Db.Offset((r.Current-1)*r.PageSize).Limit(r.PageSize).
		Select("s.id", "s.name", "s.fields", "d.name database_name").
		Scan(&sqls)

	return
}

func GetSql(id int) (sql api.GetSqlResponse) {
	g.DB().Model(&model.Sql{}).First(&sql, id)
	return
}

func EditSql(r api.EditSqlRequest) error {
	reg, _ := regexp.Compile("[\\{\\}]")
	s := reg.ReplaceAll([]byte(r.Sql), []byte(""))
	fields, err := sqlparse.GetColAsName(string(s))
	if err != nil {
		return err
	}
	databaseId, err := strconv.Atoi(r.DatabaseId)
	if err != nil {
		return err
	}
	sql := model.Sql{
		Id:         r.Id,
		Name:       r.Name,
		Sql:        r.Sql,
		Fields:     fields,
		DatabaseId: uint(databaseId),
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
			Name:   r.Name,
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
		d  model.Database
	)
	Db := g.DB().Table(fmt.Sprintf("%v us", us.TableName())).
		Joins(fmt.Sprintf("left join %v s on s.ID = us.sql_id", s.TableName())).
		Joins(fmt.Sprintf("left join %v d on d.ID = s.database_id", d.TableName()))

	if r.Name != "" {
		Db.Where("us.name like ?", "%"+r.Name+"%")
	}

	if r.Fields != "" {
		Db.Where("s.fields like ?", "%"+r.Fields+"%")
	}

	if r.DatabaseId != "" {
		Db.Where("s.database_id = ?", r.DatabaseId)
	}

	Db.Where("us.user_id = ?", user.Id)
	Db.Order("us.id DESC")
	Db.Count(&count)
	Db.Offset((r.Current-1)*r.PageSize).Limit(r.PageSize).
		Select("us.id", "us.name", "s.fields", "us.sql_id", "d.name database_name").
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

func GetDownloadSql(id int) api.GetDownloadSqlResponse {
	var sql model.Sql
	g.DB().First(&sql, id)

	r, _ := regexp.Compile("\\{[^\\}]+\\}")
	fields := r.FindAllString(sql.Sql, -1)
	fields = funk.UniqString(fields)
	res := api.GetDownloadSqlResponse{
		Id:  sql.Id,
		Sql: sql.Sql,
	}
	for _, field := range fields {
		field = strings.Trim(strings.Trim(field, "{"), "}")
		f := strings.ToLower(field)
		switch {
		case strings.Contains(f, "datetime"):
			res.Fields = append(res.Fields, api.GetDownloadSqlField{
				Name: field,
				Type: "datetime",
			})
		case strings.Contains(f, "date"):
			res.Fields = append(res.Fields, api.GetDownloadSqlField{
				Name: field,
				Type: "date",
			})
		case strings.Contains(f, "time"):
			res.Fields = append(res.Fields, api.GetDownloadSqlField{
				Name: field,
				Type: "time",
			})
		case strings.Contains(f, "text"):
			res.Fields = append(res.Fields, api.GetDownloadSqlField{
				Name: field,
				Type: "text",
			})
		default:
			res.Fields = append(res.Fields, api.GetDownloadSqlField{
				Name: field,
				Type: "text",
			})
		}
	}
	return res
}
