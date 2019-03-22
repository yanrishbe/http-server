package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := time.Now()
	r.FormValue("location")
	if location := vars["key"]; location != "" {
		switch location {
		case "moscow":
			newLocation, err := time.LoadLocation("Europe/Moscow")
			if err != nil {
				log.Println("Error in 'moscow' case adding location: ", err)
			}
			tM := t.In(newLocation)
			currTime := tM.Format(time.RFC1123Z)
			_, errWrite := w.Write([]byte("Moscow:" + currTime))
			if errWrite != nil {
				log.Println("Error in 'minsk' case writing data: ", err)
			}

		case "minsk":
			currTime := t.Format(time.RFC1123Z)
			_, err := w.Write([]byte("Minsk:" + currTime))
			if err != nil {
				log.Println("Error in 'minsk' case writing data: ", err)
			}
		default:
			t := t.UTC()
			defTime := t.Format(time.RFC1123Z)
			_, err := w.Write([]byte("Wrong input, sending UTC time instead: " + defTime))
			if err != nil {
				log.Println("Error in default: ", err)
			}

		}
	} else {
		_, err := w.Write([]byte("There is now input"))
		if err != nil {
			log.Println("Input error: ", err)
		}
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/time/{key}", handler)
	log.Fatal(http.ListenAndServe(":9093", router))
}
