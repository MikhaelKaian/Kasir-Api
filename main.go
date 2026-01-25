package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kasir-api/handlers"
)

func main() {
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetProdukByID(w, r)
		} else if r.Method == "PUT" {
			handlers.UpdateProduk(w, r)
		} else if r.Method == "DELETE" {
			handlers.DeleteProduk(w, r)
		}
	})

	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			handlers.GetAllProduk(w, r)
		} else if r.Method == "POST" {
			handlers.CreateProduk(w, r)
		}
	})

	http.HandleFunc("/api/kategori/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlers.GetKategoriByID(w, r)
		} else if r.Method == "PUT" {
			handlers.UpdateKategori(w, r)
		} else if r.Method == "DELETE" {
			handlers.DeleteKategori(w, r)
		}
	})

	http.HandleFunc("/api/kategori", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			handlers.GetAllKategori(w, r)
		} else if r.Method == "POST" {
			handlers.CreateKategori(w, r)
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server Sedang Running di locallhost::8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
