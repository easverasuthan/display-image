package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
    	"io/ioutil"
	"log"
	"os"
	"strconv"
	"net/http"
)

var db *sql.DB
var err error

func init() {
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }

}
func main() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	db, err = sql.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	if err != nil {
		panic(err.Error())
	}
	router := mux.NewRouter()
	router.HandleFunc("/images/{id}", getPost).Methods("GET")
	http.ListenAndServe(":80", router)
	defer db.Close()
}
func getPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := db.Query("SELECT image_url FROM images WHERE image_id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var imagestr string
	for result.Next() {
		err := result.Scan(&imagestr)
		if err != nil {
			panic(err.Error())
		}
	}
	log.Println(imagestr)
	
    response, e := http.Get(imagestr)
    if e != nil {
        log.Fatal(e)
    }
    defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("response length of " + params["id"] + " : " + strconv.Itoa(len(bodyBytes)))

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(bodyBytes)))
	if _, err := w.Write(bodyBytes); err != nil {
		log.Println("unable to write image.")
	}
}
