package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/countries", switchHandler)
}

func switchHandler(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getCountries(rw, r)

	case http.MethodPost:
		addCountry(rw, r)

	default:
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(rw, "Not found")
		return
	}

}
