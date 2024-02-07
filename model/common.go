package model

import (
	"time"

	"gorm.io/gorm"
)

const CtxKeyAuthorizedUser = "ckau"

type Common struct {
	ID        uint64         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"index;<-:create"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
