package model

type UserSql struct {
	Id     uint
	UserId uint
	SqlId  uint
	Name   string
	TimeModel
}

func (UserSql) TableName() string {
	return prefix + "user_sql"
}
