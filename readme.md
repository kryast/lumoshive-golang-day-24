<!-- create order -->
curl -X POST http://localhost:8080/create-order \
  -H "Content-Type: application/json" \
  -d '{
    "customer_name": "Ahmad 3",
    "order_number": "P0003",
    "order_status": "Completed",
    "order_date": "2024-11-07"
  }'


<!-- create admin -->
curl -X POST http://localhost:8080/create-admin -d '{
    "name": "Ahmad 1",
    "username": "user1",
    "password": "pass"
}' -H "Content-Type: application/json"
