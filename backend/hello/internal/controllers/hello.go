package controllers

import "net/http"

type HelloController struct {
}

func AddHelloControllerRoutes() {
	controller := &HelloController{}

	http.HandleFunc("/hello", controller.Hello)
}

func (c *HelloController) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
