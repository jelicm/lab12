package main

import (
	"log"
	"context"
	"net/http"
	"ccmainproject/handler"
	"ccmainproject/store"
	"ccmainproject/service"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gorilla/mux"
	"os"
)

func main() {

	port := os.Getenv("MAIN_DB_PORT")
	host := os.Getenv("MAIN_DB_HOST")

	mongoClient, err := store.GetClient(host, port)
	if err != nil {
		log.Fatal(err)
	}

	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			log.Printf("error closing db: %s\n", err)
		}
	}(mongoClient, context.Background())

	userStore := store.NewUsersMongoDBStore(mongoClient)

	userService, err := service.NewUsersService(*userStore)
	handleErr(err)

	userHandler := handler.NewUserHandler(&userService)

	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/{userId}", userHandler.Borrow).Methods("GET")
	r.HandleFunc("/", userHandler.GetAll).Methods("GET")
	r.HandleFunc("/return/{userId}", userHandler.Return).Methods("GET")


	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	log.Fatal(srv.ListenAndServe())
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
