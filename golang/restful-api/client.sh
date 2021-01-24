API='http://127.0.0.1:9001/api/v1'
TYPE_TEXT='text/plain'
TYPE_JSON='Content-Type: application/json'
TYPE_FORM="application/x-www-form-urlencoded"
TOKEN='ABCDEF0123456789'
HEADER_TOKEN="Authorization:$TOKEN"

echo 'Test GET /users'
curl -X GET -H $HEADER_TOKEN -H $TYPE_JSON $API/users

echo 'Test POST /users/1'
curl -X POST -H $HEADER_TOKEN -H $TYPE_JSON -d '{"name":"FF", "age":1}' $API/users

echo 'Test PUT /users/1'
curl -X PUT -H $HEADER_TOKEN -H $TYPE_JSON -d '{"age":"8"}' $API/users/1

echo 'Test DELETE /users/2'
curl -X DELETE -H $HEADER_TOKEN -H $TYPE_JSON $API/users/2
curl -X GET -H $HEADER_TOKEN -H $TYPE_JSON $API/users

echo 'Test completed.'