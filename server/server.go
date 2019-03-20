package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)
type timeFormat struct{
	hour int
	minute int
	second int
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for _, value := range vars{
		switch value {
		case "moscow":
			location, err := time.LoadLocation("Europe/Moscow")
			if err != nil {
				log.Println("Error in 'moscow' case adding location: ", err)
			}
			tM := time.Time{}.In(location)
			currTime := timeFormat{tM.Hour(), tM.Minute(), tM.Second(),}
			timeNow := strconv.Itoa(currTime.hour) + ":" + strconv.Itoa(currTime.minute) + ":" + strconv.Itoa(currTime.second)
			_, errWrite := w.Write([]byte("Moscow:" + timeNow))
			if errWrite != nil {
				log.Println("Error in 'minsk' case writing data: ", err)
			}


		case "default":
			_, err := w.Write([]byte("Wrong input, sending Minsk time instead: \n"))
			if err != nil {
				log.Println("Error in default: ", err)
			}
			fallthrough
		case "minsk":
			t:= time.Now()
			currTime := timeFormat{t.Hour(), t.Minute(), t.Second(),}
			timeNow := strconv.Itoa(currTime.hour) + ":" + strconv.Itoa(currTime.minute) + ":" + strconv.Itoa(currTime.second)
			_, err := w.Write([]byte("Minsk:" + timeNow))
			if err != nil {
				log.Println("Error in 'minsk' case writing data: ", err)
			}

		}

	}

}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/time/{key}", handler)
	http.ListenAndServe(":9093", router)
}