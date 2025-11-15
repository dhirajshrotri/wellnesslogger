
This project is a Golang API server that provides CRUD operations for tracking workout exercises. Data is stored in MongoDB using the official Go MongoDB driver.

## 📂 Project Structure
WellnessLogger/
│
├── main.go
├── config/
│ └── db.go
├── models/
│ └── exercise.go
├── handlers/
│ └── exercise_handler.go
└── routes/
└── exercise_routes.go


## Run the Server

### Prerequisites
- Go 1.20+
- MongoDB running locally (`mongodb://localhost:27017`)
- Curl or Postman for testing

### Install Dependencies
```
go mod tidy
```

## Start the API
```
go run main.go
```
Server will start at:

```
http://localhost:8080
```
## Exercise Model
```
{
  "userId": 1,
  "exercise": "Tricep Curls",
  "category": "ARMS",
  "sets": 3,
  "reps": 20,
  "date": "2025-11-09T00:00:00Z"
}
```

## API Endpoints & Example Requests

Create Exercise
```
POST /exercise
curl -X POST http://localhost:8080/exercise \
-H "Content-Type: application/json" \
-d '{
  "userId": 1,
  "exercise": "Tricep Curls",
  "category": "ARMS",
  "sets": 3,
  "reps": 20,
  "date": "2025-11-09T00:00:00Z"
}'
```

Get All Exercises

GET /exercise
```
curl http://localhost:8080/exercise
```

Update Exercise

PUT /exercise?id=<objectID>

```
curl -X PUT "http://localhost:8080/exercise?id=676fccd04993248d7d4b714e" \
-H "Content-Type: application/json" \
-d '{
  "exercise": "Push Ups",
  "category": "CHEST",
  "sets": 4,
  "reps": 15
}'
```

Delete Exercise

DELETE /exercise?id=<objectID>

```
curl -X DELETE "http://localhost:8080/exercise?id=676fccd04993248d7d4b714e"
```

