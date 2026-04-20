🚛 Recolecta API

Recolecta es una API de gestión integral diseñada para administrar múltiples componentes de un sistema de recolección y mantenimiento. Permite manejar usuarios con autenticación, rutas, camiones, mantenimientos, reportes, notificaciones y alertas, centralizando toda la lógica del sistema en una arquitectura robusta y escalable.


📌 Características principales

Autenticación de usuarios mediante JWT

Gestión de usuarios y roles

Administración de rutas, camiones y tipos de camiones

Gestión de mantenimientos y reportes

Sistema de notificaciones y alertas

Control de incidencias y fallas críticas

Arquitectura Hexagonal

Manejo de CORS

Conexión a base de datos PostgreSQL



🧰 Tecnologías utilizadas

Lenguaje: Go 1.24

Framework Web: Gin

Base de datos: PostgreSQL

Driver de BD: pgx v5

Autenticación: JWT

Variables de entorno: godotenv

Arquitectura: Hexagonal


Dependencias principales

github.com/gin-gonic/gin

github.com/jackc/pgx/v5

github.com/golang-jwt/jwt/v5

github.com/joho/godotenv

github.com/gin-contrib/cors



🏗️ Arquitectura Hexagonal

El proyecto sigue el enfoque de Arquitectura Hexagonal, separando claramente las responsabilidades y permitiendo un sistema desacoplado, mantenible y fácil de escalar.

📂 Estructura de carpetas


.
├── domain
│   └── Entidades y reglas de negocio
├── application
│   └── Casos de uso
├── infrastructure
│   └── Controladores, base de datos, HTTP, JWT, CORS
└── main.go




🔐 Seguridad

JWT para autenticación y autorización

CORS configurado para permitir clientes frontend específicos

Uso de variables de entorno para datos sensibles



📊 Tablas del sistema

La API gestiona las siguientes entidades:

rol
usuario 
domicilio 
colonia 
alerta_usuario 
aviso_general 
notificacion (parte de usuario) 
tipo_camion
camion
estado_camion
historial_asignacion_camion
ruta
punto_recoleccion
ruta_camion
relleno_sanitario 
registro_vaciado 
tipo_mantenimiento 
alerta_mantenimiento 
registro_mantenimiento 
incidencia 
reporte_falla_critica 
seguimiento_falla_critica 
anomalia 
reporte_conductor 
reporte_mantenimiento_generado 
notificacion (parte de fallas/mantenimiento)




🌐 Endpoints
Endpoint base
http://localhost:8000/api

Métodos HTTP soportados

GET

POST

PUT

PATCH

DELETE

Los endpoints están organizados por dominio siguiendo la arquitectura hexagonal.




⚙️ Variables de entorno

Crear un archivo .env en la raíz del proyecto con el siguiente contenido:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=1234
DB_NAME=proyecto_recolecta

CORS_ORIGINS=http://localhost:3000,http://localhost:5173,http://localhost:8000

JWT_SECRET=Halo_odt_es_un_gran_juego



Ejecución del proyecto

Instalar dependencias:

go mod tidy

Ejecutar la API:

go run .


o

go run main.go


La API quedará disponible en:

http://localhost:8000


📄 Notas adicionales

La API está pensada para ser consumida por aplicaciones web o móviles

Facilita la escalabilidad y el mantenimiento gracias a su arquitectura

Ideal para sistemas municipales, de logística o gestión urbana


## Documentacion Swagger para equipo mobile

La documentacion OpenAPI/Swagger ya fue generada en la carpeta [docs/swagger.json](docs/swagger.json) y [docs/swagger.yaml](docs/swagger.yaml).

### Como levantar y consultar Swagger UI

1. Levanta la API en local (puerto actual: `8080`).
2. Abre en navegador: `http://localhost:8080/swagger/index.html`.
3. Usa esa UI para revisar endpoints, parametros, respuestas y modelos.

### Archivos utiles para integracion mobile

- Especificacion JSON: [docs/swagger.json](docs/swagger.json)
- Especificacion YAML: [docs/swagger.yaml](docs/swagger.yaml)

Estos archivos pueden importarse directamente en herramientas como Postman, Insomnia o generadores de cliente (OpenAPI Generator / Swagger Codegen).

### Autenticacion para endpoints protegidos

La API usa header `Authorization` con formato:

`Authorization: Bearer <token_jwt>`

Recomendacion para mobile: centralizar el manejo del token en un interceptor/filtro HTTP para adjuntarlo automaticamente a requests autenticadas.

### Regenerar documentacion cuando cambie la API

Ejecuta este comando en la raiz del proyecto:

```bash
go run github.com/swaggo/swag/cmd/swag@v1.16.6 init -g engine.go -o docs --parseDependency --parseInternal
```