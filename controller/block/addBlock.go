package block

import(
	"encoding/json"
	"log"
	"net/http"

	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/models"
	"github.com/RaihanMalay21/web-gudang/helper"
)

func AddBlock(w http.ResponseWriter, r *http.Request) {
	var field map[string]uint

	JSON := json.NewDecoder(r.Body)
	if err :=  JSON.Decode(&field); err != nil {
		log.Println("Error cant decode json:", err)
		msg := map[string]string{"message": "Error Tidak Dapat Mendecode data json"}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}

	idRow  := field["id_row"]
	nomorblock := field["nomor_block"]

	capacitybarang := float64(field["capacity_barang"])
	capacityblock := float64(field["capacity_block"])

	if float64(nomorblock) > capacityblock {
		msg := map[string]string{"qoutaHabis": "Kouta Block Sudah Habis"}
		helper.Response(w, msg, http.StatusBadRequest)
		return
	}

	Blocks := models.Block{
		NomorBlock: uint(nomorblock),
		CapacityBarang: capacitybarang,
		RowID: uint(idRow),
	}

	if err := config.DB_AddBlock(Blocks); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := map[string]string{"message": "berhasil Membuat Block"}
	helper.Response(w, msg, http.StatusOK)
}