package db

import "github.com/TrackAddressCorp/EthGlobalBrussels/models"

func GetPetition(id uint) (models.Petition, error) {
	var petition models.Petition

	err := DB.Model(&models.Petition{}).Preload("Pdfs").Find(&petition, id).Error

	return petition, err
}

func PetitionExists(id uint) bool {
	var count int64

	DB.Model(&models.Petition{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func ListPetitions() ([]models.Petition, error) {
	var petitions []models.Petition

	err := DB.Where("finished = ?", true).Find(&petitions).Error
	return petitions, err
}
