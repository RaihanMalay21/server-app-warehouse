package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_AddBlock(block models.Block) error {

	if err := DB.Create(&block).Error; err != nil {
		return err
	}

	return nil
}