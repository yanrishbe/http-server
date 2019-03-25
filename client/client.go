package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"http-project/entities"
)

func GetData(r *http.Response) {
	dataFromServer := new(entities.Data)
	err := json.NewDecoder(r.Body).Decode(&dataFromServer)
	defer func() {
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		log.Println("Decoding server's output error: ", err)
	}
	log.Println(string(dataFromServer.City + ": " + dataFromServer.Time))

}

func main() {

	resp, err := http.Get("http://localhost:9093/time/minsk")
	resp1, err1 := http.Get("http://localhost:9093/time/moscow")

	if err != nil || err1 != nil {
		log.Fatalln("Error getting time")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(data))

	data1, err1 := ioutil.ReadAll(resp1.Body)
	if err1 != nil {
		panic(err1)
	}
	log.Println(string(data1))

	jsonMinsk := []byte(`
{
    "city": "minsk"
}`)

	jsonMoscow := []byte(`
{
	"city": "moscow"
}`)

	respPostMinsk, errMinsk := http.Post("http://localhost:9093/time", "application/json",
		bytes.NewBuffer(jsonMinsk))
	respPostMoscow, errMoscow := http.Post("http://localhost:9093/time", "application/json",
		bytes.NewBuffer(jsonMoscow))

	if errMinsk != nil || errMoscow != nil {
		log.Fatalln("Error getting time by POST method")
	}

	GetData(respPostMinsk)
	GetData(respPostMoscow)

	defer func() {
		err := resp.Body.Close()
		errM := resp1.Body.Close()
		if err != nil || errM != nil {
			log.Fatalln()
		}
	}()

}
