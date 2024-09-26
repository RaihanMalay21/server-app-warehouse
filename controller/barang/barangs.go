package barang

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RaihanMalay21/web-gudang/helper"
	"github.com/RaihanMalay21/web-gudang/config"
)

func Barangs(w http.ResponseWriter, r *http.Request) {
	id_block := r.FormValue("id_block")

	// konversi to uint64
	valueID, _ := strconv.ParseUint(id_block, 10, 0)

	block, err := config.DB_Barangs(uint(valueID))
	if err != nil {
		log.Println(err)
		msg := map[string]interface{}{"message": err.Error()}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	} 

	helper.Response(w, block, http.StatusOK)
}