package models

import (
	"gorm.io/gorm"
	"time"
)

type Row struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"ID"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	NumberRow uint `gorm:"not null" json:"nomor_row" validate:"required,number"`
	CapacityBlock  float64 `gorm:"type:DECIMAL(10, 0);not null" json:"kapasitas" validate:"required,number"`
	AmountCapacityBarang float64 `gorm:"-" json:"total_capacity_barang"`
	ShelfID uint  `gorm:"not null" validate:"required" json:"shelf_id"`
	Shelf Shelf `gorm:"foreignKey:ShelfID;referenses:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"shelf" validate:"-"`
	Blocks []Block `json:"-"`
}