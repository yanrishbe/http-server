package main

import (
	"encoding/json"
	"http-project/entities"
	"log"
	"net/http"

	"http-project/server/execute"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if location := vars["key"]; location != "" {
		timeInLocation, errLocation := execute.TimeByLocation(location)
		if errLocation != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error in "+location+"case loading time: ", errLocation)
			return
		}
		_, errWrite := w.Write([]byte(location + ": " + timeInLocation))
		if errWrite != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error in"+location+"case writing data: ", errWrite)
			return
		}

	} else {
		_, err := w.Write([]byte("There is no input"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Input error: ", err)
			return
		}
	}
}

func handlerQuery(w http.ResponseWriter, r *http.Request) {
	if location := r.FormValue("location"); location != "" {
		timeInLocation, errLocation := execute.TimeByLocation(location)
		if errLocation != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error in "+location+"case loading time: ", errLocation)
			return
		}
		_, errWrite := w.Write([]byte(location + ": " + timeInLocation))
		if errWrite != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error in"+location+"case writing data: ", errWrite)
			return
		}
	} else {
		_, err := w.Write([]byte("There is no input"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Input error: ", err)
			return
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
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Decoding client's input error: ", err)
		return
	}
	timeInLocation, errLocation := execute.TimeByLocation(dataFromClient.City)
	if errLocation != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error in "+dataFromClient.City+"case loading time: ", errLocation)
		return
	}
	dataFromClient.Time = timeInLocation
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if errWrite := json.NewEncoder(w).Encode(dataFromClient); errWrite != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("(POST) Error in"+dataFromClient.City+"case encoding data: ", errWrite)
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/time", handlerQuery).Methods(http.MethodGet)
	router.HandleFunc("/time", handlerPost).Methods(http.MethodPost)
	router.HandleFunc("/time/{key}", handler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":9093", router))
}
