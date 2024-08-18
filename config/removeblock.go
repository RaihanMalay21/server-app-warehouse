package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_RemoveBlock(id_block uint) error {
	if err := DB.Where("id = ?", id_block).Delete(&models.Block{}).Error; err != nil {
		return err
	}
	return nil
}

func DB_CheckItemsExistOrNotBlock(id_block uint) (bool, error) {
	var count int64

	if err := DB.Model(&models.Barang{}).
		Joins("JOIN block_barangs ON block_barangs.barang_id = barangs.id").
		Where("block_barangs.block_id = ?", id_block).
		Count(&count).Error; err != nil {
		return false, err
	}
	
	if count > 0 {
		return false, nil
	}

	return true, nil
}