package config

import (
	"github.com/RaihanMalay21/web-gudang/models"

	"gorm.io/gorm"
)

func DB_AddBarang(tx *gorm.DB, id_Block uint, barang models.Barang) error {
	// Simpan barang terlebih 
    if err := tx.Create(&barang).Error; err != nil {
        return err
    }

	block := models.Block{
		ID: id_Block,
	}

	// tambah kan barang ke tabel relations
	if err := tx.Model(&block).Association("Barangs").Append(&barang); err != nil {
		return err
	}

	return nil
}