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

	//err := sms.Send("7735048753", "this is a test message")
	//log.Printf("sms err %v", err)
	//return

	//mongo.Connect()

	//go store.StartCleaning()

	mainRouter := mux.NewRouter()
	//router.HandleFunc("/", controller.Index).Methods("GET")
	//router.HandleFunc("/doc/{id}", controller.ShowDocument).Methods("GET")

	//router.HandleFunc("/api/docs", api.ListDocuments).Methods("GET")

	smsRouter := mux.NewRouter()
	smsRouter.HandleFunc("/sms/inbound", sms.Inbound).Methods("POST")
	mainRouter.Handle("/sms/{.*}", smsRouter)

	// TODO see adescaper-adserver for router tips

	port := config.Get().Host.Port
	log.Printf("Serving on port %v", port)
	log.Fatal(http.ListenAndServe(port, mainRouter))
}
