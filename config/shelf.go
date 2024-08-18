package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func GetShelfs() (*[]models.Shelf, error) {
	var shelves []models.Shelf
	if err := DB.Preload("Rows").Preload("Rows.Blocks").Preload("Rows.Blocks.Barangs").Find(&shelves).Error; err != nil {
		return nil, err
	}

	// count number of row block and barang
	for i := range shelves {
		shelf := shelves[i]

		// menghitung count row
		shelves[i].JumlahRows = uint32(len(shelf.Rows))

		// Menghitung jumlah blocks
		blockCount := 0
		for _, row := range shelf.Rows {
			blockCount += len(row.Blocks)
		}
		shelf.JumlahBlocks = uint32(blockCount)

		// Menghitung jumlah barangs
		barangCount := 0
		for _, row := range shelf.Rows {
			for _, block := range row.Blocks {
				barangCount += len(block.Barangs)
			}
		}
		shelf.JumlahBarangs = uint32(barangCount)
	}

	return &shelves, nil
}