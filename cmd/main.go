package main

import (
    "log"
    "net/http"
    "paypal-clone/controllers"
    "paypal-clone/db"
    "paypal-clone/middleware"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    
    db.Connect()

    // Auth routes
    router.HandleFunc("/register", controllers.Register).Methods("POST")
    router.HandleFunc("/login", controllers.Login).Methods("POST")
    router.HandleFunc("/verify", controllers.VerifyEmail).Methods("POST")
    
    // Password reset routes
    router.HandleFunc("/request-password-reset", controllers.RequestPasswordReset).Methods("POST")
    router.HandleFunc("/reset-password", controllers.ResetPassword).Methods("POST")
    
    // Secure routes
    api := router.PathPrefix("/api").Subrouter()
    api.Use(middleware.AuthMiddleware)
    
    // User routes
    api.HandleFunc("/users", controllers.CreateUser).Methods("POST")
    
    // Transaction routes
    api.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")
    
    // Balance routes
    api.HandleFunc("/balance", controllers.GetBalance).Methods("GET")

    // Currency conversion route
    api.HandleFunc("/convert-currency", controllers.ConvertCurrency).Methods("POST")
    
    // Stripe payment route
    api.HandleFunc("/create-payment", controllers.CreatePayment).Methods("POST")
    
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}