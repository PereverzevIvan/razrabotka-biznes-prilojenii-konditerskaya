package controllers

import "net/http"

type GoodbyeController struct {
}

func AddGoodbyeControllerRoutes() {
	controller := &GoodbyeController{}

	controller.AddGoodbyeRoute()
}

func (c *GoodbyeController) AddGoodbyeRoute() {
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Goodbye, World!"))
	})
}
