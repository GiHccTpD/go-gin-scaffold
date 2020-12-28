package models

import (
	"github.com/jinzhu/gorm"
)

type Module struct {
	gorm.Model
	ModuleName string `gorm:"type:varchar(100);not null;default:'';comment:'模块名称'"`
}
