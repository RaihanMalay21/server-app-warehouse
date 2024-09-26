package barang

import (
	"net/http"
	"strconv"
	"log"
	"os"
	"fmt"
	"github.com/RaihanMalay21/web-gudang/models"
	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
)

func BarangKeluar(w http.ResponseWriter, r *http.Request) {
	id_barang := r.FormValue("id_barang")
	barangKeluar := r.FormValue("barang_keluar")

	idBarang, _ := strconv.ParseUint(id_barang, 10, 0)
	barang_keluar, _ := strconv.ParseFloat(barangKeluar, 64)

	tx := config.DB.Begin()

	var barang models.Barang
	if err := tx.Where("id = ?", idBarang).First(&barang).Error; err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// kalkulasi pengurangan barang 
	if (barang_keluar > barang.AmountBarang) {
		msg := map[string]string{"messageKeluarBarang": fmt.Sprintf("Total Barang hanya %g", barang.AmountBarang)}
		helper.Response(w, msg, http.StatusOK)
		return
	} 

	sisaBarang := barang.AmountBarang - barang_keluar
	
	if (sisaBarang == 0) {

		// menghapus data
		if err := tx.Delete(&models.Barang{}, barang.ID).Error; err != nil {
			log.Println("Error cant delete barang", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// menghapus gambar
		pathImage := helper.DestinationFolder("C:\\Users\\raiha\\Documents\\web-gudang\\static\\src\\source\\images", barang.Image)
		if err := os.Remove(pathImage); err != nil {
			tx.Rollback()
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {

		if err := tx.Model(&models.Barang{}).Where("id = ?", barang.ID).Update("amount_barang", sisaBarang).Error; err != nil {
			tx.Rollback()
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	tx.Commit()

	msg := map[string]string{"messageKeluarSuccess": "Barang Berhasil Keluar"}
	helper.Response(w, msg, http.StatusOK)
}