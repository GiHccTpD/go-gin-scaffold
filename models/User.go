package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username   string   `gorm:"unique;type:varchar(100);not null;comment:'登录名'"  validate:"min=1"`
	UsernameZH string   `gorm:"type:varchar(100);not null;default:'';comment:'姓名'"`
	Password   string   `gorm:"type:varchar(255);not null" validate:"min=11"`
	Secrets    []Secret `gorm:"many2many:user_secret;"`
	Status     bool     `gorm:"default:'0';comment:'用户状态 0正常 1禁用'"`
	Avatar     string   `gorm:"type:varchar(512);not null;default:'http://image.thepaper.cn/wap/image/18/396/362.png';comment:'用户头像'"`
	Email      string   `gorm:"type:varchar(256);not null;default:'';comment:'用户邮箱'"`
	Role       string   `gorm:"type:varchar(128);not null;default:'';comment:'用户角色'"`
}

type Secret struct {
	ID        uint   `gorm:"primary_key"`
	Secretid  string `gorm:"type:varchar(255);unique;not null" validate:"min=1"`
	Secretkey string `gorm:"type:varchar(255);unique;not null" validate:"min=1"`
}
