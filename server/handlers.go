package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getCountries(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("content-type", "application/json")
	json.NewEncoder(res).Encode(countries)
	// fmt.Fprintf(rw, "%v", countries)
}

func addCountry(res http.ResponseWriter, req *http.Request) {
	country := &Country{}
	err := json.NewDecoder(req.Body).Decode(country)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "%v", err)
		return
	}
	countries = append(countries, country)
	res.WriteHeader(http.StatusCreated)

	myRes := make(map[string]string)
	myRes["msg"] = "Country creado con exito"
	json.NewEncoder(res).Encode(myRes)
}

func index(rw http.ResponseWriter, r *http.Request) {
	// rw ResponseWriter,
	// r *Request
	name := "Post"
	if r.Method == http.MethodGet {
		name = "GET"
		rw.WriteHeader(http.StatusOK)
	} else {

		rw.WriteHeader(http.StatusCreated)
	}
	fmt.Fprintf(rw, "Hola Mundo! %s", name)
}
