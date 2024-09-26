package models

import (
	"gorm.io/gorm"
	"time"
)

type Block struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"ID"`
	NomorBlock uint `gorm:"not null" json:"nomor_block" validate:"number,max=3,required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	CapacityBarang float64 `gorm:"type:DECIMAL(10, 0);not null" json:"kapasitas" validate:"number,max=10,required"`
	RowID uint `gorm:"not null" json:"row_id" validate:"required"`
	Row Row `gorm:"foreignKey:RowID" json:"row" validate:"-"`
	Barangs []Barang `json:"-"`
}