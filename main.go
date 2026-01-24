package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	// "strconv"
	// "strings"
)

// in-memory storage (sementara, nanti ganti database)
type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 100ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "Kecap", Harga: 12000, Stok: 20},
}

func getProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalaid produk ID", http.StatusBadRequest)
		return
	}
	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	//dapet get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	// loop produk, cari id, ganti sesuai data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	// loop produk cari ID
	for i, p := range produk {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Sukses Delete",
			})
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusBadRequest)

}

func main() {
	//GET localhost:8080/api/produk{id}
	//POST localhost:8080/api/produk{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})

	//GET localhost:8080/api/produk
	//POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			// baca data dari request
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			//masukan data ke dalam variable produk
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) //201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// localhost:8080/health
	// ini endpoint pertama kita buat cek server hidup atau tidak
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

// http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request){
// 	if r.Method == "GET" {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(produk)

// 	// baca data dari request
// 	}else if r.Method == "POST"
// 	var produkBaru Produk{
// 	err := json.NewDecoder(r.Body).Decode(&produkBaru)
// 		if err != nil{
// 			http.Error(w, "Invalid request", http.StatusBadRequest)
// 			return
// 		}
// 		// Masukin data ke dalam variable produk
// 		produkBaru.ID = len(produk) + 1
// 		produk = append(produk, produkBaru)

// 		w.Header().Set("Content-type", "application/json")
// 		w.WriteHeader(http.StatusCreated) //201
// 		json.NewEncoder(w).Encode(produkBaru)
// 	}
// })

// func getProdukByID(w http.ResponseWriter, r *http.Request){
// 	//Parse ID dari URL path
// 	//URL: /api/produk/123 -> ID = 123
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Cari produk dengan ID tersebut
// 	for _, p := range produk {
// 		if p.ID == id {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(p)
// 			return
// 		}
// 	}

// 	// Kalau tidak found
// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }

// //put localhost:8080/api/produk/{id}
// func updateProduk(w http.ResponseWriter, r *http.Request){
// 	// get id dari request
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk")

// 	//ganti int
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}

// 	//get data dari request
// 	var updateProduk Produk
// 	err = json.NewDecoder(r.Body).Decode(&updateProduk)
// 	if err != nil {
// 		http.Error(w, "Invalid Request", http.StatusBadRequest)
// 		return
// 	}

// 	// loop produk, cari id, ganti sesuai data dari request
// 	for i := range produk {
// 		if produk[i].ID == id {
// 			updateProduk.ID = id
// 			produk[i] = updateProduk

// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(updateProduk)
// 			return
// 		}
// 	}
// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }

// func deleteProduk(w http.ResponseWriter, r *http.Request) {
// 	//get id
// 	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

// 	//ganti id
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "invalid Produk ID", http.StatusBadRequest)
// 		return
// 	}

// 	//loop produk cari ID, dapet index yang mau dihapus
// 	for i, p := range produk {
// 		if p.ID == id {
// 			//bikin slice baru dengan data sebelum dan sesudah index
// 			produk = append(produk[:i], produk [i+1:]...)

// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(map[string]string{
// 				"message": "sukses delete",
// 			})
// 			return
// 		}
// 	}
// 	http.Error(w, "Produk belum ada", http.StatusNotFound)
// }
