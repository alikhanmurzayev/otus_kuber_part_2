package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	time.Sleep(time.Second * 2)

	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write([]byte(`{"status": "ok"}`))
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		uri := request.RequestURI
		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write([]byte(fmt.Sprintf("Requested url: %s", uri)))
		if err != nil {
			panic(err)
		}
	})

	log.Fatalln(http.ListenAndServe(":80", nil))
}
