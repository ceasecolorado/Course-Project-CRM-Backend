package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	// other imports
	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
)

// define a global variable for the content type header
var contentType string = "Content-Type"
var contentTypeJSON string = "application/json"

// define const for random string generation
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// define the Customer and error structs
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

var customerNotFound = errorResponse{
	StatusCode: http.StatusNotFound,
	Message:    "Customer not found",
}

var customers = []Customer{}

// Function to generate a random string
func generateRandomStringID(n int) string {
	// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// Functions to handle all requests
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
	json.NewEncoder(w).Encode(customerNotFound)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeJSON)

	var newCustomer Customer

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newCustomer)

	newCustomer.ID = generateRandomStringID(10)
	customers = append(customers, newCustomer)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeJSON)
	id := mux.Vars(r)["id"]
	var updatedCustomer Customer

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedCustomer)

	for i, customer := range customers {
		if customer.ID == id {
			customers[i] = updatedCustomer
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedCustomer)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(customerNotFound)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, contentTypeJSON)
	id := mux.Vars(r)["id"]

	for i, customer := range customers {
		if customer.ID == id {
			customers = slices.Delete(customers, i, i+1)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customers)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(customerNotFound)
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
	fmt.Printf("customer's type is: %T\n", customers)
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
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Start the server
	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
