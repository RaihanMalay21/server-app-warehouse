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
		

		// menghitung capacity row
		shelves[i].TotalRows = uint32(len(shelf.Rows))

		// Menghitung total capacity blocks
		total_block := 0
		capacity_block := 0
		for _, row := range shelf.Rows {
			total_block += len(row.Blocks)
			capacity_block += int(row.CapacityBlock)
		}
		shelves[i].CapacityBlocks = uint32(capacity_block)
		shelves[i].TotalBlocks = uint32(total_block)

		// Menghitung capacity barangs
		capacity_barang := 0
		total_barangs := 0
		for _, row := range shelf.Rows {
			for _, block := range row.Blocks {
				total_barangs += len(block.Barangs)
				capacity_barang += int(block.CapacityBarang)
			}
		}
		shelves[i].CapacityBarangs = uint32(capacity_barang)
		shelves[i].TotalBarangs = uint32(total_barangs)

	}

	return &shelves, nil
}