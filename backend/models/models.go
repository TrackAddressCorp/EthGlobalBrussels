package models

import "gorm.io/gorm"

type Petition struct {
	gorm.Model
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	PdfURL          *string `json:"pdf_url"`
	Votes           int     `json:"votes"`
	IndividualVotes []Vote  `json:"individual_votes" gorm:"foreignKey:PetitionID"`
}

type Vote struct {
	gorm.Model
	PetitionID    uint   `json:"petition_id"`
	NullifierHash string `json:"signature"`
}
