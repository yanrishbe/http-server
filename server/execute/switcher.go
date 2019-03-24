package execute

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/yanrishbe/http-server/entities"
)

func ManageLocation(location string, t time.Time, w http.ResponseWriter) {
	switch strings.ToLower(location) {
	case "moscow":
		newLocation, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			log.Println("Error in 'moscow' case adding location: ", err)
		}
		tM := t.In(newLocation)
		currTime := tM.Format(time.RFC1123Z)
		_, errWrite := w.Write([]byte("Moscow:" + currTime))
		if errWrite != nil {
			log.Println("Error in 'moscow' case writing data: ", err)
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
}

func ManageLocationPost(location string, t time.Time, w http.ResponseWriter) {
	switch strings.ToLower(location) {
	case "moscow":
		newLocation, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			log.Println("Error in 'moscow' case adding location: ", err)
		}
		tM := t.In(newLocation)
		currTime := tM.Format(time.RFC1123Z)
		moscowData := new(entities.Data)
		moscowData.City = "Moscow"
		moscowData.Time = currTime
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if errWrite := json.NewEncoder(w).Encode(moscowData); errWrite != nil {
			log.Println("(POST) Error in 'moscow' case encoding data: ", errWrite)
		}

	case "minsk":
		currTime := t.Format(time.RFC1123Z)
		minskData := new(entities.Data)
		minskData.City = "Minsk"
		minskData.Time = currTime
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if errWrite := json.NewEncoder(w).Encode(minskData); errWrite != nil {
			log.Println("(POST) Error in 'minsk' case encoding data: ", errWrite)
		}
	default:
		t := t.UTC()
		defTime := t.Format(time.RFC1123Z)

		defData := new(entities.Data)
		defData.City = "Wrong input: UTC time"
		defData.Time = defTime
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if errWrite := json.NewEncoder(w).Encode(defData); errWrite != nil {
			log.Println("(POST) Error in default encoding data: ", errWrite)
		}

	}
}
