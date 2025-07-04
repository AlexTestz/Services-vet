# Update Service Microservice

## Descripción
Microservicio encargado de actualizar servicios veterinarios registrados en el sistema, permitiendo actualizar todos los servicios.

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
   docker build -t alexmpz/update-service:qa 
   ```
2. Ejecuta el contenedor:
   ```sh
   docker run -p 3016:3016 --env-file .env alexmpz/update-service:qa
   ```

---

## Endpoints principales

### Obtener todos los servicios

- **GET** `/api/services/`

### Buscar servicio por id 

- **GET** `/api/services?name=<nombre_del_servicio>`

### Buscar servicio por nombre (parámetro de ruta, menos recomendado)

- **GET** `/api/services/name/:name`

---

## Ejemplo de request

**Obtener todos los servicios:**
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

## Ejemplo de respuesta exitosa

**Respuesta para todos los servicios:**
```json
{
    "message": "Service updated ✅"
}

## Notas

- El endpoint con query param (`/api/services?name=...`) es más robusto para búsquedas con tildes, ñ o espacios.
- El endpoint por ruta (`/api/services/name/:name`) puede fallar con caracteres especiales debido a la codificación de la URL.
- El CORS está abierto para facilitar el desarrollo, pero se recomienda restringirlo en producción.
- La conexión a MongoDB debe estar correctamente configurado desde el