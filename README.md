# epimarket
ecommerce application made for https://priostack.com
A complete e-commerce example with multiple microservices.

**Services:**
- `catalog/` - Product catalog service
- `cart/` - Shopping cart service
- `order/` - Order management service
- `inventory/` - Inventory tracking service
- `payment/` - Payment processing service
- `shipping/` - Shipping service
- `notification/` - Notification service

## Usage

Each example can be deployed as a bundle:

```bash
# Import example bundle
curl -X POST -F "zip=@ecommerce/ecommerce.zip" http://localhost:8080/api/bundles/import/zip

# Start the bundle
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"ecommerce","workers":5}' \
  http://localhost:8080/api/bundles/run
```

## License

Apache 2.0
