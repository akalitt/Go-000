package model

import (
	"time"
)

type Base struct {
	IsDel     int       `json:"is_del"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(20);not null"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(20);not null"`
	DeletedBy string    `json:"-" gorm:"type:varchar(20);not null"`
	RequestId string    `json:"-" gorm:"type:char(36);not null"`
}
