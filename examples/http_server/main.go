package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	aerostack "github.com/aerostackdev/sdks/packages/go"
	"github.com/aerostackdev/sdks/packages/go/pkg/models/operations"
)

/*
 * HTTP Server Example
 *
 * Demonstrates usage with standard library net/http
 */

var sdk *aerostack.SDK

func main() {
	// Initialize SDK
	// apiKey := os.Getenv("AEROSTACK_API_KEY")
	sdk = aerostack.New(
	// aerostack.WithSecurity(shared.Security{APIKeyAuth: &apiKey}),
	)

	http.HandleFunc("/signup", handleSignup)
	http.HandleFunc("/users", handleUsers)

	port := ":8080"
	fmt.Printf("Server starting on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	res, err := sdk.Authentication.AuthSignup(ctx, operations.AuthSignupRequestBody{
		Email:    req.Email,
		Password: req.Password,
		Name:     &req.Name,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res.StatusCode == 201 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res.AuthResponse)
	} else {
		http.Error(w, "Signup failed", http.StatusBadRequest)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Example DB Query
	res, err := sdk.Database.DbQuery(ctx, operations.DbQueryRequestBody{
		SQL: "SELECT * FROM users LIMIT 10",
	}, nil, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if res != nil && res.DbQueryResult != nil {
		json.NewEncoder(w).Encode(res.DbQueryResult)
	} else {
		json.NewEncoder(w).Encode(map[string]any{"results": []any{}})
	}
}
