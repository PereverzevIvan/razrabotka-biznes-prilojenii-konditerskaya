package controllers

import (
	"fmt"
	"net/http"
)

func Init(port int) {
	AddHelloControllerRoutes()

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
