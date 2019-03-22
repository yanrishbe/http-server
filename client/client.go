package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:9093/time/minsk")
	resp1, err1 := http.Get("http://localhost:9093/time/moscow")
	if err != nil || err1 != nil {
		log.Fatalln("Error getting time")
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln()
		}
	}()
	defer
		func() {
			err := resp1.Body.Close()
			if err != nil {
				log.Fatalln()
			}
		}()
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
}
