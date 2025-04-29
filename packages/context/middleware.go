package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// using context with middlewareAuth

func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.FormValue("userId")
		password := r.FormValue("password")

		// no db configured here
		userName, dbPassword := database.getUserData(userId)
		if userName == "" || password != dbPassword {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "userName", userName)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// userId := r.FormValue("userId")
	// username, _ := database.getUserData(userId)
	userName := r.Context().Value("userName").(string)
	fmt.Printf(w, "Hello %s!", userName)
}

func apiHandler() {
	http.HandleFunc("/hello", authenticate(helloHandler))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
