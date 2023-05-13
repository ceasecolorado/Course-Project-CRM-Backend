package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// define a global variable for the content type header
var contentType string = "Content-Type"
var contentTypeJSON string = "application/json"

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type errorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

var customers = []Customer{}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeJSON)
	id := mux.Vars(r)["id"]
	for _, customer := range customers {
		if customer.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	error := errorResponse{
		StatusCode: http.StatusNotFound,
		Message:    "Customer not found",
	}
	json.NewEncoder(w).Encode(error)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeJSON)
	w.WriteHeader(http.StatusCreated)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
}
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeJSON)
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Initialize mock data
	customers = append(customers,
		Customer{
			ID:        "1",
			Name:      "John Doe",
			Role:      "CEO",
			Email:     "joDoe@email.com",
			Phone:     5550199,
			Contacted: true,
		},
		Customer{
			ID:        "2",
			Name:      "Jane Doe",
			Role:      "CTO",
			Email:     "jaDoe@email.com",
			Phone:     5550199,
			Contacted: false,
		},
		Customer{
			ID:        "3",
			Name:      "Cesar Colorado",
			Role:      "Software Engineer",
			Email:     "cesar@mail.com",
			Phone:     555555,
			Contacted: true,
		})
	// Init a new router by invoking the "NewRouter" handler
	router := mux.NewRouter()

	// Define the URl constant instead of duplicating the string
	customerUri := "/customers"
	customerIdUri := "/customers/{id}"

	// Define the routes
	router.HandleFunc(customerUri, getCustomers).Methods("GET")
	router.HandleFunc(customerIdUri, getCustomer).Methods("GET")

	router.HandleFunc(customerUri, addCustomer).Methods("POST")
	router.HandleFunc(customerIdUri, deleteCustomer).Methods("DELETE")
	router.HandleFunc(customerIdUri, updateCustomer).Methods("PUT")

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
