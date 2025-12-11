package ecommerce

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/devrevolt/qubit-memory"
	"github.com/devrevolt/qubit-service"
)

// HandleCatalogAPI handles catalog service API requests
func HandleCatalogAPI(w http.ResponseWriter, r *http.Request, path string, endpoint service.Endpoint, mem *memory.MemoryManager) {
	switch endpoint.Name {
	case "listProducts":
		// Get products from memory
		products := getDemoProducts()

		// Apply filters from query params
		category := r.URL.Query().Get("category")
		if category != "" {
			filtered := []map[string]interface{}{}
			for _, p := range products {
				if p["category"] == category {
					filtered = append(filtered, p)
				}
			}
			products = filtered
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"products": products,
			"total":    len(products),
			"page":     1,
			"limit":    20,
			"hasMore":  false,
		})

	case "getProduct":
		id := extractPathParam(path, "/products/{id}", "id")
		products := getDemoProducts()
		for _, p := range products {
			if p["id"] == id {
				json.NewEncoder(w).Encode(p)
				return
			}
		}
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "product not found"})

	case "searchProducts":
		query := r.URL.Query().Get("q")
		products := getDemoProducts()
		filtered := []map[string]interface{}{}
		for _, p := range products {
			name := p["name"].(string)
			if strings.Contains(strings.ToLower(name), strings.ToLower(query)) {
				filtered = append(filtered, p)
			}
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"products": filtered,
			"total":    len(filtered),
		})

	case "getCategories":
		json.NewEncoder(w).Encode([]map[string]interface{}{
			{"id": "electronics", "name": "Electronics", "slug": "electronics"},
			{"id": "clothing", "name": "Clothing", "slug": "clothing"},
			{"id": "home", "name": "Home & Garden", "slug": "home"},
		})

	default:
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Catalog endpoint: " + endpoint.Name})
	}
}

// HandleCartAPI handles cart service API requests
func HandleCartAPI(w http.ResponseWriter, r *http.Request, path string, endpoint service.Endpoint, mem *memory.MemoryManager) {
	switch endpoint.Name {
	case "getCart":
		userId := extractPathParam(path, "/cart/{userId}", "userId")
		// Return empty cart for demo
		json.NewEncoder(w).Encode(map[string]interface{}{
			"userId": userId,
			"items":  []interface{}{},
			"total":  0,
		})
	default:
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Cart endpoint: " + endpoint.Name})
	}
}

// HandleOrderAPI handles order service API requests
func HandleOrderAPI(w http.ResponseWriter, r *http.Request, path string, endpoint service.Endpoint, mem *memory.MemoryManager) {
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Order endpoint: " + endpoint.Name})
}

// getDemoProducts returns sample product data
func getDemoProducts() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":          "prod-001",
			"name":        "Wireless Headphones",
			"description": "High-quality wireless headphones with noise cancellation",
			"price":       149.99,
			"currency":    "USD",
			"category":    "electronics",
			"images":      []string{"/images/headphones.jpg"},
			"stock":       50,
			"sku":         "WH-001",
		},
		{
			"id":          "prod-002",
			"name":        "Smart Watch",
			"description": "Feature-rich smartwatch with health monitoring",
			"price":       299.99,
			"currency":    "USD",
			"category":    "electronics",
			"images":      []string{"/images/smartwatch.jpg"},
			"stock":       30,
			"sku":         "SW-002",
		},
		{
			"id":          "prod-003",
			"name":        "Cotton T-Shirt",
			"description": "Comfortable cotton t-shirt, available in multiple colors",
			"price":       29.99,
			"currency":    "USD",
			"category":    "clothing",
			"images":      []string{"/images/tshirt.jpg"},
			"stock":       100,
			"sku":         "TS-003",
		},
		{
			"id":          "prod-004",
			"name":        "LED Desk Lamp",
			"description": "Adjustable LED desk lamp with multiple brightness levels",
			"price":       59.99,
			"currency":    "USD",
			"category":    "home",
			"images":      []string{"/images/lamp.jpg"},
			"stock":       75,
			"sku":         "DL-004",
		},
	}
}

// SeedDemoData populates memory with demo data
func SeedDemoData(mem *memory.MemoryManager) {
	products := getDemoProducts()
	for _, p := range products {
		mem.Store(p["id"].(string), p)
	}
}

// Helper: extract path parameter from path using pattern
func extractPathParam(path, pattern, paramName string) string {
	parts := strings.Split(path, "/")
	patt := strings.Split(pattern, "/")
	for i, p := range patt {
		if p == "{"+paramName+"}" && i < len(parts) {
			return parts[i]
		}
	}
	return ""
}

// writeJSON is expected to be in main package utils, but implement a local fallback for the ecommerce package
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
