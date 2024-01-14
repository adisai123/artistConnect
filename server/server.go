package server

import (
	"artistConnect/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/artists", allArtist).Methods("GET")
	myRouter.HandleFunc("/artists/{id}", deleteArtist).Methods("DELETE")
	//myRouter.HandleFunc("/artists/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/artist", newArtist).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func allArtist(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []model.Artists
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func deleteArtist(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var artist model.Artists
	db.Where("id = ?", id).Find(&artist)
	db.Delete(&artist)

	fmt.Fprintf(w, "%s Successfully Deleted!!!", artist.Name)
}

func newArtist(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var artist model.Artists
	_ = json.NewDecoder(r.Body).Decode(&artist)
	fmt.Printf("data %#v", artist)
	db.Create(&artist)
	err = json.NewEncoder(w).Encode(&artist)
	if err != nil {
		log.Println(err)
	}
}
