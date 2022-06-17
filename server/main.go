package main

import (
	"log"
	"net/http"
	"server/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)



func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Login).Methods("POST")
	r.HandleFunc("/signup",controllers.Signup).Methods("POST")
	r.HandleFunc("/forgetpassword",controllers.Forgetpassword).Methods("POST")
	r.HandleFunc("/products", controllers.Getproducts).Methods("GET")
	r.HandleFunc("/getoneproducts/{id}", controllers.Getoneproducts).Methods("GET")
	r.HandleFunc("/updateproduct/{id}", controllers.Updateproduct).Methods("PUT")
	r.HandleFunc("/addproduct", controllers.Addproduct).Methods("POST")
	r.HandleFunc("/deleteproduct/{id}", controllers.Deleteproduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}