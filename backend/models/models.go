package models

import "gorm.io/gorm"

type Petition struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Signatures  int    `json:"signatures"`
}
