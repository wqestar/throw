package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("qwe")	

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is empty")
	}
	fmt.Println("port:", portString)

	router := chi.NewRouter()	
	router.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{"https://*","http://*"},
			AllowedMethods: []string{"GET", "POST", "DELETE", "UPDATE", "OPTIONS"},
			AllowedHeaders: []string{"*"},
			ExposedHeaders: []string{"Link"},
			AllowCredentials: false,
			MaxAge: 300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/ready", handlerReadiness)

	v1Router.Get("/err", handlerError)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr : ":"+portString,

	}
	log.Printf("server start in porn %v",portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
