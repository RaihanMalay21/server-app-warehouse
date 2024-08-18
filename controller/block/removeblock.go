package block

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
)

func RemoveBlock(w http.ResponseWriter, r *http.Request) {
	idBlock := r.FormValue("id_block")

	IDBlock, err := strconv.ParseUint(idBlock, 10, 32)
	if err != nil {
		log.Println("Error parsing shelf ID:", err)
		msg := map[string]interface{}{"message": err.Error()}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}	

	exists, err := config.DB_CheckItemsExistOrNotBlock(uint(IDBlock))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		if err := config.DB_RemoveBlock(uint(IDBlock)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		msg := map[string]string{"message": "Berhasil Menghapus Block"}
		helper.Response(w, msg, http.StatusOK)
	} else {
		msg := map[string]string{"message": "Tidak Dapat Menghapus Block, Barang Masih Tersedia"}
		helper.Response(w, msg, http.StatusBadRequest)
	}
} 