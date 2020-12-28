package models

import (
	"github.com/jinzhu/gorm"
)

type Workspace struct {
	gorm.Model
	WorkspaceName string `gorm:"type:varchar(100);not null;default:'';comment:'项目名称'"`
}
