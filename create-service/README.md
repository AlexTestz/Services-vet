# Create Service Microservice

## Descripción
Microservicio encargado de registrar servicios veterinarios en el sistema, permitiendo la gestión y consulta de los mismos.

---

## Tecnologías utilizadas

- **Lenguaje:** Go 1.23.3
- **Framework:** Fiber v2
- **Base de datos:** MongoDB

---

## Estilo de arquitectura

- **API RESTful:** Todas las operaciones se exponen como endpoints HTTP siguiendo el estilo REST.

---

## Patrones de diseño aplicados

- **KISS (Keep It Simple, Stupid):** Código sencillo y directo, evitando complejidad innecesaria.
- **DRY (Don't Repeat Yourself):** Reutilización de funciones y lógica para evitar duplicidad.
- **Separación de responsabilidades:** Rutas, controladores y lógica de acceso a datos están claramente diferenciados.

---

## Base de datos

- **MongoDB:** Conexión mediante el driver oficial de Go (`mongo-driver`). La URI y el nombre de la base de datos se configuran por variables de entorno.

---

## Arquitectura interna

- **N-capas:** Separación entre rutas (`routes`), controladores (`controllers`), modelos (`models`) y configuración de base de datos (`database`).
- **Modelo similar a MVC:** Aunque no se usa un ORM, la estructura sigue la separación de responsabilidades típica de MVC.

---

## Seguridad y Middleware

- **CORS:** Configurado para aceptar peticiones desde cualquier origen, útil para desarrollo y pruebas.
- **Validaciones:** Validación de datos de entrada en los controladores.
- **Logs:** Registro de eventos importantes como la conexión a la base de datos.

---

## Estructura del proyecto

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

## Variables de entorno

El microservicio utiliza variables de entorno para la configuración de la base de datos y el puerto de escucha. Estas se definen en el archivo `.env`.



## Ejecución local

1. Instala Go 1.23.3 o superior.
2. Instala las dependencias:
   ```sh
   go mod download
   ```
3. Ejecuta el microservicio:
   ```sh
   go run cmd/main.go
   ```

---

## Docker local

1. Construye la imagen:
   ```sh
   docker build -t alexmpz/create-service:qa .
   ```
2. Ejecuta el contenedor:
   ```sh
   docker run -p 3015:3015 --env-file .env alexmpz/create-service:qa
   ```

PORT 3018
---

## Endpoints principales

### Registrar un servicio

- **POST** `/api/services/`

#### Ejemplo de request

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

## Notas

- Antes de registrar un servicio, se valida que los campos requeridos estén presentes y sean válidos.
- El microservicio está preparado para ser desplegado en entornos Docker y Kubernetes.
- El CORS está abierto para facilitar el desarrollo, pero se recomienda restringirlo en producción.
- La conexión a MongoDB debe estar correctamente configurada y accesible desde el contenedor.