package models

import "gorm.io/gorm"

type Petition struct {
	gorm.Model
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	Finished        bool    `json:"finished"`
	Signs           int     `json:"signs"`
	IndividualSigns []Sign  `json:"individual_votes" gorm:"foreignKey:PetitionID"`
	Pdfs            []Pdf   `json:"pdfs" gorm:"foreignKey:PetitionID"`
}

type Sign struct {
	gorm.Model
	PetitionID    uint   `json:"petition_id"`
	NullifierHash string `json:"signature"`
}

type Pdf struct {
	gorm.Model
	PetitionID uint   `json:"petition_id"`
	PdfURL     string `json:"pdf_url"`
	Hash       string `json:"hash"`
}
