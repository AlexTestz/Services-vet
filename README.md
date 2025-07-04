# Services-Vet Domain

## Descripción
El dominio `services-vet` está diseñado para gestionar los servicios de una guardería de mascotas, específicamente para perros. Incluye microservicios para registrar, actualizar, eliminar y consultar reservas de estancia para las mascotas, todo implementado bajo una arquitectura de microservicios y respaldado por Swagger para la documentación automática.

## Microservicios incluidos
- **register-reservation**: Registrar una nueva reserva para la estancia de un perro.
- **update-reservation**: Modificar los detalles de una reserva existente.
- **delete-reservation**: Eliminar una reserva existente.
- **get-reservation**: Consultar detalles de una reserva (por ID).

## Tecnologías utilizadas
- **Framework**: Go (Golang) para la creación de microservicios.
- **Base de Datos**: MongoDB para el almacenamiento de datos de reservas.
- **Swagger**: Para la documentación automática de los microservicios.
- **Docker**: Para la contenedorización y despliegue de los microservicios.
- **Fiber**: Framework web de Go para la creación de APIs.

## Estilo de arquitectura
El sistema sigue una arquitectura de **microservicios**, en la que cada microservicio es independiente y está encargado de una funcionalidad específica dentro del dominio. Los microservicios se comunican entre sí a través de API RESTful, permitiendo escalabilidad y mantenimiento independiente.

## Patrones de diseño aplicados
- **Microservicios**: Cada funcionalidad está encapsulada en un microservicio independiente.
- **Singleton**: Se utiliza para la configuración de la base de datos y la inicialización de servicios que requieren una instancia única.
- **Factory**: Para la creación de servicios de reserva con diferentes configuraciones.

## Arquitectura interna
El dominio está compuesto por varios microservicios, cada uno con su propia lógica de negocio, bases de datos y configuraciones de servicio. La comunicación entre los servicios es asincrónica y basada en HTTP, utilizando endpoints RESTful.

## Seguridad y Middleware
- **Autenticación**: Implementación de autenticación basada en tokens JWT para asegurar las rutas que modifican datos.
- **Middleware**: Se utilizan middlewares para manejar la validación de entradas, control de acceso y gestión de errores.

## Estructura del dominio
/services-vet
│
├── create-service/
├── delete-service/
├── update-service/
├── get-service/
├── .gitignore
├── .gitattributes
└── README.md


## Variables de entorno
- **DB_HOST**: Dirección de la base de datos MongoDB.
- **DB_PORT**: Puerto de la base de datos.
- **JWT_SECRET**: Clave secreta para la creación de tokens JWT.
- **SERVICE_PORT**: Puerto en el que el microservicio estará corriendo.

## Endpoints principales
- **POST /reservations**: Registrar una nueva reserva.
- **PUT /reservations/{id}**: Actualizar una reserva existente.
- **DELETE /reservations/{id}**: Eliminar una reserva.
- **GET /reservations/{id}**: Consultar una reserva por ID.

## Notas
- Asegúrese de tener configuradas correctamente las variables de entorno para el acceso a la base de datos y la autenticación.
- La API está documentada automáticamente a través de Swagger, accesible a través de la ruta `/docs`.

