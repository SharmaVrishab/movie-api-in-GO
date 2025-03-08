# Movie API in Go (Golang)

## Overview
This is a simple RESTful API built using **Golang** and the **Gorilla Mux** router. The API allows users to perform CRUD (Create, Read, Update, Delete) operations on movies. Each movie has an `ID`, `ISBN`, `Title`, and a `Director` (which contains `Firstname` and `Lastname`).

## Features
- Get all movies (`GET /movies`)
- Get a specific movie by ID (`GET /movie/{id}`)
- Create a new movie (`POST /movies`)
- Update an existing movie (`PUT /movie/{id}`)
- Delete a movie (`DELETE /movie/{id}`)

## Technologies Used
- **Go (Golang)**
- **Gorilla Mux** (for routing)
- **Encoding/Decoding JSON**
- **Net/HTTP** (for handling requests and responses)

## Installation & Setup

### Prerequisites
- **Go** (Download and install [Go](https://go.dev/dl/))

### Clone the Repository
```sh
git clone https://github.com/your-username/movie-api-go.git
cd movie-api-go
```

### Install Dependencies
```sh
go mod tidy
```

### Run the Server
```sh
go run main.go
```
The server will start at: `http://localhost:8080`

## API Endpoints

### 1. Get All Movies
**Endpoint:** `GET /movies`
```sh
curl -X GET http://localhost:8080/movies
```
**Response:**
```json
[
  {
    "id": "1",
    "isbn": "10",
    "title": "MOVIE ONE",
    "director": {
      "firstname": "best",
      "lastname": "director"
    }
  }
]
```

### 2. Get a Movie by ID
**Endpoint:** `GET /movie/{id}`
```sh
curl -X GET http://localhost:8080/movie/1
```
**Response:**
```json
{
  "id": "1",
  "isbn": "10",
  "title": "MOVIE ONE",
  "director": {
    "firstname": "best",
    "lastname": "director"
  }
}
```

### 3. Create a New Movie
**Endpoint:** `POST /movies`
```sh
curl -X POST http://localhost:8080/movies \
  -H "Content-Type: application/json" \
  -d '{"isbn":"20","title":"NEW MOVIE","director":{"firstname":"John","lastname":"Doe"}}'
```
**Response:**
```json
{
  "id": "87432912",
  "isbn": "20",
  "title": "NEW MOVIE",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

### 4. Update an Existing Movie
**Endpoint:** `PUT /movie/{id}`
```sh
curl -X PUT http://localhost:8080/movie/1 \
  -H "Content-Type: application/json" \
  -d '{"isbn":"30","title":"UPDATED MOVIE","director":{"firstname":"Jane","lastname":"Smith"}}'
```
**Response:**
```json
{
  "id": "1",
  "isbn": "30",
  "title": "UPDATED MOVIE",
  "director": {
    "firstname": "Jane",
    "lastname": "Smith"
  }
}
```

### 5. Delete a Movie
**Endpoint:** `DELETE /movie/{id}`
```sh
curl -X DELETE http://localhost:8080/movie/1
```
**Response:**
```json
[
  {
    "id": "2",
    "isbn": "11",
    "title": "MOVIE TWO",
    "director": {
      "firstname": "best",
      "lastname": "director"
    }
  }
]
```

## Project Structure
```
movie-api-go/
│── main.go        # Main server logic
│── go.mod         # Go module file
│── go.sum         # Dependencies
```

## Explanation of Key Parts

### 1. Movie Struct with JSON Tags
```go
type Movie struct {
    ID       string    `json:"id"`
    ISBN     string    `json:"isbn"`
    Title    string    `json:"title"`
    Director *Director `json:"director"`
}
```
**Why use JSON tags?**
- When encoding/decoding JSON, the field names will match the specified JSON keys instead of struct field names.

### 2. Handling Request Parameters with Mux
```go
params := mux.Vars(r) // Gets URL parameters as a map
```
Example: If request is `GET /movie/1`, then `params["id"] == "1"`

### 3. Slicing for Deleting Movies
```go
movies = append(movies[:index], movies[index+1:]...)
```
- This removes the movie at `index` from the `movies` slice.
- **`...` (spread operator)** is used to unpack the remaining slice elements.

## Contributing
1. Fork the repository
2. Create a new feature branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m "Added new feature"`)
4. Push the branch (`git push origin feature-branch`)
5. Open a Pull Request

---

🚀 **Happy Coding!**

