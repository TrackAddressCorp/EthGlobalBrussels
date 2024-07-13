package db

import "github.com/TrackAddressCorp/EthGlobalBrussels/models"

func GetPetition(id uint) (models.Petition, error) {
	var petition models.Petition

	err := DB.Model(&models.Petition{}).Preload("Pdfs").Find(&petition, id).Error

	return petition, err
}

func ListPetitions() ([]models.Petition, error) {
	var petitions []models.Petition

	err := DB.Find(&petitions).Error
	return petitions, err
}
