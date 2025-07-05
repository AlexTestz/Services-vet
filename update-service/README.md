# 🔄 Update Service Microservice

## Description
Microservice responsible for updating registered veterinary services in the system, allowing the modification of existing services.

---

## 🧪 Technologies Used

- **Language:** Go 1.23.3  
- **Framework:** Fiber v2  
- **Database:** MongoDB

---

## ⚙️ Architecture Style

- **RESTful API:** All operations are exposed as HTTP endpoints following REST principles.

---

## 🧠 Design Patterns Applied

- **KISS (Keep It Simple, Stupid):** Simple and straightforward code avoiding unnecessary complexity.  
- **DRY (Don't Repeat Yourself):** Reuse of logic and functions to avoid duplication.  
- **Separation of Concerns:** Routes, controllers, and data access logic are clearly separated.

---

## 🗄️ Database

- **MongoDB:** Connected via Go’s official driver (`mongo-driver`). URI and database name are configured via environment variables.

---

## 🧱 Internal Architecture

- **Layered Architecture:** Separation between `routes`, `controllers`, `models`, and database config.  
- **MVC-like Model:** Even without using an ORM, the structure follows typical MVC separation of concerns.

---

## 🔐 Security and Middleware

- **CORS:** Configured to accept requests from any origin, useful for development and testing.  
- **Validation:** Input data is validated in the controllers.  
- **Logging:** Logs important events such as MongoDB connection.

---

## 📁 Project Structure



```
update-service/
│
├── cmd/
│   └── main.go
├── controllers/
│   └── service_controller.go
├── routes/
│   └── service_routes.go
├── models/
│   └── service.go
├── database/
│   └── mongo.go
├── config/
│   └── config.go
├── go.mod
├── go.sum
├── Dockerfile
├── .env
└── README.md
```

---

## Environment variables

The microservice uses environment variables for database configuration and listening port. These are defined in the `.env` file.



## Local execution

1. Install Go 1.23.3 or higher.
2. Install dependencies:
```sh
   go mod download
   ```
3. Run the microservice:
```sh
   go run cmd/main.go
   ```

---

## Local Docker

1. Build the image:
```sh
   docker build -t alexmpz/update-service:qa 
   ```
2. Run the container:
```sh
   docker run -p 3016:3016 --env-file .env alexmpz/update-service:qa
   ```

port 3015
---

## Main endpoints

### Get all services

- **GET** `/api/services/`

### Search for service by ID 

- **GET** `/api/services?name=<service_name>`

### Search for service by name (route parameter, less recommended)

- **GET** `/api/services/name/:name`

---

## Request example

**Get all services:**
```
GET http://localhost:3017/api/services/name/corte
```

{
  "name": " cortee",
  "description": "Servicio completo de  corte de pelo con fragancia",
  "price": 30,
  "duration": 60
}
```

---

## Example of a successful response

**Response for all services:**
```json
{
    “message”: “Service updated ✅”
}

## Notes

- The endpoint with query param (`/api/services?name=...`) is more robust for searches with accents, ñ, or spaces.
- The endpoint by route (`/api/services/name/:name`) may fail with special characters due to URL encoding.
- CORS is open to facilitate development, but it is recommended to restrict it in production.
- The connection to MongoDB must be correctly configured from the