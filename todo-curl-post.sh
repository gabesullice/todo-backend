curl -X POST -H "Content-Type: application/json" http://localhost:8080/todos --data @- <<EOF
{
  "data": {
    "type": "todos",
      "attributes": {
        "title": "Create database",
        "body": "This will help the students use $http to practice GET and POST",
        "done": false
      }
  }
}
EOF
