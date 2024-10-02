package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	loadEnv()

	var err error
	db, err = connectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Миграция базы данных
	migrateDB(db)

	// Создание роутера
	router := mux.NewRouter()
	router.HandleFunc("/songs", getSongs).Methods("GET")
	router.HandleFunc("/songs", addSong).Methods("POST")
	router.HandleFunc("/songs/{id:[0-9]+}", updateSong).Methods("PUT")
	router.HandleFunc("/songs/{id:[0-9]+}", deleteSong).Methods("DELETE")

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}

// Получение всех песен
func getSongs(w http.ResponseWriter, r *http.Request) {
	var songs []Song
	db.Find(&songs)
	json.NewEncoder(w).Encode(songs)
}

// Добавление новой песни
func addSong(w http.ResponseWriter, r *http.Request) {
	var newSong Song
	if err := json.NewDecoder(r.Body).Decode(&newSong); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&newSong)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSong)
}

// Обновление существующей песни
func updateSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var song Song
	if err := db.First(&song, id).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	var updatedSong Song
	if err := json.NewDecoder(r.Body).Decode(&updatedSong); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Model(&song).Updates(updatedSong)
	json.NewEncoder(w).Encode(song)
}

// Удаление песни
func deleteSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := db.Delete(&Song{}, id).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
