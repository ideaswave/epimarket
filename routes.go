package ecommerce

import (
	"fmt"

	"github.com/devrevolt/qubit-gateway"
	"github.com/devrevolt/qubit-memory"
)

// StartEcommerceRoutes starts e-commerce-specific routes
func StartEcommerceRoutes(routes []gateway.Route, memoryManager *memory.MemoryManager) {
	for _, route := range routes {
		fmt.Printf("[Qubit] Starting %s route: %s\n", route.Type, route.Name)
		switch route.Type {
		case "http":
			gateway.LoadSwagger(route, memoryManager)
		case "grpc":
			// gRPC services like payment and inventory
			fmt.Printf("[Qubit] gRPC service ready: %s at %s\n", route.Name, route.Host)
		}
	}
}
