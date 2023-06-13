package main

//go mod init example.com/api
//go get -u github.com/gorilla/mux
// :wq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Operacion struct {
	Left  int `json:"left"`
	Right int `json:"right"`
}
type Numero struct {
	Value int `json:"value"`
}

func suma(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var operacion Operacion
	json.Unmarshal(reqBody, &operacion)
	result := strconv.Itoa(operacion.Left + operacion.Right)
	fmt.Fprintf(w, result)
}

func resta(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var operacion Operacion
	json.Unmarshal(reqBody, &operacion)
	result := strconv.Itoa(operacion.Left - operacion.Right)
	fmt.Fprintf(w, result)
}

func primo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var numero Numero
	json.Unmarshal(reqBody, &numero)
	var result string
	if numero.Value < 2 {
		result = strconv.Itoa(numero.Value) + " no es primo"
	} else {
		limite := int(math.Sqrt(float64(numero.Value)))

		for i := 2; i <= limite; i++ {
			if numero.Value%i == 0 {
				result = strconv.Itoa(numero.Value) + " no es primo"
				break
			} else if numero.Value%i != 0 {
				result = strconv.Itoa(numero.Value) + " es primo"
				break
			}
		}
	}
	fmt.Fprintf(w, result)
}

func main() {
	fmt.Println("Api corriendo en 'http://localhost:8000")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/suma", suma).Methods("POST")
	router.HandleFunc("/resta", resta).Methods("POST")
	router.HandleFunc("/primo", primo).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
