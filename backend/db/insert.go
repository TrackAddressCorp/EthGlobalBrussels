package db

import "github.com/TrackAddressCorp/EthGlobalBrussels/models"

func AddPetition(petition models.Petition) error {
	return DB.Create(&petition).Error
}
