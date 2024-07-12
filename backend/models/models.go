package models

import "gorm.io/gorm"

type Petition struct {
	gorm.Model
	Title           string `json:"title"`
	Description     string `json:"description"`
	Votes           int    `json:"votes"`
	IndividualVotes []Vote `json:"individual_votes" gorm:"foreignKey:PetitionID"`
}

type Vote struct {
	gorm.Model
	PetitionID uint   `json:"petition_id"`
	Signature  string `json:"signature"`
}
