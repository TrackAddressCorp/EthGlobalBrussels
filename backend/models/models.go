package models

import "gorm.io/gorm"

type Petition struct {
	gorm.Model
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	PdfURL          *string `json:"pdf_url"`
	Signs           int     `json:"signs"`
	IndividualVotes []Sign  `json:"individual_votes" gorm:"foreignKey:PetitionID"`
}

type Sign struct {
	gorm.Model
	PetitionID    uint   `json:"petition_id"`
	NullifierHash string `json:"signature"`
}
