package models

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	WorkspaceID   Workspace `gorm:"foreignKey:WorkspaceRefer"`
	ReleaseID     Release   `gorm:"foreignKey:ReleaseRefer"`
	ModuleID      Module    `gorm:"foreignKey:ModuleRefer"`
	ServiceName   string    `gorm:"type:varchar(100);not null;default:'';comment:'服务名称'"`
	Owner         string    `gorm:"type:varchar(100);not null;default:'';comment:'负责人'"`
	Order         uint      `gorm:"type:not null;default:1;comment:'发布顺序'"validate:"min=1"`
	MirrorVersion string    `gorm:"type:varchar(125);not null;default:'';comment:'镜像版本'"`
	ConfigName    string    `gorm:"type:varchar(125);not null;default:'';comment:'配置名称'"`
	PostStatus    uint      `gorm:"type:not null;default:-1;comment:'发布状态: -1：未定义；1：已锁定；2：发布中；3：已发布；4：发布失败'"`
	OpRelease     bool      `gorm:"type:not null;default:1;comment:'发布操作 0：不可发布 1：可以发布'"`
	OpLock        bool      `gorm:"type:not null;default:0;comment:'锁定操作 0：已锁定 1：可以锁定'"`
	OpDelete      bool      `gorm:"type:not null;default:1;comment:'删除操作 0：可以删除 1：可以删除'"`
}
