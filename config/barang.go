package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_Barangs(id_block uint) ([]models.Barang, error) {
	var barang []models.Barang

	if err := DB.Where("block_id = ?", id_block).Find(&barang).Error; err != nil {
		return barang, err
	}

	return barang, nil
}