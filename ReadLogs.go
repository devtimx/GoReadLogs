package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	dir := "/var/log/"
	file := ListLogs(dir)
	for i, name := range file {
		fmt.Println(i, dir+name)
	}

	route := mux.NewRouter()

	route.HandleFunc("/readlog", ReadLog).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(route)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func ListLogs(ruta string) []string {
	var name []string
	archivos, err := ioutil.ReadDir(ruta)
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {
		suffix := filepath.Ext(archivo.Name())
		if !archivo.IsDir() && suffix == ".log" {
			name = append(name, archivo.Name())

		}
	}
	return name
}

func ReadLog(w http.ResponseWriter, r *http.Request) {
	Ruta := r.URL.Query().Get("ruta")

	files, err := ioutil.ReadFile(Ruta)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(files)
}
