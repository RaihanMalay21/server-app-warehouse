package config 

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_Blocks(id_row uint) ([]models.Block, error) {
	var Blocks []models.Block

	if err := DB.Where("row_id = ?", id_row).Find(&Blocks).Error; err != nil {
		return nil, err
	}

	return Blocks, nil
}