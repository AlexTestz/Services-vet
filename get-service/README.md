# 📄 Get Service Microservice

## Description
Microservice responsible for retrieving veterinary services registered in the system, allowing you to get all services or search by name.

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

- **KISS (Keep It Simple, Stupid):** Clean and straightforward code avoiding unnecessary complexity.  
- **DRY (Don't Repeat Yourself):** Logic and functions are reused to avoid duplication.  
- **Separation of Concerns:** Routes, controllers, and data access logic are clearly separated.

---

## 🗄️ Database

- **MongoDB:** Connected using the official Go driver (`mongo-driver`). URI and database name are configured through environment variables.

---

## 🧱 Internal Architecture

- **Layered Architecture:** Clear separation between `routes`, `controllers`, `models`, and database configuration.  
- **MVC-like Structure:** While not using an ORM, the structure follows typical MVC responsibilities.

---

## 🔐 Security and Middleware

- **CORS:** Configured to accept requests from any origin, useful for development and testing.  
- **Validation:** Input validation is handled inside the controllers.  
- **Logs:** Logs important events such as database connection status.

---

## 📁 Project Structure



```
get-service/
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
2. Install the dependencies:
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
  docker build -t alexmpz/get-service:qa .
  ```
2. Run the container:
```sh
  docker run -p 3016:3016 --env-file .env alexmpz/get-service:qa
  ```

---

## Main endpoints

### Get all services

- **GET** `/api/services/`

### Search for service by name (query param)

- **GET** `/api/services?name=<service_name>`

### Search for service by name (route parameter, less recommended)

- **GET** `/api/services/name/:name`

---

## Request example

**Get all services:**
```
GET http://localhost:3016/api/services/
```


```

---
## EXAMPLE

**Example response all services:**
```json
[
  {
    "id": "6857768c1c50b7293e2fd6cc",
    "name": "Baño y corte",
    "description": "Servicio completo de baño y corte de pelo",
    "price": 25.5,
    "duration": 45
  },
  {
    "id": "6857769c1c50b7293e2fd6cd",
    "name": "Baño",
    "description": "Servicio completo de baño y corte de pelo",
    "price": 25.5,
    "duration": 45
  }
]
```



## Notes

- The endpoint with query param (`/api/services?name=...`) is more robust for searches with accents, ñ, or spaces.
- The endpoint by route (`/api/services/name/:name`) may fail with special characters due to URL encoding.
- CORS is open to facilitate development, but it is recommended to restrict it in production.
- The connection to MongoDB must be correctly configured and accessible from the