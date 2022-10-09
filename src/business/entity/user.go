package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type UserUpdateParam struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type UserSelectParam struct {
	Id uint `uri:"user_id" binding:"required"`
}
