# 🧠 Go Fiber User CRUD API

A clean architecture-based REST API built with **Go Fiber**, **PostgreSQL**, **SQLC**, and **Uber Zap** logging.

---

##  Features
- Create, Read, Update, Delete users
- Input validation using go-playground/validator
- Structured logging with Uber Zap
- Global request logging middleware
- PostgreSQL + SQLC for database layer

---

##  Tech Stack
- **Language:** Go
- **Framework:** Fiber
- **Database:** PostgreSQL
- **ORM/Queries:** SQLC
- **Logging:** Uber Zap
- **Validation:** go-playground/validator

---

### Setup
```bash
1) git clone https://github.com/salwinsolomon/go-user-crud-api.git
    cd task
    go mod tidy

2) Create the database 
    go to neon website
    sigup for a free account 
    copy paste the link for the "connect to your database"

3) Create a .env file in the root of the project 
    paste DATABASE_URL=<paste-your-neon-database-key-here>

go run ./cmd/server

To test the API
 you can go to postman 

paste your base url :http://localhost:8080
Endpoints
| Method | Endpoint     | Description         |
| ------ | ------------ | ------------------- |
| POST   | `/users`     | Create a new user   |
| GET    | `/users/:id` | Get a user by ID    |
| GET    | `/users`     | List all users      |
| PUT    | `/users/:id` | Update a user       |
| DELETE | `/users/:id` | Delete a user       |
| GET    | `/health`    | Check server status |



