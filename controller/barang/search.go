package barang

import (
	"log"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"fmt"
	"github.com/RaihanMalay21/web-gudang/models"
	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
)

func SearchBarang(w http.ResponseWriter, r *http.Request) {
	keyValue := r.FormValue("key_value")
	
	var barang models.Barang

	// if err := config.DB.Preload("Block", func(db *gorm.DB) *gorm.DB {
	// 		return db.Select("nomor_block")
	// 	}).
	// 	Preload("Block.Row", func(db *gorm.DB) *gorm.DB {
	// 		return db.Select("number_row")
	// 	}).
	// 	Preload("Block.Row.Shelf", func(db *gorm.DB) *gorm.DB {
	// 		return db.Select("name_shelf")
	// 	}).
	// 	Where("name_barang = ?", keyValue).
	// 	First(&barang).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		if err := config.DB.Preload("Block", func(db *gorm.DB) *gorm.DB {
	// 				return db.Select("nomor_block")
	// 			}).
	// 			Preload("Block.Row", func(db *gorm.DB) *gorm.DB {
	// 				return db.Select("number_row")
	// 			}).
	// 			Preload("Block.Row.Shelf", func(db *gorm.DB) *gorm.DB {
	// 				return db.Select("name_shelf")
	// 			}).
	// 			Where("kode = ?", keyValue).
	// 			First(&barang).Error; err != nil {
	// 			log.Println(err)
	// 			http.Error(w, err.Error(), http.StatusInternalServerError)
	// 			return
	// 		}
	// 	} else {
	// 		log.Println(err)
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }
	fmt.Println(keyValue)

	if err := config.DB.Preload("Block").Preload("Block.Row").Preload("Block.Row.Shelf").Where("kode = ?", keyValue).First(&barang).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := config.DB.Preload("Block").Preload("Block.Row").Preload("Block.Row.Shelf").Where("name_barang = ?", keyValue).First(&barang).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					log.Println("Error Can't found barang:", err)
					msg := map[string]string{"messageNotFound": "Barang Tidak Di temukan"}
					helper.Response(w, msg, http.StatusBadRequest)
					return
				} else {
					log.Println(err);
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}	
		} else {
			log.Println(err);
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	helper.Response(w, barang, http.StatusOK)
}