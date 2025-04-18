package main

import (
	"fmt"
	"net/http"
	
	"github.com/go-chi/chi"
	"github.com/avukadin/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Printf("Starting GO API Services...")
	fmt.Println("Let's GO")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil{
		log.Error(err)
	}
}
