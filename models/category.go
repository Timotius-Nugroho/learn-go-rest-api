package models

import "time"

type Category struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	NamaCategory string `gorm:"type:varchar(300)" json:"nama_category" validate:"required"`
	Deskripsi    string `gorm:"type:text" json:"deskripsi" validate:"required"`
	CreatedAt    time.Time
}
