# 🛠️ Create Service Microservice

## Description
This microservice handles the registration of veterinary services in the system, enabling their management and retrieval.

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

- **KISS (Keep It Simple, Stupid):** Clean and straightforward code with no unnecessary complexity.  
- **DRY (Don't Repeat Yourself):** Functions and logic are reused to avoid redundancy.  
- **Separation of Concerns:** Routes, controllers, and data access logic are clearly separated.

---

## 🗄️ Database

- **MongoDB:** Connected using the official Go `mongo-driver`. The URI and database name are configured through environment variables.

---

## 🧱 Internal Architecture

- **Layered Architecture:** Clear separation into `routes`, `controllers`, `models`, and `database` configuration.  
- **MVC-like Structure:** While no ORM is used, the project follows a typical MVC-like separation of responsibilities.

---

## 🔐 Security & Middleware

- **CORS:** Configured to accept requests from all origins (useful for development).  
- **Validation:** Input data is validated in the controller layer.  
- **Logging:** Important events like database connection are logged.

---

## 📁 Project Structure


```
create-service/
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


---

## 🌐 Environment Variables

The microservice uses environment variables to configure the database connection and the listening port. These are defined in the `.env` file.

---

## 🚀 Local Execution

1. Install Go 1.23.3 or higher.  
2. Download dependencies:
   ```bash
   go mod download

   ```
3. Execute microservice:
   ```sh
   go run cmd/main.go
   ```

---

## Docker local

1. build docker image:
   ```sh
   docker build -t alexmpz/create-service:qa .
   ```
2. run container:
   ```sh
   docker run -p 3015:3015 --env-file .env alexmpz/create-service:qa
   ```

PORT 3018
---

## Endpoints

### Register a service

- **POST** `/api/services/`

#### Request example

```json
{
  “name”: “Veterinary consultation”,
  “description”: “General consultation for pets”,
  “price”: 100,
  “duration”: 30
}
```

#### Example of a successful response

```json
{
  "name": "Consulta veterinaria",
  "description": "Consulta general para mascotas",
  "price": 100,
  "duration": 30
}
```

#### Ejemplo de respuesta exitosa

```json
{
  "message": "Service created successfully ✅",
  "service": {
    "id": "60c72b2f9b1e8b001c8e4b8a",
    "name": "Consulta veterinaria",
    "description": "Consulta general para mascotas",
    "price": 100,
    "duration": 30
  }
}
```

---

## Notes

- Before registering a service, the required fields are validated to ensure they are present and valid.
- The microservice is ready to be deployed in Docker and Kubernetes environments.
- CORS is open to facilitate development, but it is recommended to restrict it in production.
- The connection to MongoDB must be correctly configured and accessible from the container.