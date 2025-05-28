# TaskManager
Run the command: `go run main.go` from the root directory(inside task-manager)

## Sample Requests:
### Create tasks:
```curl --location 'http://localhost:8080/tasks' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Backup Job 8",
  "status": "completed",
  "trigger_time": "2025-06-01T08:00:00Z"
}'```


### Update task:
```curl --location --request PUT 'http://localhost:8080/tasks/790aa9d6-4a4b-4be0-8c04-bb772ffeb43d' \
--header 'Content-Type: application/json' \
--data '{
  "name": "New Name Job",
  "status": "active",
  "trigger_time": "2025-06-01T08:00:00Z"
}'```

### Get task:
```curl --location --request GET 'http://localhost:8080/tasks/790aa9d6-4a4b-4be0-8c04-bb772ffeb43d' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Backup Job",
  "status": "active",
  "trigger_time": "2025-06-01T08:00:00Z"
}'```


### Delete task:
`curl --location --request DELETE 'http://localhost:8080/tasks/8a297319-c3fe-4650-a57a-5acdf4da556b'`


### List tasks:
```curl --location --request GET 'http://localhost:8080/tasks?page=3&limit=1&status=completed' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Backup Job",
  "status": "active",
  "trigger_time": "2025-06-01T08:00:00Z"
}'
'```

