# 🗑️ Delete Service Microservice

## Description
This microservice is responsible for deleting veterinary services from the system, allowing efficient management by ID or name.

---

## 🧪 Technologies Used

- **Language:** Go 1.23.3  
- **Framework:** Fiber v2  
- **Database:** MongoDB

---

## ⚙️ Architecture Style

- **RESTful API:** All operations are exposed via HTTP endpoints following REST principles.

---

## 🧠 Design Patterns Applied

- **KISS (Keep It Simple, Stupid):** Clean, straightforward code avoiding unnecessary complexity.  
- **DRY (Don't Repeat Yourself):** Logic and functions are reused to avoid duplication.  
- **Separation of Concerns:** Routes, controllers, and data access logic are clearly separated.

---

## 🗄️ Database

- **MongoDB:** Connected via the official Go `mongo-driver`. The URI and database name are configured using environment variables.

---

## 🧱 Internal Architecture

- **Layered Architecture:** Code is separated into `routes`, `controllers`, `models`, and `database` configuration.  
- **MVC-like Pattern:** Follows MVC principles even without an ORM.

---

## 🔐 Security & Middleware

- **CORS:** Configured to accept requests from any origin (useful for development).  
- **Validation:** Input validation is handled within the controllers.  
- **Logging:** Logs important events such as database connection.

---

## 📁 Project Structure



```
delete-service/
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
  docker build -t alexmpz/delete-service:qa .
  ```
2. Run the container:
```sh
  docker run -p 3018:3018 --env-file .env alexmpz/delete-service:qa
  ```

---

## Main endpoints

### Delete a service by ID

- **DELETE** `/api/services/:id`

### Delete a service by name

- **DELETE** `/api/services/name/:name`

---

### Request example

**By ID:**
```
DELETE http://localhost:3018/api/services/60c72b2f9b1e8b001c8e4b8a
```

**By name:**
```
DELETE http://localhost:3018/api/services/name/Consulta%20veterinaria
```

port 3017
---

### Ejemplo de respuesta exitosa

```json
{
  "message": "✅ Service deleted successfully"
}
```

---

## Notes

- The endpoint by name deletes only the first service that exactly matches the name provided.
- CORS is open to facilitate development, but it is recommended to restrict it in production.
- The connection to MongoDB must be correctly configured and accessible from the container.
- It is recommended to validate that the service name is unique if you are going to use the delete endpoint by name in