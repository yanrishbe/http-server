package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for _, value := range vars {
		switch value {
		case "moscow":
			newLocation, err := time.LoadLocation("Europe/Moscow")
			if err != nil {
				log.Println("Error in 'moscow' case adding location: ", err)
			}
			tM := time.Now().In(newLocation)
			currTime := tM.Format(time.RFC1123Z)
			_, errWrite := w.Write([]byte("Moscow:" + currTime))
			if errWrite != nil {
				log.Println("Error in 'minsk' case writing data: ", err)
			}

		case "minsk":
			t := time.Now()
			currTime := t.Format(time.RFC1123Z)
			_, err := w.Write([]byte("Minsk:" + currTime))
			if err != nil {
				log.Println("Error in 'minsk' case writing data: ", err)
			}
		default:
			t := time.Now().UTC()
			defTime := t.Format(time.RFC1123Z)
			_, err := w.Write([]byte("Wrong input, sending UTC time instead: " + defTime))
			if err != nil {
				log.Println("Error in default: ", err)
			}

		}

	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/time/{key}", handler)
	err := http.ListenAndServe(":9093", router)
	if err != nil {
		panic(err)
	}
}
