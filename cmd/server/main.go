package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/octiq-test/go-cart-service/pkg/cart"
)

func main() {
	c := cart.New()
	c.Add("sku-1", 2, 1999)

	http.HandleFunc("/cart", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]int{"items": c.Count(), "total_cents": c.TotalCents()})
	})
	log.Println("cart service on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
