package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_AddRows(row models.Row) error {
	
	if err := DB.Create(&row).Error; err != nil {
		return err
	}

	return nil
}