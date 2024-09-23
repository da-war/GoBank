package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

// User represents a user in the system
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Product represents a product in the store
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Order represents a user's order
type Order struct {
	ID       int      `json:"id"`
	UserID   int      `json:"user_id"`
	ProductID int     `json:"product_id"`
	Quantity int      `json:"quantity"`
}

// Store simulates an in-memory database
var store = struct {
	sync.RWMutex
	users    []User
	products []Product
	orders   []Order
}{
	users:    []User{},
	products: []Product{},
	orders:   []Order{},
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/orders", ordersHandler)
	http.ListenAndServe(":8080", nil)
}

// usersHandler handles requests to the /users endpoint
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w)
	case http.MethodPost:
		createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// productsHandler handles requests to the /products endpoint
func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getProducts(w)
	case http.MethodPost:
		createProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ordersHandler handles requests to the /orders endpoint
func ordersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getOrders(w)
	case http.MethodPost:
		createOrder(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getUsers retrieves all users
func getUsers(w http.ResponseWriter) {
	store.RLock()
	defer store.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.users)
}

// createUser adds a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	store.Lock()
	newUser.ID = len(store.users) + 1
	store.users = append(store.users, newUser)
	store.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// getProducts retrieves all products
func getProducts(w http.ResponseWriter) {
	store.RLock()
	defer store.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.products)
}

// createProduct adds a new product
func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	store.Lock()
	newProduct.ID = len(store.products) + 1
	store.products = append(store.products, newProduct)
	store.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

// getOrders retrieves all orders
func getOrders(w http.ResponseWriter) {
	store.RLock()
	defer store.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.orders)
}

// createOrder adds a new order
func createOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	store.Lock()
	newOrder.ID = len(store.orders) + 1
	store.orders = append(store.orders, newOrder)
	store.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}
