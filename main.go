package main

import (
	"fmt"
	"go-contacts/app"
	"go-contacts/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/customers/new", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/customers/delete", controllers.DeteleCustomerFor).Methods("DELETE")
	router.HandleFunc("/api/list/new", controllers.CreateList).Methods("POST")
	router.HandleFunc("/api/producttolist/new", controllers.AddList).Methods("POST")
	router.HandleFunc("/api/producttolist/delete", controllers.RemoveListFor).Methods("DELETE")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
