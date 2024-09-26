package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_Rows(id_shelf uint) ([]models.Row, error) {
	var Rows []models.Row

	if err := DB.Preload("Blocks").Preload("Blocks.Barangs").Where("shelf_id = ?", id_shelf).Find(&Rows).Error; err != nil {
		return nil, err
	}

	for i, value := range Rows {
		var amountCapacityBarangs float64
		for _, block := range value.Blocks {
			amountCapacityBarangs += block.CapacityBarang
		}
		Rows[i].AmountCapacityBarang = amountCapacityBarangs
	}

	return Rows, nil
}