# Fiber CRM Basic

A simple Customer Relationship Management (CRM) REST API built with Go. This is my beginner project created to learn Go programming and web API development.

## Project Overview

I built this project to gain practical experience with Go and to understand how to create REST APIs. The application allows for basic lead management with CRUD operations (Create, Read, Update, Delete).

## Technologies Used

- **Go (Golang)** - The core programming language
- **Fiber v2** - A fast, Express-inspired web framework
- **GORM** - ORM library for Go
- **SQLite** - Lightweight disk-based database
- **Postman** - For testing API endpoints

## Project Structure and Design

The project follows a modular design pattern:

```
go_fiber_crm_basic/
├── database/
│   └── database.go  # Database connection configuration
├── lead/
│   └── lead.go      # Lead model and handlers
├── go.sum           # intialized file
├── go.mod           # mod file for all the dependencies
├── main.exe         # compiled file
├── main.go          # Entry point and route setup
├── test.db          # SQLite database
└── README.md
```

I designed the application with a clean separation of concerns:
- **Models**: Define the data structure
- **Handlers**: Process requests and return responses
- **Database**: Manages connections and operations

## Features

- Create new leads with contact information
- Retrieve individual leads or all leads
- Delete leads
- Database auto-migration for easy setup

## What I Learned

Through this project, I gained valuable experience with:

1. **Go Programming**:
   - Structuring a Go application
   - Working with packages and imports
   - Error handling patterns
   - Using pointers and structs

2. **REST API Development**:
   - Implementing different HTTP methods (GET, POST, DELETE)
   - Request/response handling
   - Status codes and error responses

3. **Database Operations**:
   - Using GORM for ORM capabilities
   - Connecting to SQLite database
   - Performing CRUD operations
   - Automatic migrations

4. **Testing with Postman**:
   - Creating API requests
   - Validating responses
   - Testing different endpoints

## Setup and Running

1. Clone the repository
2. Make sure Go is installed
3. Install dependencies:
   ```
   go mod download
   ```
4. Run the application:
   ```
   go run main.go
   ```
5. The server will start on port 3000

## API Endpoints

- `GET /api/v1/lead` - Get all leads
- `GET /api/v1/lead/:id` - Get a specific lead
- `POST /api/v1/lead` - Create a new lead
- `DELETE /api/v1/lead/:id` - Delete a lead

## Future Improvements

As I continue learning, I plan to add:
- Authentication and authorization
- More complex data relationships
- Frontend UI
- Deployment to cloud services

---

This project, while simple, provided me with a solid foundation in Go web development and API design. It helped me understand the Go ecosystem and how to build efficient backend services.

---
