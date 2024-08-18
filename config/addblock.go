package config

import (
	"github.com/RaihanMalay21/web-gudang/models"
)

func DB_AddBlock(block models.Block, idRow uint) error {
	Row := models.Row{
		ID: idRow,
	}

	DB.Model(&Row).Association("Blocks").Append(&block)

	return nil
}