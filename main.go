package main

import (
	"github.com/gorilla/mux"
	"github.com/sirsean/textback/config"
	"github.com/sirsean/textback/redis"
	"github.com/sirsean/textback/run"
	"github.com/sirsean/textback/sms"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	log.Printf("Starting Up")

	rand.Seed(time.Now().UnixNano())

	redis.Connect()
	sms.Connect()

	go run.Start()

	mainRouter := mux.NewRouter()

	smsRouter := mux.NewRouter()
	smsRouter.HandleFunc("/sms/inbound", sms.Inbound).Methods("POST")
	mainRouter.Handle("/sms/{.*}", smsRouter)

	port := config.Get().Host.Port
	log.Printf("Serving on port %v", port)
	log.Fatal(http.ListenAndServe(port, mainRouter))
}
