package execute

import (
	"log"
	"net/http"
	"time"
)

func ManageLocation(location string, t time.Time, w http.ResponseWriter) {
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
}
