package main

import (
	"log"
	"net/http"

	"github.com/sebasromero/api/authorization"
	"github.com/sebasromero/api/handler"
	"github.com/sebasromero/api/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Could not get the certificates: %v", err)
	}
	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)
	log.Println("Listen in port: 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error in server: %v\n", err)
	}
}
