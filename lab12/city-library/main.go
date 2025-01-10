package main

import (
	"ccproject/handler"
	"ccproject/service"
	"ccproject/store"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
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

	bookStore := store.NewBooksMongoDBStore(mongoClient)

	bookService, err := service.NewBooksService(*bookStore)
	handleErr(err)

	bookHandler := handler.NewBookHandler(&bookService)

	r := mux.NewRouter()

	r.HandleFunc("/register", bookHandler.Register).Methods("POST")
	r.HandleFunc("/borrow", bookHandler.BorrowBook).Methods("POST")
	r.HandleFunc("/", bookHandler.GetAll).Methods("GET")
	r.HandleFunc("/{userId}/{isbn}", bookHandler.Return).Methods("DELETE")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
	}
	log.Fatal(srv.ListenAndServe())
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
