package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Release struct {
	gorm.Model
	ReleaseName string `gorm:"type:varchar(100);not null;default:'';comment:'发布名称'"`
	ReleaseDate time.Time
}
