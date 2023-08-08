package models

import "gorm.io/gorm"

type Logs struct {
	gorm.Model
    Content string `gorm:"type:text" json:"content"`
    UserID  uint
}
