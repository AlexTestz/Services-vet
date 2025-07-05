# 🐾 Services-Vet Domain

## 📘 Description
The `services-vet` domain is designed to manage pet daycare services, specifically for dogs. It includes microservices for registering, updating, deleting, and retrieving reservation records, all implemented under a microservices architecture and documented automatically using Swagger.

---

## 🧩 Included Microservices

- **register-reservation**: Registers a new dog daycare reservation.
- **update-reservation**: Modifies the details of an existing reservation.
- **delete-reservation**: Deletes an existing reservation.
- **get-reservation**: Retrieves reservation details (by ID).

---

## 🛠️ Technologies Used

- **Framework**: Go (Golang) for microservices development.
- **Database**: MongoDB for storing reservation data.
- **Swagger**: For automatic API documentation.
- **Docker**: For containerization and deployment of services.
- **Fiber**: Go web framework for building APIs.

---

## 🏗️ Architecture Style

The system follows a **microservices architecture**, where each microservice is independent and responsible for a specific functionality within the domain. Services communicate with each other via RESTful APIs, enabling scalability and independent maintenance.

---

## 🧠 Design Patterns Applied

- **Microservices**: Each functionality is encapsulated in an independent service.
- **Singleton**: Used for database configuration and services requiring a single instance.
- **Factory**: For building reservation services with different configurations.

---

## 🧱 Internal Architecture

The domain is composed of multiple microservices, each with its own business logic, database, and service configurations. Communication between services is asynchronous and based on HTTP using RESTful endpoints.

---

## 🔐 Security and Middleware

- **Authentication**: JWT-based token authentication for protecting routes that modify data.
- **Middleware**: Used for input validation, access control, and error handling.

---

## 📁 Domain Structure
/services-vet
│
├── create-service/
├── delete-service/
├── update-service/
├── get-service/
├── .gitignore
├── .gitattributes
└── README.md



---

## 🌐 Environment Variables

- **DB_HOST**: MongoDB database host address.  
- **DB_PORT**: MongoDB port.  
- **JWT_SECRET**: Secret key used to generate JWT tokens.  
- **SERVICE_PORT**: Port the microservice will listen on.

---

## 📡 Main Endpoints

- **POST** `/reservations` – Register a new reservation.  
- **PUT** `/reservations/{id}` – Update an existing reservation.  
- **DELETE** `/reservations/{id}` – Delete a reservation.  
- **GET** `/reservations/{id}` – Retrieve reservation details by ID.

---

## 📝 Notes

- Ensure that all required environment variables are correctly configured for database access and authentication.
- The API is automatically documented using Swagger and is accessible at the `/docs` endpoint.



