package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", Name).Methods("GET")
	router.HandleFunc("/bad", Bad).Methods("GET")
	router.HandleFunc("/data", Data).Methods("POST")
	router.HandleFunc("/header", Header).Methods("GET")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func Name(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	fmt.Fprintf(w, "Hello, %s!", param)
}

func Bad(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Status: 500", http.StatusInternalServerError)
}

func Data(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(w, "I got message:\n%s", string(bytes))
}

func Header(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("A"))
	if err != nil {
		log.Println(err)
		return
	}

	b, err := strconv.Atoi(r.Header.Get("B"))
	if err != nil {
		log.Println(err)
		return
	}
	sum := a + b
	result := fmt.Sprint(sum)
	w.Header().Set("a+b", result)
}
