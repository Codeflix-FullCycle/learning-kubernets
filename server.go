package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var createAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/myFamily", GetFamily)
	http.HandleFunc("/secret", Secret)

	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s", name, age)
}
func Secret(w http.ResponseWriter, r *http.Request) {

	user := os.Getenv("USER")
	password := os.Getenv("password")
	fmt.Fprintf(w, "User %s. Password %s", user, password)
}
func GetFamily(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/myFamily/family.txt")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "My family: %s", string(data))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(createAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
