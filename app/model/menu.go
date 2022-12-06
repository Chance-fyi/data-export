package model

type Menu struct {
	Id       uint
	Name     string
	Path     string
	Icon     string
	ParentId uint
	TimeModel
}
