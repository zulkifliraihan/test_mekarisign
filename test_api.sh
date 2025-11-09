#!/bin/bash

# Collaborative Todo List API - Test Script
# Make sure the server is running before executing this script

BASE_URL="http://localhost:8080"

echo "====================================="
echo "Collaborative Todo List API - Tests"
echo "====================================="
echo ""

echo "1. Health Check"
echo "GET $BASE_URL/health"
curl -s $BASE_URL/health | jq .
echo -e "\n"

echo "2. Get API Info"
echo "GET $BASE_URL/"
curl -s $BASE_URL/ | jq .
echo -e "\n"

echo "3. Get All Users"
echo "GET $BASE_URL/users"
curl -s $BASE_URL/users | jq .
echo -e "\n"

echo "4. Create Todo for John Doe (User 1)"
echo "POST $BASE_URL/todos"
TODO1=$(curl -s -X POST $BASE_URL/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Review code", "user_id": 1, "completed": false}')
echo $TODO1 | jq .
TODO1_ID=$(echo $TODO1 | jq -r '.id')
echo -e "\n"

echo "5. Create Todo for Jane Smith (User 2)"
echo "POST $BASE_URL/todos"
TODO2=$(curl -s -X POST $BASE_URL/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Write documentation", "user_id": 2, "completed": false}')
echo $TODO2 | jq .
TODO2_ID=$(echo $TODO2 | jq -r '.id')
echo -e "\n"

echo "6. Create Another Todo for John Doe"
echo "POST $BASE_URL/todos"
TODO3=$(curl -s -X POST $BASE_URL/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Fix bugs", "user_id": 1, "completed": false}')
echo $TODO3 | jq .
echo -e "\n"

echo "7. Get All Todos"
echo "GET $BASE_URL/todos"
curl -s $BASE_URL/todos | jq .
echo -e "\n"

echo "8. Get Todos for User 1 (John Doe)"
echo "GET $BASE_URL/todos?user_id=1"
curl -s "$BASE_URL/todos?user_id=1" | jq .
echo -e "\n"

echo "9. Get Todos for User 2 (Jane Smith)"
echo "GET $BASE_URL/todos?user_id=2"
curl -s "$BASE_URL/todos?user_id=2" | jq .
echo -e "\n"

echo "10. Toggle Todo Completion (ID: $TODO1_ID)"
echo "PATCH $BASE_URL/todos/$TODO1_ID/toggle"
curl -s -X PATCH $BASE_URL/todos/$TODO1_ID/toggle | jq .
echo -e "\n"

echo "11. Update Todo (ID: $TODO2_ID)"
echo "PUT $BASE_URL/todos/$TODO2_ID"
curl -s -X PUT $BASE_URL/todos/$TODO2_ID \
  -H "Content-Type: application/json" \
  -d '{"text": "Write documentation and examples", "user_id": 2, "completed": true}' | jq .
echo -e "\n"

echo "12. Get All Todos (after updates)"
echo "GET $BASE_URL/todos"
curl -s $BASE_URL/todos | jq .
echo -e "\n"

echo "13. Delete Todo (ID: $TODO1_ID)"
echo "DELETE $BASE_URL/todos/$TODO1_ID"
curl -s -X DELETE $BASE_URL/todos/$TODO1_ID | jq .
echo -e "\n"

echo "14. Get All Todos (after deletion)"
echo "GET $BASE_URL/todos"
curl -s $BASE_URL/todos | jq .
echo -e "\n"

echo "15. Test Error - User Not Found (user_id: 999)"
echo "POST $BASE_URL/todos"
curl -s -X POST $BASE_URL/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "Test", "user_id": 999, "completed": false}' | jq .
echo -e "\n"

echo "16. Test Error - Empty Text"
echo "POST $BASE_URL/todos"
curl -s -X POST $BASE_URL/todos \
  -H "Content-Type: application/json" \
  -d '{"text": "", "user_id": 1, "completed": false}' | jq .
echo -e "\n"

echo "====================================="
echo "All tests completed!"
echo "====================================="
