package controllers

import (
	"fmt"
	"net/http"
)

type HelloController struct {
}

func AddHelloControllerRoutes() {
	controller := &HelloController{}

	controller.AddHelloRoute()
}

func (c *HelloController) AddHelloRoute() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}

func (c *HelloController) AddHelloIdRoute() {
	http.HandleFunc("/hello/:id", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(
			"Hello, World-%s!",
			r.URL.Query().Get("id"),
		)))
	})
}
