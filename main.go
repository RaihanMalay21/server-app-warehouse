package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/RaihanMalay21/web-gudang/controller/shelf"
	"github.com/RaihanMalay21/web-gudang/controller/row"
	"github.com/RaihanMalay21/web-gudang/controller/block"
	"github.com/RaihanMalay21/web-gudang/controller/barang"
	"github.com/RaihanMalay21/web-gudang/config"
)

func main() {
	r := mux.NewRouter()
	config.DB_Connection()
	api := r.PathPrefix("/warehouse").Subrouter()
	
	api.HandleFunc("/shelfs", shelf.Shelfs).Methods("GET") // berhasil
	api.HandleFunc("/add/shelf", shelf.AddShelf).Methods("post") // berhasil
	api.HandleFunc("/shelf/rows", row.GetRows).Methods("POST") // berhasil
	api.HandleFunc("/shelf/row/add", row.AddRow).Methods("POST") // berhasil
	api.HandleFunc("/shelf/row/block", block.GetBlocks).Methods("POST")
	api.HandleFunc("/shelf/row/block/add", block.AddBlock).Methods("POST")  //berhasil
	api.HandleFunc("/shelf/row/block/barangs", barang.Barangs).Methods("POST")
	api.HandleFunc("/shelf/row/block/barang/add", barang.AddBarang).Methods("POST") // berhasil
	api.HandleFunc("/shelf/row/block/barang/out", barang.BarangKeluar).Methods("POST")
	api.HandleFunc("/search/barang", barang.SearchBarang).Methods("POST")
	
	// api.HandleFunc("/remove/shelf", shelf.RemoveShelf).Methods("POST") // berhasil
	// api.HandleFunc("/shelf/row/remove", row.RemoveRow).Methods("POST") // berhasil
	// api.HandleFunc("/shelf/row/block/remove", block.RemoveBlock).Methods("POST") // berhasil
	// api.HandleFunc("/shelf/row/block/barang/remove", ).Methods("POST")

	fmt.Println("App running on port http://localhost:8080/warehouse")
	log.Fatal(http.ListenAndServe(":8080", 
		handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowCredentials(),
		)(r)))
}