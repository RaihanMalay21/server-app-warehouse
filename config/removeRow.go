package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_RemoveRow(ID_Row uint) error {
	if err := DB.Unscoped().Where("id = ?", ID_Row).Delete(&models.Row{}).Error; err != nil {
		return err
	}
	return nil
}

// cek apakah barang masih tersida di dalam shelf rows block
func DB_CheckItemsExistOrNotRow(ID_Row uint) (bool, error) {
	var count uint64
	var row models.Row

	if err := DB.Preload("Blocks.Barangs").First(&row, "id = ?", ID_Row).Error; err != nil {
		return false, err
	}

	for _, block := range row.Blocks{
		count += uint64(len(block.Barangs))
	}

	// Jika ada satu atau lebih record yang ditemukan
	if count > 0 {
		return false, nil
	}

	// Jika tidak ada record yang ditemukan
	return true, nil
}