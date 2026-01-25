package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/models"
)

var kategori = []models.Kategori{
	{ID: 1, Nama: "Makanan Berat", Deskripsi: "Makanan Berat Top Seller"},
	{ID: 2, Nama: "Makanan Ringan", Deskripsi: "Makanan Ringan Top Seller"},
	{ID: 3, Nama: "Minuman Dingin", Deskripsi: "Minuman Top Seller Karna Murah hehe"},
}

func GetAllKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

func GetKategoriByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	for _, k := range kategori {
		if k.ID == id {
			json.NewEncoder(w).Encode(k)
			return
		}
	}
	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

func CreateKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var kategoriBaru models.Kategori

	err := json.NewDecoder(r.Body).Decode(&kategoriBaru)
	if err != nil {
		http.Error(w, "Invalid kategori ID", http.StatusBadRequest)
		return
	}

	kategoriBaru.ID = len(kategori) + 1

	kategori = append(kategori, kategoriBaru)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kategoriBaru)
}

func UpdateKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	var updateKategori models.Kategori
	err = json.NewDecoder(r.Body).Decode(&updateKategori)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	for i := range kategori {
		if kategori[i].ID == id {
			updateKategori.ID = id
			kategori[i] = updateKategori

			json.NewEncoder(w).Encode(updateKategori)
			return
		}
	}
	http.Error(w, "kategori belum ada", http.StatusNotFound)
}

func DeleteKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	for i, k := range kategori {
		if k.ID == id {
			kategori = append(kategori[:i], kategori[i+1:]...)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "Sukses Delete",
			})
			return
		}
	}
	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}
