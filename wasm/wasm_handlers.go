package main

import (
	"encoding/json"
	"strings"
	"unsafe"
)

//export handleCatalogAPI
func handleCatalogAPI(inputPtr, inputLen uint32) (uint32, uint32) {
	inputBytes := make([]byte, inputLen)
	src := (*[1 << 30]byte)(unsafe.Pointer(uintptr(inputPtr)))[:inputLen:inputLen]
	copy(inputBytes, src)

	var input map[string]string
	json.Unmarshal(inputBytes, &input)

	path := input["path"]
	method := input["method"]
	body := input["body"]

	result := handleCatalogAPIImpl(path, method, body)
	resultBytes := []byte(result)

	resultPtr := uint32(1024) // fixed offset for demo
	copy((*[1 << 30]byte)(unsafe.Pointer(uintptr(resultPtr)))[:len(resultBytes)], resultBytes)

	return resultPtr, uint32(len(resultBytes))
}

//export handleCartAPI
func handleCartAPI(inputPtr, inputLen uint32) (uint32, uint32) {
	inputBytes := make([]byte, inputLen)
	src := (*[1 << 30]byte)(unsafe.Pointer(uintptr(inputPtr)))[:inputLen:inputLen]
	copy(inputBytes, src)

	var input map[string]string
	json.Unmarshal(inputBytes, &input)

	path := input["path"]
	method := input["method"]
	body := input["body"]

	result := handleCartAPIImpl(path, method, body)
	resultBytes := []byte(result)

	resultPtr := uint32(2048) // fixed offset
	copy((*[1 << 30]byte)(unsafe.Pointer(uintptr(resultPtr)))[:len(resultBytes)], resultBytes)

	return resultPtr, uint32(len(resultBytes))
}

//export handleOrderAPI
func handleOrderAPI(inputPtr, inputLen uint32) (uint32, uint32) {
	inputBytes := make([]byte, inputLen)
	src := (*[1 << 30]byte)(unsafe.Pointer(uintptr(inputPtr)))[:inputLen:inputLen]
	copy(inputBytes, src)

	var input map[string]string
	json.Unmarshal(inputBytes, &input)

	path := input["path"]
	method := input["method"]
	body := input["body"]

	result := handleOrderAPIImpl(path, method, body)
	resultBytes := []byte(result)

	resultPtr := uint32(3072) // fixed offset
	copy((*[1 << 30]byte)(unsafe.Pointer(uintptr(resultPtr)))[:len(resultBytes)], resultBytes)

	return resultPtr, uint32(len(resultBytes))
}

func handleCatalogAPIImpl(path, method, body string) string {
	switch getEndpointName(path) {
	case "listProducts":
		products := getDemoProducts()
		result, _ := json.Marshal(map[string]interface{}{
			"products": products,
			"total":    len(products),
			"page":     1,
			"limit":    20,
			"hasMore":  false,
		})
		return string(result)
	case "getProduct":
		id := extractPathParam(path, "/products/{id}", "id")
		products := getDemoProducts()
		for _, p := range products {
			if p["id"] == id {
				result, _ := json.Marshal(p)
				return string(result)
			}
		}
		return `{"error": "product not found"}`
	case "searchProducts":
		query := getQueryParam(body, "q")
		products := getDemoProducts()
		filtered := []map[string]interface{}{}
		for _, p := range products {
			name := p["name"].(string)
			if strings.Contains(strings.ToLower(name), strings.ToLower(query)) {
				filtered = append(filtered, p)
			}
		}
		result, _ := json.Marshal(map[string]interface{}{
			"products": filtered,
			"total":    len(filtered),
		})
		return string(result)
	case "getCategories":
		result, _ := json.Marshal([]map[string]interface{}{
			{"id": "electronics", "name": "Electronics", "slug": "electronics"},
			{"id": "clothing", "name": "Clothing", "slug": "clothing"},
			{"id": "home", "name": "Home & Garden", "slug": "home"},
		})
		return string(result)
	default:
		return `{"message": "Catalog endpoint: ` + getEndpointName(path) + `"}`
	}
}

func handleCartAPIImpl(path, method, body string) string {
	switch getEndpointName(path) {
	case "getCart":
		userId := extractPathParam(path, "/cart/{userId}", "userId")
		result, _ := json.Marshal(map[string]interface{}{
			"userId": userId,
			"items":  []interface{}{},
			"total":  0,
		})
		return string(result)
	default:
		return `{"message": "Cart endpoint: ` + getEndpointName(path) + `"}`
	}
}

func handleOrderAPIImpl(path, method, body string) string {
	return `{"message": "Order endpoint: ` + getEndpointName(path) + `"}`
}

// Helper functions
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

func getEndpointName(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 4 {
		return parts[4]
	}
	return ""
}

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

func getQueryParam(body, key string) string {
	return ""
}

func main() {}
