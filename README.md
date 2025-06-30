[![Go Version](https://img.shields.io/badge/go-1.21-blue?logo=go&logoColor=white)](https://golang.org/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)


# 🧑‍🎓 Students-api

A lightweight and modular RESTful API in Go to manage student records using SQLite. Built with clean architecture, structured logging, and input validation.

## 📌 Features

- ➕ Create new students (`POST`)
- 📄 Get all students (`GET`)
- ✏️ Update student info (`PUT`)
- ❌ Delete a student (`DELETE`)
- ✅ Input validation using `go-playground/validator`
- 📦 Persistent storage using SQLite
- ⚙️ Configurable via YAML
- 🧱 Clean folder structure with interfaces


## 🛠️ Tech Stack

- **Go 1.21+**
- **SQLite** via `mattn/go-sqlite3`
- **HTTP Server** using `net/http`
- **Structured Logging** using `log/slog`
- **Input Validation** using `go-playground/validator`
- **YAML Configuration**

## 🗂 Folder Structure
```
students-api-go/
├── cmd/
│ └── students-api/ # Entry point
│ └── main.go
├── config/
│ └── local.yaml # App configuration
├── internal/
│ ├── config/ # Config loader
│ ├── http/
│ │ └── handlers/
│ │ └── student/ # API handler
│ ├── storage/
│ │ ├── sqlite/ # SQLite driver
│ │ └── storage.go # Interface (future-proofing)
│ ├── types/ # Domain types
│ └── utils/response/ # Standard response writer
├── storage/
│ └── storage.db # SQLite DB file (auto-created)
├── go.mod
└── go.sum
```

## ⚙️ Configuration

`config/local.yaml`:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server: 
  address: "localhost:8082"
```
## 🚀 How to Run

## 1. Clone the repo
```
git clone https://github.com/Ashank007/students-api-go.git
cd students-api-go
```
## 2. Run the app
```
CONFIG_PATH=config/local.yaml go run cmd/students-api/main.go
```

ℹ️ The server will start at http://localhost:8082
The SQLite database is auto-created at storage/storage.db


## 📡 API Reference

🔹 POST /api/students

Create a new student record.

🔸 Request Body
```
{
  "name": "John",
  "email": "john@example.com",
  "age": 21
}
```
🔸 Response
```
{
  "id": 1
}
```

🔹 GET /api/students

Fetch all students

🔸 Response
```
[
  {
    "id": 1,
    "name": "John",
    "email": "John@example.com",
    "age": 21
  }
]
```


## 🧪 Sample curl Commands

Create a new student
```
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"John@example.com","age":21}'
```
Get all students
```
curl http://localhost:8082/api/students
```
Update student
```
curl -X PUT http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"id":1,"name":"Updated Name","email":"updated@mail.com","age":23}'
```
Delete student
```
curl -X DELETE "http://localhost:8082/api/students?id=1"
```
## 📌 Validation Rules
```
Field	Rule
name	required
email	required
age	    required
```

## 🧠 Future Enhancements

- Add support for pagination

- Add GET /api/students/{id} for single record

- Swagger/OpenAPI auto docs

- Dockerfile for containerization

- CI pipeline via GitHub Actions

## 🪪 License

This project is licensed under the MIT License.