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
	router.HandleFunc("/bad", ViewBad).Methods(http.MethodGet)
	router.HandleFunc("/name/{param}", ViewName).Methods(http.MethodGet)
	router.HandleFunc("/data", ViewPostParam).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func ViewBad(w http.ResponseWriter, r *http.Request) {

	http.Error(w, "bad", 500)
	return
}

func ViewName(w http.ResponseWriter, r *http.Request) {

	param, _ := mux.Vars(r)["param"] // it'll be of type string,

	resp := []byte("Hello, " + param + "!")

	w.Write(resp)
	return
}

func ViewPostParam(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	resp := []byte("I got message:\n" + string(body))

	w.Write(resp)
	return
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
