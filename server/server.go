package server

import "net/http"

type Country struct {
	Name     string
	Language string
}

// var countries := []Country{Country{name: "Honduras", language: "Español"}}
var countries []*Country = []*Country{} //&Country{Name: "Honduras", Language: "Español"}

func New(address string) *http.Server {
	initRoutes()

	return &http.Server{Addr: address}
}
