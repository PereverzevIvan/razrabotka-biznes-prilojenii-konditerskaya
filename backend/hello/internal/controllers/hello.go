package controllers

import "net/http"

type HelloController struct {
}

func AddHelloControllerRoutes(mux *http.ServeMux) {
	controller := &HelloController{}

	mux.HandleFunc("/hello", controller.Hello)
}

func (c *HelloController) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
