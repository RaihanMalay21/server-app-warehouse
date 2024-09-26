package barang

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/RaihanMalay21/web-gudang/config"
	"github.com/RaihanMalay21/web-gudang/helper"
	"github.com/RaihanMalay21/web-gudang/models"
)

func AddBarang(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile{
			log.Println(err.Error())
			msg := map[string]string{"message": "Tidak ada file yang di unggah"}
			helper.Response(w, msg, http.StatusBadRequest)
			return
		}
		log.Println(err)
		msg := map[string]string{"Message": "Image Gagal Di Input"}
		helper.Response(w, msg, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// mengambil ext dari nama file
	ext := filepath.Ext(handler.Filename)
	if ext == "" || (ext != ".jpg" && ext != ".png" && ext != ".gif") {
		log.Println("Tipe Gambar harus jpg, png, dan gift")
		msg := map[string]string{"messageImg": "Tipe Gambar harus jpg, png, dan gift"}
		helper.Response(w, msg, http.StatusBadRequest)
		return
	}

	// size image 
	fileSize := handler.Size

	// authentikasi ukuran file 
	if fileSize > 2000000 {
		log.Println("error on line 61 function input barang : Ukuran FIle terlalu besar")
		message := map[string]string{"messageImg":"Ukuran Image Terlalu Besar, max 2MB"}
		helper.Response(w, message, http.StatusBadRequest)
		return
	}

	// mengambil nama filenya 
	nameOnly := filepath.Base(handler.Filename[:len(handler.Filename) - len(ext)])
	
	// menkonversi nama file menggunakan sha256 menjadi byte dan ubah menjadi string
	hasher := sha256.Sum256([]byte(nameOnly))
	namaFileStringByte := hex.EncodeToString(hasher[:])

	idBlockStr := r.FormValue("id_block")
	if idBlockStr == "" {
		fmt.Println("id_block is missing or empty")
		return
	}

	idBlock, err := strconv.ParseUint(idBlockStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return
	}

	amountBarangStr := r.FormValue("amount_barang")
	if amountBarangStr == "" {
		fmt.Println("amount_barang is missing or empty")
		return
	}

	amountBarang, err := strconv.ParseUint(amountBarangStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return
	}

	capacityBarangStr := r.FormValue("capacity_barang")
	if capacityBarangStr == "" {
		fmt.Println("capacity_barang is missing or empty")
		return
	}

	capacity_barang, err:= strconv.ParseUint(capacityBarangStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return
	}

	currentCapacityStr := r.FormValue("current_capacity_barang")
	if currentCapacityStr == "" {
		fmt.Println("current_capacity_barang is missing or empty")
		return
	}

	currentCapacity, err := strconv.ParseUint(currentCapacityStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing string to uint:", err)
		return
	}

	if amountBarang > (capacity_barang - currentCapacity) {
		amount := capacity_barang - currentCapacity
		msg := map[string]string{"messageBarang": fmt.Sprintf("Kapasitas Tidak Cukup Penyimpanan Tersisa  %d Barang", amount)}
		helper.Response(w, msg, http.StatusBadRequest)
		return
	}

	barang := models.Barang{
		Kode: r.FormValue("kode"),
		NameBarang: r.FormValue("nama_barang"),
		Material: r.FormValue("material"),
		Diameter: r.FormValue("diameter"),
		Fitur: r.FormValue("fitur"),
		AmountBarang: float64(amountBarang),
		Image: namaFileStringByte + ext,
		BlockID: uint(idBlock),
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	Trans := helper.TranslatorIDN()

	if err := validate.Struct(&barang); err != nil {
		errors := make(map[string]string)

		// menyimpan errors kedalam map error berupa field dan pesannya
		for _, err := range err.(validator.ValidationErrors) {
			NameField := err.StructField()
			errTranlate := err.Translate(Trans)
			errors[NameField] = errTranlate
		}

		helper.Response(w, errors, http.StatusInternalServerError)
		return
	}

	// cek apakah nama image sudah tersedia di database
	var count int64
	if err := config.DB.Model(&models.Barang{}).Where("image = ?", barang.Image).Count(&count).Error; err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if count > 0 {
		barang.Image = namaFileStringByte + strconv.FormatInt(count, 10) + ext
	} 

	tx := config.DB.Begin()

	if err := config.DB_AddBarang(tx, barang); err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	fileDir := helper.DestinationFolder("C:\\Users\\raiha\\Documents\\web-gudang\\static\\src\\source\\images", barang.Image)

	outFile, err := os.Create(fileDir)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, file); err != nil {
		tx.Rollback()
		// menghapus file yang baru saja di buat
		if err := os.Remove(fileDir); err != nil {
			log.Println(err.Error())
		}
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()

	msg := map[string]string{"message": "Berhasil Menambahkan Barang"}
	helper.Response(w, msg, http.StatusOK)
}