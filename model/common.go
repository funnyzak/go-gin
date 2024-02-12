package model

import (
	"gorm.io/gorm"
)

const CtxKeyAuthorizedUser = "ckau"

type Common struct {
	ID        uint64         `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt int64          `gorm:"autoCreateTime:milli;column:created_at" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli;column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
}
