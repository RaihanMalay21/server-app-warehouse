package row

import (
	"log"
	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
	"net/http"
	"strconv"
)

func GetRows(w http.ResponseWriter, r *http.Request) {
	id_shelf := r.FormValue("id_shelf")

	// Mengonversi string ke uint64
    id, _:= strconv.ParseUint(id_shelf, 10, 0)

	data, err := config.DB_Rows(uint(id))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.Response(w, data, http.StatusOK)
}