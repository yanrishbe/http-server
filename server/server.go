package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/yanrishbe/http-server/entities"

	"github.com/gorilla/mux"
	"github.com/yanrishbe/http-server/server/execute"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := time.Now()
	if location := vars["key"]; location != "" {
		execute.ManageLocation(location, t, w)
	} else {
		_, err := w.Write([]byte("There is no input"))
		if err != nil {
			log.Println("Input error: ", err)
		}
	}

}

func handlerQuery(w http.ResponseWriter, r *http.Request) {
	if location := r.FormValue("location"); location != "" {
		t := time.Now()
		execute.ManageLocation(location, t, w)
	} else {
		_, err := w.Write([]byte("There is no input"))
		if err != nil {
			log.Println("Input error: ", err)
		}
	}

}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	dataFromClient := new(entities.Data)
	err := json.NewDecoder(r.Body).Decode(&dataFromClient)
	defer func() {
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		log.Println("Decoding client's input error: ", err)
	}
	t := time.Now()
	execute.ManageLocationPost(dataFromClient.City, t, w)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/time", handlerQuery).Methods(http.MethodGet)
	router.HandleFunc("/time", handlerPost).Methods(http.MethodPost)
	router.HandleFunc("/time/{key}", handler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":9093", router))
}
