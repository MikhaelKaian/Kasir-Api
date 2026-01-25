package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/models"
)

var produk = []models.Produk{
	{ID: 1, Nama: "Nasi Goreng", Harga: 12000, Stok: 15},
	{ID: 2, Nama: "Kentang Goreng", Harga: 8000, Stok: 30},
	{ID: 3, Nama: "Es Teh Manis", Harga: 3000, Stok: 30},
}

func GetAllProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

func GetProdukByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func CreateProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var produkBaru models.Produk
	err := json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	produkBaru.ID = len(produk) + 1
	produk = append(produk, produkBaru)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(produkBaru)
}

func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid produk ID", http.StatusBadRequest)
		return
	}
	var updateProduk models.Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
	http.Error(w, "produk belum ada", http.StatusNotFound)
}

func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}
	http.Error(w, "produk belum ada", http.StatusNotFound)
}
