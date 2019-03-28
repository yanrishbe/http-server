package execute

import (
	"strings"
	"time"
)

func TimeByLocation(location string) (string, error) {
	t := time.Now()
	switch strings.ToLower(location) {
	case "moscow":
		newLocation, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			return "", err
		}
		tM := t.In(newLocation)
		return tM.Format(time.RFC1123Z), nil

	case "minsk":
		return t.Format(time.RFC1123Z), nil
	default:
		t := t.UTC()
		return "wrong input, sending UTC time instead: " + t.Format(time.RFC1123Z), nil

	}
}
