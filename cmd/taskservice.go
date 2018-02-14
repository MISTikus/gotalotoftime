package main

import (
	"alotoftime/taskservice"
	"net/http"

	"log"

	"github.com/julienschmidt/httprouter"
)

const apiroute = "/api/"

func main() {
	service := taskservice.NewService()
	router := httprouter.New()
	for _, r := range service.Routes {
		router.Handle(r.Method, apiroute+r.Route, r.Handler)
	}

	log.Println("Task service started ...")

	log.Fatal(http.ListenAndServe(":9091", router))

	log.Println("Task service stopped .")
}
