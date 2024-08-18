package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_RemoveShelf(dataShelf models.Shelf) error {
	if err := DB.Unscoped().Delete(&dataShelf).Error; err != nil {
		return err
	}
	return nil
} 

// cek apakah barang masih tersida di dalam shelf rows block
func DB_CheckItemsExistOrNot(ID_Shelf uint) (bool, error) {
	var shelf models.Shelf
	var count int64

	if err := DB.Preload("Rows.Blocks.Barangs").First(&shelf, "id = ?", ID_Shelf).Error; err != nil {
		return false, err
	}

	// Hitung jumlah barangs dalam shelf
	for _, row := range shelf.Rows {
		for _, block := range row.Blocks {
			count += int64(len(block.Barangs))
		}
	}

	// Jika ada satu atau lebih record yang ditemukan
	if count > 0 {
		return false, nil
	}

	// Jika tidak ada record yang ditemukan
	return true, nil
}
