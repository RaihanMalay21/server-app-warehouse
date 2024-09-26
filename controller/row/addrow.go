package row

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
	"github.com/RaihanMalay21/web-gudang/models"
)

func AddRow(w http.ResponseWriter, r *http.Request) {
	var field map[string]uint

	JSON := json.NewDecoder(r.Body)
	if err :=  JSON.Decode(&field); err != nil {
		log.Println("Error cant decode json:", err)
		msg := map[string]string{"message": "Error Tidak Dapat Mendecode data json"}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}

	idShelf := field["id_shelf"]
	numberrow := field["number_row"]

	capacityblock := float64(field["capacity_block"])
	capacityrow := float64(field["capacity_row"])

	fmt.Println(numberrow)
	fmt.Println(capacityrow)
	if float64(numberrow) > capacityrow{
		msg := map[string]string{"qoutaHabis": "Kouta Row Sudah habis"}
		helper.Response(w, msg, http.StatusBadRequest)
		return
	}

	rows := models.Row{
		// data yang dibutuhkan 
		NumberRow: uint(numberrow),
		CapacityBlock: capacityblock,
		ShelfID: uint(idShelf),
	}

	if err := config.DB_AddRows(rows); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := map[string]string{"message": "berhasil Membuat Row"}
	helper.Response(w, msg, http.StatusOK)
}