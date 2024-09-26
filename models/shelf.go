package models

import (
	"gorm.io/gorm"
	"time"
)

type Shelf struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"ID"`
	NameShelf string `gorm:"type:varchar(10); unique; not null" json:"nama_shelf" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	CapacityRow float64 `gorm:"type:DECIMAL(10, 0);not null" json:"kapasitas_row" validate:"required,number"`
	TotalRows uint32 `gorm:"-" json:"total_rows"`
	CapacityBlocks uint32 `gorm:"-" json:"kapasitas_block"`
	TotalBlocks uint32 `gorm:"-" json:"total_blocks"`
	CapacityBarangs uint32 `gorm:"-" json:"kapasitas_barangs"`
	TotalBarangs uint32 `gorm:"-" json:"total_barangs"`
	Rows []Row `json:"-"`
}