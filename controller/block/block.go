package block

import (
	"log"
	"net/http"
	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
	"strconv"
	"fmt"
)

func GetBlocks(w http.ResponseWriter, r *http.Request) {
	id_row := r.FormValue("id_row")

	// mengkonversi string ke uint64
	id, _ := strconv.ParseUint(id_row, 10, 0)
	fmt.Println(id)
	blocks, err := config.DB_Blocks(uint(id))
	if err != nil {
		log.Println(err)
		msg := map[string]interface{}{"message": err.Error()}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}
	fmt.Println(blocks)
	helper.Response(w, blocks, http.StatusOK)
}