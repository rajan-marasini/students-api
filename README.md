# Students API

A RESTful CRUD API for managing student records, built using Go and structured with best practices for scalability and maintainability.

## 🚀 Features

-   Create, Read, Update, and Delete (CRUD) operations for student data
-   Modular project structure with `cmd/`, `config/`, and `internal/` directories
-   Configuration management for flexible environment setups
-   Clean and idiomatic Go codebase

## 📁 Project Structure

```
students-api/
├── cmd/                # Application entry point
├── config/             # Configuration files
├── internal/           # Internal application logic
├── go.mod              # Go module file
├── go.sum              # Go dependencies
└── .gitignore          # Git ignore file
```

## 🛠️ Getting Started

### Prerequisites

-   Go 1.18 or higher

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/rajan-marasini/students-api.git
    cd students-api
    ```

2\. Install dependencies:

```bash
go mod tidy
```

3\. Run the application:

```bash
go run cmd/students-api/main.go -config config/local.yml
```

## 📦 API Endpoints

| Method | Endpoint         | Description                |     |
| ------ | ---------------- | -------------------------- | --- |
| GET    | `/students`      | Retrieve all students      |     |
| GET    | `/student/{id}`  | Retrieve a student by ID   |     |
| POST   | `/student`       | Create a new student       |     |
| PUT    | `/students/{id}` | Update an existing student |     |
| DELETE | `/student/{id}`  | Delete a student           |
