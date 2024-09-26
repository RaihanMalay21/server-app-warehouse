package config

import (
	"github.com/RaihanMalay21/web-gudang/models"

	"gorm.io/gorm"
)

func DB_AddBarang(tx *gorm.DB, barang models.Barang) error {
	
    if err := tx.Create(&barang).Error; err != nil {
        return err
    }

	return nil
}