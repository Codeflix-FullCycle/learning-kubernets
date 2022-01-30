package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/myFamily", GetFamily)

	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s", name, age)
}

func GetFamily(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/myFamily/family.txt")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "My family: %s", string(data))
}
