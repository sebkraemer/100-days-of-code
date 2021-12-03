package main

import (
	"log"
	"net/http"
)

func main() {
	hw := func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello #100DaysOfCode friends 👋"))
	}
	http.HandleFunc("/", hw)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
