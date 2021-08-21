package model

import (
	"time"
)

// gorm.Model 的定义
type CommonModel struct {
	ID        uint `torm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
