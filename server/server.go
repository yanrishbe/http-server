package main

import (
	//"encoding/json"
	"http-project/server/execute"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

////// implements http.Handler interface///////////
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
func main() {

	//err := json.NewEncoder(w).Encode(todos)
	router := mux.NewRouter()
	router.HandleFunc("/time", handlerQuery).Methods(http.MethodPost, "GET")
	router.HandleFunc("/time/{key}", handler).Methods("POST", "GET")
	log.Fatal(http.ListenAndServe(":9093", router))
}
