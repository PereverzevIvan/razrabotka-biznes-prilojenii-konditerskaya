package controllers

import (
	"net/http"
)

func Init(mux *http.ServeMux) {
	AddHelloControllerRoutes(mux)
}
