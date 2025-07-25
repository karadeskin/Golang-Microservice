# Golang Microservice

Lightweight REST API written in Go with concurrency using goroutines and channels. Deployed locally using Docker and tested with Postman.

## Features

- Handles basic HTTP requests with the `net/http` package
- Demonstrates concurrency with goroutines and channels
- RESTful API design 
- Dockerized for easy deployment and testing

## Tech Stack

- **Go**: Core logic + goroutines
- **Docker**: Local containerized deployment
- **Postman**: API endpoint testing

## Project Structure

Golang-Microservice
- main.go
- Dockerfile
- go.mod
- README.md

## Getting Started

### Clone the repo
```bash
git clone https://github.com/your-username/Golang-Microservice.git
cd Golang-Microservice
```

### Run it locally 
```bash 
go run main.go
```

### Run with Docker
```bash
docker build -t golang-microservice .
docker run -p 8080:8080 golang-microservice
```

### Test the endpoint 
```bash
curl http://localhost:8080/
```

### or use postman to hit http://localhost:8080/

