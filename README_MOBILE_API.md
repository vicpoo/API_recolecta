# API Recolecta — Guía para Aplicación Mobile

Este documento especifica los endpoints y métodos necesarios para integrar la aplicación mobile con la API Recolecta.

## Base URL

- Local: `http://localhost:8080`
- Prefijo base usado en la API: `/api`

## Autenticación

La API utiliza **JWT (JSON Web Tokens)** para autenticación y autorización:

- **Esquema**: `Authorization: Bearer <token>`
- **Ubicación del token**: Header de todas las solicitudes protegidas
- **Tipos de usuario**: Ciudadano (sin rol específico) y Empleado (con rol asignado)

## Convenciones generales

- `Content-Type: application/json` para `POST`, `PATCH`.
- Fechas en formato **ISO 8601** (ejemplo: `2026-05-01T12:00:00Z`).
- Todas las contraseñas se almacenan hasheadas con **bcrypt** (no se exponen en respuestas).

### Respuesta de error estandarizada

Todas las respuestas de error siguen esta estructura:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Mensaje legible para la app",
    "details": { /* opcional */ }
  }
}
```

**Campos importantes:**
- `code`: código nominal (string) para manejar lógica en la app
- `message`: texto legible para mostrar al usuario
- `details`: información adicional contextual

### Códigos de error disponibles

| Código | HTTP | Descripción |
|---|---:|---|
| `VALIDATION_ERROR` | 400 | Validación de entrada fallida |
| `INVALID_INPUT` | 400 | Parámetro inválido |
| `BAD_REQUEST` | 400 | Solicitud incorrecta |
| `NOT_FOUND` | 404 | Recurso no encontrado |
| `UNAUTHORIZED` | 401 | Autenticación requerida o inválida |
| `FORBIDDEN` | 403 | Acceso denegado |
| `CONFLICT` | 409 | Conflicto con datos existentes |
| `INTERNAL_SERVER_ERROR` | 500 | Error interno del servidor |
| `DATABASE_ERROR` | 500 | Error de base de datos |
| `OPERATION_FAILED` | 500 | Operación fallida |

---

---

# AUTENTICACIÓN

## 1. Login de Ciudadano

### 1.1 Descripción general del login de ciudadano

El sistema de autenticación para ciudadanos permite que usuarios sin rol administrativo se autentiquen en la plataforma usando su correo electrónico. 

**Características principales:**
- **Identificador**: Usa únicamente `email`
- **Contraseña**: Se valida mediante bcrypt para mayor seguridad
- **Token generado**: Contiene solo `user_id` (sin rol específico)
- **Acceso**: Una vez autenticados, pueden acceder a recursos de ciudadanos con protección JWT

### 1.2 Endpoint: Login de Ciudadano

**Método**: `POST`  
**URL**: `/api/ciudadanos/login`  
**Ejemplo completo**: `http://localhost:8080/api/ciudadanos/login`  
**Headers**: `Content-Type: application/json`  
**Autenticación**: No requerida  

#### Request

```json
{
  "email": "juan.garcia@correo.com",
  "password": "Segura123!"
}
```

**Validaciones en servidor:**
- `email` no puede estar vacío (se trimea)
- `password` no puede estar vacío (se trimea)
- Se convierte a minúsculas para búsqueda case-insensitive
- La contraseña se valida contra el hash bcrypt almacenado

#### Response 200 OK

```json
{
  "message": "login correcto",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "data": {
    "id": 42,
    "email": "juan.garcia@correo.com",
    "alias": "juangarcia",
    "created_at": "2026-04-15T10:30:00Z"
  }
}
```

**Token JWT incluye:**
- `sub`: ID del ciudadano (42)
- `role_id`: 0 (sin rol administrativo)
- `exp`: Tiempo de expiración (usualmente 24h o configurado)


---

## 2. Login de Empleado

### 2.1 Descripción general del login de empleado

El sistema de autenticación para empleados permite que usuarios con roles administrativos se autentiquen en la plataforma. A diferencia de ciudadanos, los empleados tienen un rol asignado que determina sus permisos.

**Características principales:**
- **Identificador**: Usa únicamente `email`
- **Contraseña**: Se valida mediante bcrypt
- **Token generado**: Contiene `user_id` + `rol_id` (importante para autorización)
- **Acceso**: Después de autenticarse, pueden acceder a recursos protegidos según su rol
- **Estado de cuenta**: Solo empleados activos (no desactivados) pueden acceder

**Diferencias clave con login de ciudadano:**
- Incluye validación de rol (`rol_id`)
- Incluye verificación de estado de activación (`desactivado`)
- El token incluye `rol_id` para decisiones en frontend/backend
- Los permisos dependen del rol asignado

### 2.2 Endpoint: Login de Empleado

**Método**: `POST`  
**URL**: `/api/empleados/login`  
**Ejemplo completo**: `http://localhost:8080/api/empleados/login`  
**Headers**: `Content-Type: application/json`  
**Autenticación**: No requerida  

#### Request

```json
{
  "email": "luis.perez@recolecta.com",
  "password": "AdminPass123!"
}
```

**Validaciones en servidor:**
- `email` no puede estar vacío (se trimea)
- `password` no puede estar vacío (se trimea)
- Se convierte a minúsculas para búsqueda case-insensitive
- Se valida que el empleado no esté desactivado
- La contraseña se valida contra el hash bcrypt almacenado

#### Response 200 OK

```json
{
  "message": "login correcto",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsInJvbGVfaWQiOjIsImV4cCI6MTc0NzM...",
  "data": {
    "id": 1,
    "nombre": "Luis",
    "apellidos": "Pérez",
    "mail": "luis.perez@recolecta.com",
    "username": "luis.perez",
    "desactivado": false,
    "rol_id": 2,
    "created_at": "2026-05-01T12:00:00Z",
    "updated_at": "2026-05-01T12:00:00Z"
  }
}
```

**Token JWT incluye:**
- `sub`: ID del empleado (1)
- `role_id`: Rol del empleado (2 = gerente, 3 = supervisor, etc.)
- `exp`: Tiempo de expiración


---

# OPERACIONES CRUD

## 3. Ciudadanos - Operaciones CRUD

### 3.1 Tabla de resumen

| Método | URL | Descripción | Autenticación |
|---|---|---|---|
| POST | `/api/ciudadanos` | Crear ciudadano (registro) | No | 
| GET | `/api/ciudadanos` | Listar ciudadanos | Sí (JWT) | Admin
| GET | `/api/ciudadanos/:id` | Obtener ciudadano por ID | Sí (JWT) | Admin
| PATCH | `/api/ciudadanos/:id` | Actualizar ciudadano | Sí (JWT) | Admin
| DELETE | `/api/ciudadanos/:id` | Eliminar ciudadano | Sí (JWT) | Admin
| POST | `/api/ciudadanos/login` | Login de ciudadano | No | 

### 3.2 Objeto Ciudadano

Estructura JSON estándar devuelta por la API:

```json
{
  "id": 42,
  "email": "juan.garcia@correo.com",
  "alias": "juangarcia",
  "created_at": "2026-04-15T10:30:00Z"
}
```

**Campos:**
- `id`: Identificador único (auto-generado en BD)
- `email`: Email único del ciudadano
- `alias`: Alias único (nombre de usuario alternativo)
- `created_at`: Fecha de creación (ISO 8601)
- `password`: NO se expone en respuestas

### 3.3 Crear ciudadano (Registro)

**Método**: `POST`  
**URL**: `/api/ciudadanos`  
**Headers**: `Content-Type: application/json`  
**Autenticación**: No (registro abierto)  

#### Request

```json
{
  "email": "juan.garcia@correo.com",
  "alias": "juangarcia",
  "password": "MiPassword123!",
  "fcm_token": "fcm-device-token-abc123"
}
```

**Validaciones:**
- `email` debe ser válido y único
- `alias` debe ser único
- `password` debe tener mínimo X caracteres (validar según reglas del servidor)
- `fcm_token` es requerido
- El token FCM se almacena en Redis con clave `fcm:ciudadano:<id>`

#### Response 201 Created

```json
{
  "message": "ciudadano creado correctamente",
  "id": 42
}
```

---

### 3.4 Listar ciudadanos

**Método**: `GET`  
**URL**: `/api/ciudadanos`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido)  

#### Response 200 OK

```json
{
  "data": [
    {
      "id": 42,
      "email": "juan.garcia@correo.com",
      "alias": "juangarcia",
      "created_at": "2026-04-15T10:30:00Z"
    },
    {
      "id": 43,
      "email": "maria.lopez@correo.com",
      "alias": "marialopez",
      "created_at": "2026-04-16T14:20:00Z"
    }
  ]
}
```

---

### 3.5 Obtener ciudadano por ID

**Método**: `GET`  
**URL**: `/api/ciudadanos/:id`  
**Ejemplo**: `http://localhost:8080/api/ciudadanos/42`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido)  

#### Response 200 OK

```json
{
  "data": {
    "id": 42,
    "email": "juan.garcia@correo.com",
    "alias": "juangarcia",
    "created_at": "2026-04-15T10:30:00Z"
  }
}
```

---

### 3.6 Actualizar ciudadano

**Método**: `PATCH`  
**URL**: `/api/ciudadanos/:id`  
**Ejemplo**: `http://localhost:8080/api/ciudadanos/42`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido)  

#### Request

```json
{
  "email": "juan.nuevo@correo.com",
  "alias": "juannuevo",
  "password": "NuevaPassword456!"
}
```

**Notas:**
- Solo se actualizan los campos proporcionados (PATCH, no PUT)
- Los campos omitidos mantienen su valor actual
- Todos los campos son opcionales

#### Response 200 OK

```json
{
  "message": "ciudadano actualizado correctamente",
  "data": {
    "id": 42,
    "email": "juan.nuevo@correo.com",
    "alias": "juannuevo",
    "created_at": "2026-04-15T10:30:00Z"
  }
}
```

---

### 3.7 Eliminar ciudadano

**Método**: `DELETE`  
**URL**: `/api/ciudadanos/:id`  
**Ejemplo**: `http://localhost:8080/api/ciudadanos/42`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido)  

#### Response 200 OK

```json
{
  "message": "ciudadano eliminado correctamente"
}
```

---

## 4. Empleados - Operaciones CRUD

### 4.1 Tabla de resumen

| Método | URL | Descripción | Autenticación | Rol requerido |
|---|---|---|---|---|
| POST | `/api/empleados` | Crear empleado | Sí (JWT) | ADMIN |
| GET | `/api/empleados` | Listar empleados | Sí (JWT) | ADMIN |
| GET | `/api/empleados/:id` | Obtener empleado por ID | Sí (JWT) | ADMIN |
| PATCH | `/api/empleados/:id` | Actualizar empleado | Sí (JWT) | ADMIN |
| DELETE | `/api/empleados/:id` | Eliminar empleado | Sí (JWT) | ADMIN |
| POST | `/api/empleados/login` | Login de empleado | No | N/A |

**Nota importante**: Todas las operaciones CRUD de empleados requieren rol `ADMIN`. Solo usuarios con `rol_id` de administrador pueden crear, listar, actualizar o eliminar empleados.

### 4.2 Objeto Empleado

Estructura JSON estándar devuelta por la API:

```json
{
  "id": 1,
  "nombre": "Luis",
  "apellidos": "Pérez",
  "mail": "luis.perez@recolecta.com",
  "username": "luis.perez",
  "desactivado": false,
  "rol_id": 2,
  "created_at": "2026-05-01T12:00:00Z",
  "updated_at": "2026-05-01T12:00:00Z",
  "deleted_at": null
}
```

**Campos:**
- `id`: Identificador único
- `nombre`: Nombre del empleado
- `apellidos`: Apellidos del empleado
- `mail`: Email único del empleado
- `username`: Username único (alternativa para login)
- `desactivado`: Boolean que indica si la cuenta está activa
- `rol_id`: ID del rol asignado (determina permisos)
- `created_at`: Fecha de creación
- `updated_at`: Última actualización
- `deleted_at`: Fecha de eliminación (si aplica soft-delete)
- `password`: NO se expone en respuestas

### 4.3 Crear empleado

**Método**: `POST`  
**URL**: `/api/empleados`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con rol ADMIN)  

#### Request

```json
{
  "nombre": "Carlos",
  "apellidos": "González",
  "mail": "carlos.gonzalez@recolecta.com",
  "username": "carlos.gonzalez",
  "password": "SeguraPass123!",
  "rol_id": 3,
  "desactivado": false
}
```

**Validaciones:**
- `nombre` y `apellidos` requeridos
- `mail` debe ser válido y único
- `username` debe ser único
- `password` debe cumplir requisitos mínimos
- `rol_id` debe ser válido (existir en tabla roles)
- `desactivado` por defecto es false

#### Response 201 Created

```json
{
  "message": "empleado creado correctamente",
  "data": {
    "id": 2,
    "nombre": "Carlos",
    "apellidos": "González",
    "mail": "carlos.gonzalez@recolecta.com",
    "username": "carlos.gonzalez",
    "desactivado": false,
    "rol_id": 3,
    "created_at": "2026-05-02T09:15:00Z",
    "updated_at": "2026-05-02T09:15:00Z",
    "deleted_at": null
  }
}
```


---

### 4.4 Listar empleados

**Método**: `GET`  
**URL**: `/api/empleados`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con rol ADMIN)  

#### Response 200 OK

```json
{
  "data": [
    {
      "id": 1,
      "nombre": "Luis",
      "apellidos": "Pérez",
      "mail": "luis.perez@recolecta.com",
      "username": "luis.perez",
      "desactivado": false,
      "rol_id": 2,
      "created_at": "2026-05-01T12:00:00Z",
      "updated_at": "2026-05-01T12:00:00Z",
      "deleted_at": null
    },
    {
      "id": 2,
      "nombre": "Carlos",
      "apellidos": "González",
      "mail": "carlos.gonzalez@recolecta.com",
      "username": "carlos.gonzalez",
      "desactivado": false,
      "rol_id": 3,
      "created_at": "2026-05-02T09:15:00Z",
      "updated_at": "2026-05-02T09:15:00Z",
      "deleted_at": null
    }
  ]
}
```

---

### 4.5 Obtener empleado por ID

**Método**: `GET`  
**URL**: `/api/empleados/:id`  
**Ejemplo**: `http://localhost:8080/api/empleados/1`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con rol ADMIN)  

#### Response 200 OK

```json
{
  "data": {
    "id": 1,
    "nombre": "Luis",
    "apellidos": "Pérez",
    "mail": "luis.perez@recolecta.com",
    "username": "luis.perez",
    "desactivado": false,
    "rol_id": 2,
    "created_at": "2026-05-01T12:00:00Z",
    "updated_at": "2026-05-01T12:00:00Z",
    "deleted_at": null
  }
}
```

---

### 4.6 Actualizar empleado

**Método**: `PATCH`  
**URL**: `/api/empleados/:id`  
**Ejemplo**: `http://localhost:8080/api/empleados/1`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con rol ADMIN)  

#### Request

```json
{
  "nombre": "Luis Alberto",
  "mail": "luis.alberto@recolecta.com",
  "password": "NuevaPassword789!",
  "rol_id": 4,
  "desactivado": false
}
```

**Notas:**
- PATCH permite actualizar parcialmente (solo enviar campos a cambiar)
- Si se proporciona `password`, se hashea antes de guardar
- Los campos omitidos mantienen su valor actual

#### Response 200 OK

```json
{
  "message": "empleado actualizado correctamente",
  "data": {
    "id": 1,
    "nombre": "Luis Alberto",
    "apellidos": "Pérez",
    "mail": "luis.alberto@recolecta.com",
    "username": "luis.perez",
    "desactivado": false,
    "rol_id": 4,
    "created_at": "2026-05-01T12:00:00Z",
    "updated_at": "2026-05-02T14:30:00Z",
    "deleted_at": null
  }
}
```


---

### 4.7 Eliminar empleado

**Método**: `DELETE`  
**URL**: `/api/empleados/:id`  
**Ejemplo**: `http://localhost:8080/api/empleados/1`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con rol ADMIN)  

#### Response 200 OK

```json
{
  "message": "empleado eliminado correctamente"
}
```

---

## 5. Alertas de Usuario - Operaciones

### 5.1 Tabla de resumen

| Método | URL | Descripción | Autenticación | Rol requerido |
|---|---|---|---|---|
| POST | `/api/alertas` | Crear alerta de usuario | Sí (JWT) | SUPERVISOR |
| GET | `/api/alertas` | Listar mis alertas | Sí (JWT) | N/A |
| PUT | `/api/alertas/:id/leida` | Marcar alerta como leída | Sí (JWT) | N/A |


### 5.2 Objeto AlertaUsuario

Estructura JSON estándar devuelta por la API:

```json
{
  "alerta_id": 1,
  "usuario_id": 42,
  "titulo": "Mantenimiento programado",
  "mensaje": "El sistema estará en mantenimiento el domingo de 02:00 a 04:00",
  "leida": false,
  "creado_por": 5,
  "created_at": "2026-05-03T10:30:00Z"
}
```

**Campos:**
- `alerta_id`: Identificador único de la alerta
- `usuario_id`: ID del usuario receptor (ciudadano o empleado)
- `titulo`: Título de la alerta (corto y descriptivo)
- `mensaje`: Contenido completo de la alerta
- `leida`: Boolean indicando si el usuario ha leído la alerta
- `creado_por`: ID del supervisor que creó la alerta
- `created_at`: Fecha de creación (ISO 8601)

### 5.3 Crear alerta

**Método**: `POST`  
**URL**: `/api/alertas`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con rol SUPERVISOR)  

#### Request

```json
{
  "titulo": "Actualización de seguridad",
  "mensaje": "Se ha lanzado una actualización de seguridad crítica. Por favor, reinicia la app.",
  "usuario_id": 42
}
```

**Validaciones:**
- `titulo` requerido (no vacío)
- `mensaje` requerido (no vacío)
- `usuario_id` requerido (debe ser numérico)
- Solo supervisores pueden crear alertas (rol SUPERVISOR)

#### Response 201 Created

```json
{
  "message": "Alerta creada exitosamente"
}
```

---

### 5.4 Listar mis alertas

**Método**: `GET`  
**URL**: `/api/alertas`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido)  

#### Response 200 OK

```json
{
  "data": [
    {
      "alerta_id": 1,
      "usuario_id": 42,
      "titulo": "Mantenimiento programado",
      "mensaje": "El sistema estará en mantenimiento el domingo de 02:00 a 04:00",
      "leida": false,
      "creado_por": 5,
      "created_at": "2026-05-03T10:30:00Z"
    },
    {
      "alerta_id": 2,
      "usuario_id": 42,
      "titulo": "Actualización de seguridad",
      "mensaje": "Se ha lanzado una actualización de seguridad crítica",
      "leida": true,
      "creado_por": 5,
      "created_at": "2026-05-02T15:45:00Z"
    }
  ]
}
```

---

### 5.5 Marcar alerta como leída

**Método**: `PUT`  
**URL**: `/api/alertas/:id/leida`  
**Ejemplo**: `http://localhost:8080/api/alertas/1/leida`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido)  

#### Response 200 OK

```json
{
  "message": "Alerta marcada como leída"
}
```


**Notas importantes:**
- Solo el usuario receptor de la alerta puede marcar como leída
- La validación de propiedad se realiza en el servidor
- Un usuario solo puede marcar como leída sus propias alertas

---

## 6. Domicilios - Operaciones

### 6.1 Resumen de endpoints

| Operación | Método | Endpoint | Autenticación |
|-----------|--------|----------|----------------|
| Crear domicilio | POST | `/api/domicilios` | JWT requerido |
| Listar mis domicilios | GET | `/api/domicilios` | JWT requerido |
| Obtener domicilio | GET | `/api/domicilios/:id` | JWT requerido |
| Actualizar domicilio | PATCH | `/api/domicilios/:id` | JWT requerido |
| Eliminar domicilio | DELETE | `/api/domicilios/:id` | JWT requerido |

---

### 6.2 Objeto Domicilio

```json
{
  "id": 1,
  "ciudadano_id": 42,
  "calle": "Avenida Principal 123",
  "numero_exterior": "123",
  "numero_interior": "Apt 5B",
  "colonia": "Centro",
  "ciudad": "Ciudad de México",
  "estado": "CDMX",
  "codigo_postal": "06500",
  "referencia": "Frente a la plaza",
  "tipo_domicilio": "Residencial",
  "activo": true,
  "created_at": "2026-05-01T10:30:00Z",
  "updated_at": "2026-05-02T14:45:00Z"
}
```

---

### 6.3 Crear domicilio

**Método**: `POST`  
**URL**: `/api/domicilios`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido)  

#### Request

```json
{
  "calle": "Avenida Principal 123",
  "numero_exterior": "123",
  "numero_interior": "Apt 5B",
  "colonia": "Centro",
  "ciudad": "Ciudad de México",
  "estado": "CDMX",
  "codigo_postal": "06500",
  "referencia": "Frente a la plaza",
  "tipo_domicilio": "Residencial",
  "activo": true
}
```

**Validaciones:**
- `calle` requerido (no vacío)
- `numero_exterior` requerido (no vacío)
- `ciudad` requerido (no vacío)
- `estado` requerido (no vacío)
- `codigo_postal` requerido (no vacío)
- `tipo_domicilio` requerido (no vacío)
- El `ciudadano_id` se obtiene automáticamente del token JWT

#### Response 201 Created

```json
{
  "message": "domicilio creado correctamente",
  "id": 1
}
```

---

### 6.4 Listar mis domicilios

**Método**: `GET`  
**URL**: `/api/domicilios`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido)  

#### Response 200 OK

```json
{
  "data": [
    {
      "id": 1,
      "ciudadano_id": 42,
      "calle": "Avenida Principal 123",
      "numero_exterior": "123",
      "numero_interior": "Apt 5B",
      "colonia": "Centro",
      "ciudad": "Ciudad de México",
      "estado": "CDMX",
      "codigo_postal": "06500",
      "referencia": "Frente a la plaza",
      "tipo_domicilio": "Residencial",
      "activo": true,
      "created_at": "2026-05-01T10:30:00Z",
      "updated_at": "2026-05-02T14:45:00Z"
    },
    {
      "id": 2,
      "ciudadano_id": 42,
      "calle": "Calle Secundaria 456",
      "numero_exterior": "456",
      "numero_interior": null,
      "colonia": "Polanco",
      "ciudad": "Ciudad de México",
      "estado": "CDMX",
      "codigo_postal": "11560",
      "referencia": "Entrada lateral",
      "tipo_domicilio": "Comercial",
      "activo": true,
      "created_at": "2026-04-28T08:15:00Z",
      "updated_at": "2026-05-01T11:20:00Z"
    }
  ]
}
```

---

### 6.5 Obtener domicilio por ID

**Método**: `GET`  
**URL**: `/api/domicilios/:id`  
**Ejemplo**: `http://localhost:8080/api/domicilios/1`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido)  

#### Response 200 OK

```json
{
  "data": {
    "id": 1,
    "ciudadano_id": 42,
    "calle": "Avenida Principal 123",
    "numero_exterior": "123",
    "numero_interior": "Apt 5B",
    "colonia": "Centro",
    "ciudad": "Ciudad de México",
    "estado": "CDMX",
    "codigo_postal": "06500",
    "referencia": "Frente a la plaza",
    "tipo_domicilio": "Residencial",
    "activo": true,
    "created_at": "2026-05-01T10:30:00Z",
    "updated_at": "2026-05-02T14:45:00Z"
  }
}
```

---

### 6.6 Actualizar domicilio

**Método**: `PATCH`  
**URL**: `/api/domicilios/:id`  
**Ejemplo**: `http://localhost:8080/api/domicilios/1`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido)  

#### Request

```json
{
  "calle": "Avenida Principal Actualizada 123",
  "numero_interior": "Apt 6C",
  "referencia": "Entrada frontal",
  "activo": true
}
```

**Notas:**
- PATCH permite actualizar parcialmente (solo enviar campos a cambiar)
- Los campos omitidos mantienen su valor actual
- Solo el propietario del domicilio puede actualizarlo

#### Response 200 OK

```json
{
  "message": "domicilio actualizado correctamente"
}
```

---

### 6.7 Eliminar domicilio

**Método**: `DELETE`  
**URL**: `/api/domicilios/:id`  
**Ejemplo**: `http://localhost:8080/api/domicilios/1`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido)  

#### Response 200 OK

```json
{
  "message": "domicilio eliminado correctamente"
}
```

**Notas importantes:**
- Solo el propietario del domicilio puede eliminarlo
- La validación de propiedad se realiza en el servidor
- Un usuario solo puede eliminar sus propios domicilios
- La eliminación es permanente

---

### 10.0 Anomalías

Base: `/api/anomalias`

| Método | URL | Descripción | Roles autorizados |
|---|---|---|---|
| POST | `/api/anomalias/` | Crear anomalía | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/anomalias/` | Listar anomalías | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/anomalias/:id` | Obtener anomalía por ID | ADMIN, SUPERVISOR, COORDINADOR |
| PUT | `/api/anomalias/:id` | Actualizar anomalía | ADMIN, SUPERVISOR, COORDINADOR |
| DELETE | `/api/anomalias/:id` | Eliminar anomalía | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/anomalias/punto/:puntoId` | Filtrar por punto | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/anomalias/chofer/:choferId` | Filtrar por chofer | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/anomalias/estado` | Filtrar por `?estado=` | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/anomalias/tipo` | Filtrar por `?tipo_anomalia=` | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/anomalias/por-fecha` | Filtrar por `?fecha_inicio=&fecha_fin=` | ADMIN, SUPERVISOR, COORDINADOR |

### 10.0.1 Objeto anomalía

```json
{
  "anomalia_id": 21,
  "punto_id": 8,
  "tipo_anomalia": "bloqueo",
  "descripcion": "Calle cerrada por obras",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "estado": "PENDIENTE",
  "fecha_resolucion": null,
  "id_chofer_id": 5
}
```

### 10.0.2 Crear anomalía

**Método**: `POST`  
**URL**: `/api/anomalias/`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "punto_id": 8,
  "tipo_anomalia": "bloqueo",
  "descripcion": "Calle cerrada por obras",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "estado": "PENDIENTE",
  "id_chofer_id": 5
}
```

#### Response 201 Created

```json
{
  "anomalia_id": 21,
  "punto_id": 8,
  "tipo_anomalia": "bloqueo",
  "descripcion": "Calle cerrada por obras",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "estado": "PENDIENTE",
  "fecha_resolucion": null,
  "id_chofer_id": 5
}
```

### 10.0.3 Listar anomalías

**Método**: `GET`  
**URL**: `/api/anomalias/`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "anomalia_id": 21,
      "punto_id": 8,
      "tipo_anomalia": "bloqueo",
      "descripcion": "Calle cerrada por obras",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "estado": "PENDIENTE",
      "fecha_resolucion": null,
      "id_chofer_id": 5
    }
  ]
}
```

### 10.0.4 Obtener anomalía por ID

**Método**: `GET`  
**URL**: `/api/anomalias/:id`  
**Ejemplo**: `http://localhost:8080/api/anomalias/21`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "anomalia_id": 21,
  "punto_id": 8,
  "tipo_anomalia": "bloqueo",
  "descripcion": "Calle cerrada por obras",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "estado": "PENDIENTE",
  "fecha_resolucion": null,
  "id_chofer_id": 5
}
```

### 10.0.5 Actualizar anomalía

**Método**: `PUT`  
**URL**: `/api/anomalias/:id`  
**Ejemplo**: `http://localhost:8080/api/anomalias/21`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "punto_id": 8,
  "tipo_anomalia": "bloqueo",
  "descripcion": "Calle cerrada por obras corregida",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "estado": "RESUELTA",
  "fecha_resolucion": "2026-04-28T10:00:00Z",
  "id_chofer_id": 5
}
```

#### Response 200 OK

```json
{
  "anomalia_id": 21,
  "punto_id": 8,
  "tipo_anomalia": "bloqueo",
  "descripcion": "Calle cerrada por obras corregida",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "estado": "RESUELTA",
  "fecha_resolucion": "2026-04-28T10:00:00Z",
  "id_chofer_id": 5
}
```

### 10.0.6 Eliminar anomalía

**Método**: `DELETE`  
**URL**: `/api/anomalias/:id`  
**Ejemplo**: `http://localhost:8080/api/anomalias/21`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "message": "Anomalía eliminada exitosamente"
}
```

### 10.0.7 Filtrar por punto

**Método**: `GET`  
**URL**: `/api/anomalias/punto/:puntoId`  
**Ejemplo**: `http://localhost:8080/api/anomalias/punto/8`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "anomalia_id": 21,
      "punto_id": 8,
      "tipo_anomalia": "bloqueo",
      "descripcion": "Calle cerrada por obras",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "estado": "PENDIENTE",
      "fecha_resolucion": null,
      "id_chofer_id": 5
    }
  ]
}
```

### 10.0.8 Filtrar por chofer

**Método**: `GET`  
**URL**: `/api/anomalias/chofer/:choferId`  
**Ejemplo**: `http://localhost:8080/api/anomalias/chofer/5`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "anomalia_id": 21,
      "punto_id": 8,
      "tipo_anomalia": "bloqueo",
      "descripcion": "Calle cerrada por obras",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "estado": "PENDIENTE",
      "fecha_resolucion": null,
      "id_chofer_id": 5
    }
  ]
}
```

### 10.0.9 Filtrar por estado

**Método**: `GET`  
**URL**: `/api/anomalias/estado?estado=PENDIENTE`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "anomalia_id": 21,
      "punto_id": 8,
      "tipo_anomalia": "bloqueo",
      "descripcion": "Calle cerrada por obras",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "estado": "PENDIENTE",
      "fecha_resolucion": null,
      "id_chofer_id": 5
    }
  ]
}
```

### 10.0.10 Filtrar por tipo

**Método**: `GET`  
**URL**: `/api/anomalias/tipo?tipo_anomalia=bloqueo`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "anomalia_id": 21,
      "punto_id": 8,
      "tipo_anomalia": "bloqueo",
      "descripcion": "Calle cerrada por obras",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "estado": "PENDIENTE",
      "fecha_resolucion": null,
      "id_chofer_id": 5
    }
  ]
}
```

### 10.0.11 Filtrar por rango de fechas

**Método**: `GET`  
**URL**: `/api/anomalias/por-fecha?fecha_inicio=&fecha_fin=`  
**Ejemplo**: `http://localhost:8080/api/anomalias/por-fecha?fecha_inicio=2026-04-01T00:00:00Z&fecha_fin=2026-04-30T23:59:59Z`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "anomalia_id": 21,
      "punto_id": 8,
      "tipo_anomalia": "bloqueo",
      "descripcion": "Calle cerrada por obras",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "estado": "PENDIENTE",
      "fecha_resolucion": null,
      "id_chofer_id": 5
    }
  ]
}
```

---

### 10.1 Reportes de Conductor

| Método | URL | Descripción | Roles autorizados |
|---|---|---|---|
| POST | `/api/reportes-conductor/` | Crear reporte de conductor | CONDUCTOR, ADMIN, COORDINADOR, SUPERVISOR |
| GET | `/api/reportes-conductor/` | Listar reportes | CONDUCTOR, ADMIN, COORDINADOR, SUPERVISOR |
| GET | `/api/reportes-conductor/:id` | Obtener reporte por ID | CONDUCTOR, ADMIN, COORDINADOR, SUPERVISOR |
| PUT | `/api/reportes-conductor/:id` | Actualizar reporte | CONDUCTOR, ADMIN, COORDINADOR, SUPERVISOR |
| DELETE | `/api/reportes-conductor/:id` | Eliminar reporte | CONDUCTOR, ADMIN, COORDINADOR, SUPERVISOR |
| GET | `/api/reportes-conductor/camion/:camion_id` | Filtrar por camion | CONDUCTOR, ADMIN, COORDINADOR, SUPERVISOR |
| GET | `/api/reportes-conductor/conductor/:conductor_id` | Filtrar por conductor | CONDUCTOR, ADMIN, COORDINADOR, SUPERVISOR |
| GET | `/api/reportes-conductor/ruta/:ruta_id` | Filtrar por ruta | CONDUCTOR, ADMIN, COORDINADOR, SUPERVISOR |
| GET | `/api/reportes-conductor/fecha?fecha_inicio=&fecha_fin=` | Filtrar por rango de fechas (ISO8601) | CONDUCTOR, ADMIN, COORDINADOR, SUPERVISOR |

### 10.1.1 Crear reporte

**Método**: `POST`  
**URL**: `/api/reportes-conductor/`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, ADMIN, COORDINADOR o SUPERVISOR)  

#### Request

```json
{
  "conductor_id": 5,
  "camion_id": 12,
  "ruta_id": 3,
  "descripcion": "Conductor reporta bloqueo en ruta"
}
```

Response ejemplo 201 Created:

```json
{
  "data": {
    "reporte_id": 101,
    "conductor_id": 5,
    "camion_id": 12,
    "ruta_id": 3,
    "descripcion": "Conductor reporta bloqueo en ruta",
    "created_at": "2026-05-11T08:30:00Z"
  }
}
```

### 10.1.2 Listar reportes

**Método**: `GET`  
**URL**: `/api/reportes-conductor/`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, ADMIN, COORDINADOR o SUPERVISOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "reporte_id": 101,
      "conductor_id": 5,
      "camion_id": 12,
      "ruta_id": 3,
      "descripcion": "Conductor reporta bloqueo en ruta",
      "created_at": "2026-05-11T08:30:00Z"
    },
    {
      "reporte_id": 102,
      "conductor_id": 5,
      "camion_id": 15,
      "ruta_id": 4,
      "descripcion": "Conductor reporta retraso por falla mecánica",
      "created_at": "2026-05-11T09:10:00Z"
    }
  ]
}
```

### 10.1.3 Obtener reporte por ID

**Método**: `GET`  
**URL**: `/api/reportes-conductor/:id`  
**Ejemplo**: `http://localhost:8080/api/reportes-conductor/101`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, ADMIN, COORDINADOR o SUPERVISOR)  

#### Response 200 OK

```json
{
  "data": {
    "reporte_id": 101,
    "conductor_id": 5,
    "camion_id": 12,
    "ruta_id": 3,
    "descripcion": "Conductor reporta bloqueo en ruta",
    "created_at": "2026-05-11T08:30:00Z"
  }
}
```

### 10.1.4 Actualizar reporte

**Método**: `PUT`  
**URL**: `/api/reportes-conductor/:id`  
**Ejemplo**: `http://localhost:8080/api/reportes-conductor/101`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, ADMIN, COORDINADOR o SUPERVISOR)  

#### Request

```json
{
  "conductor_id": 5,
  "camion_id": 12,
  "ruta_id": 3,
  "descripcion": "Conductor reporta bloqueo levantado y ruta reanudada"
}
```

#### Response 200 OK

```json
{
  "data": {
    "reporte_id": 101,
    "conductor_id": 5,
    "camion_id": 12,
    "ruta_id": 3,
    "descripcion": "Conductor reporta bloqueo levantado y ruta reanudada",
    "created_at": "2026-05-11T08:30:00Z"
  }
}
```

### 10.1.5 Eliminar reporte

**Método**: `DELETE`  
**URL**: `/api/reportes-conductor/:id`  
**Ejemplo**: `http://localhost:8080/api/reportes-conductor/101`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, ADMIN, COORDINADOR o SUPERVISOR)  

#### Response 200 OK

```json
{
  "status": "Reporte eliminado exitosamente"
}
```

### 10.1.6 Filtrar por camion

**Método**: `GET`  
**URL**: `/api/reportes-conductor/camion/:camion_id`  
**Ejemplo**: `http://localhost:8080/api/reportes-conductor/camion/12`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, ADMIN, COORDINADOR o SUPERVISOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "reporte_id": 101,
      "conductor_id": 5,
      "camion_id": 12,
      "ruta_id": 3,
      "descripcion": "Conductor reporta bloqueo en ruta",
      "created_at": "2026-05-11T08:30:00Z"
    },
    {
      "reporte_id": 104,
      "conductor_id": 8,
      "camion_id": 12,
      "ruta_id": 6,
      "descripcion": "Reporte de conductor por falla eléctrica",
      "created_at": "2026-05-11T12:40:00Z"
    }
  ]
}
```

### 10.1.7 Filtrar por conductor

**Método**: `GET`  
**URL**: `/api/reportes-conductor/conductor/:conductor_id`  
**Ejemplo**: `http://localhost:8080/api/reportes-conductor/conductor/5`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, ADMIN, COORDINADOR o SUPERVISOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "reporte_id": 101,
      "conductor_id": 5,
      "camion_id": 12,
      "ruta_id": 3,
      "descripcion": "Conductor reporta bloqueo en ruta",
      "created_at": "2026-05-11T08:30:00Z"
    },
    {
      "reporte_id": 103,
      "conductor_id": 5,
      "camion_id": 18,
      "ruta_id": 7,
      "descripcion": "Conductor reporta demora por lluvia",
      "created_at": "2026-05-11T10:05:00Z"
    }
  ]
}
```

### 10.1.8 Filtrar por ruta

**Método**: `GET`  
**URL**: `/api/reportes-conductor/ruta/:ruta_id`  
**Ejemplo**: `http://localhost:8080/api/reportes-conductor/ruta/3`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, ADMIN, COORDINADOR o SUPERVISOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "reporte_id": 101,
      "conductor_id": 5,
      "camion_id": 12,
      "ruta_id": 3,
      "descripcion": "Conductor reporta bloqueo en ruta",
      "created_at": "2026-05-11T08:30:00Z"
    },
    {
      "reporte_id": 105,
      "conductor_id": 9,
      "camion_id": 21,
      "ruta_id": 3,
      "descripcion": "Ruta con incidente vial reportado",
      "created_at": "2026-05-11T13:25:00Z"
    }
  ]
}
```

### 10.1.9 Filtrar por rango de fechas

**Método**: `GET`  
**URL**: `/api/reportes-conductor/fecha?fecha_inicio=&fecha_fin=`  
**Ejemplo**: `http://localhost:8080/api/reportes-conductor/fecha?fecha_inicio=2026-05-01T00:00:00Z&fecha_fin=2026-05-11T23:59:59Z`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, ADMIN, COORDINADOR o SUPERVISOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "reporte_id": 101,
      "conductor_id": 5,
      "camion_id": 12,
      "ruta_id": 3,
      "descripcion": "Conductor reporta bloqueo en ruta",
      "created_at": "2026-05-11T08:30:00Z"
    },
    {
      "reporte_id": 106,
      "conductor_id": 5,
      "camion_id": 19,
      "ruta_id": 8,
      "descripcion": "Conductor reporta unidad detenida por revisión",
      "created_at": "2026-05-10T16:45:00Z"
    }
  ]
}
```

### 10.2 Reportes de Falla Crítica

| Método | URL | Descripción | Roles autorizados |
|---|---|---|---|
| POST | `/api/reportes-falla-critica/` | Crear reporte de falla crítica | CONDUCTOR, SUPERVISOR, ADMIN, COORDINADOR |
| GET | `/api/reportes-falla-critica/` | Listar reportes | CONDUCTOR, SUPERVISOR, ADMIN, COORDINADOR |
| GET | `/api/reportes-falla-critica/:id` | Obtener reporte por ID | CONDUCTOR, SUPERVISOR, ADMIN, COORDINADOR |
| PUT | `/api/reportes-falla-critica/:id` | Actualizar reporte | CONDUCTOR, SUPERVISOR, ADMIN, COORDINADOR |
| DELETE | `/api/reportes-falla-critica/:id` | Eliminar reporte | CONDUCTOR, SUPERVISOR, ADMIN, COORDINADOR |
| GET | `/api/reportes-falla-critica/camion/:camionId` | Filtrar por camion | CONDUCTOR, SUPERVISOR, ADMIN, COORDINADOR |
| GET | `/api/reportes-falla-critica/conductor/:conductorId` | Filtrar por conductor | CONDUCTOR, SUPERVISOR, ADMIN, COORDINADOR |
| GET | `/api/reportes-falla-critica/por-fecha?fecha_inicio=&fecha_fin=` | Filtrar por rango de fechas (ISO8601) | CONDUCTOR, SUPERVISOR, ADMIN, COORDINADOR |

### 10.2.1 Crear reporte

**Método**: `POST`  
**URL**: `/api/reportes-falla-critica/`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, SUPERVISOR, ADMIN o COORDINADOR)  

#### Request

```json
{
  "camion_id": 12,
  "conductor_id": 5,
  "descripcion": "Falla en sistema de frenos"
}
```

Response ejemplo 201 Created:

```json
{
  "data": {
    "falla_id": 55,
    "camion_id": 12,
    "conductor_id": 5,
    "descripcion": "Falla en sistema de frenos",
    "created_at": "2026-05-11T09:00:00Z",
    "eliminado": false
  }
}
```

### 10.2.2 Listar reportes

**Método**: `GET`  
**URL**: `/api/reportes-falla-critica/`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, SUPERVISOR, ADMIN o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "falla_id": 55,
      "camion_id": 12,
      "conductor_id": 5,
      "descripcion": "Falla en sistema de frenos",
      "created_at": "2026-05-11T09:00:00Z",
      "eliminado": false
    },
    {
      "falla_id": 56,
      "camion_id": 15,
      "conductor_id": 8,
      "descripcion": "Pérdida de presión en llantas",
      "created_at": "2026-05-11T10:25:00Z",
      "eliminado": false
    }
  ]
}
```

### 10.2.3 Obtener reporte por ID

**Método**: `GET`  
**URL**: `/api/reportes-falla-critica/:id`  
**Ejemplo**: `http://localhost:8080/api/reportes-falla-critica/55`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, SUPERVISOR, ADMIN o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": {
    "falla_id": 55,
    "camion_id": 12,
    "conductor_id": 5,
    "descripcion": "Falla en sistema de frenos",
    "created_at": "2026-05-11T09:00:00Z",
    "eliminado": false
  }
}
```

### 10.2.4 Actualizar reporte

**Método**: `PUT`  
**URL**: `/api/reportes-falla-critica/:id`  
**Ejemplo**: `http://localhost:8080/api/reportes-falla-critica/55`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, SUPERVISOR, ADMIN o COORDINADOR)  

#### Request

```json
{
  "camion_id": 12,
  "conductor_id": 5,
  "descripcion": "Falla en sistema de frenos corregida"
}
```

#### Response 200 OK

```json
{
  "data": {
    "falla_id": 55,
    "camion_id": 12,
    "conductor_id": 5,
    "descripcion": "Falla en sistema de frenos corregida",
    "created_at": "2026-05-11T09:00:00Z",
    "eliminado": false
  }
}
```

### 10.2.5 Eliminar reporte

**Método**: `DELETE`  
**URL**: `/api/reportes-falla-critica/:id`  
**Ejemplo**: `http://localhost:8080/api/reportes-falla-critica/55`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, SUPERVISOR, ADMIN o COORDINADOR)  

#### Response 200 OK

```json
{
  "status": "Reporte de falla crítica marcado como eliminado exitosamente",
  "message": "El reporte ha sido marcado como eliminado (soft delete)"
}
```

### 10.2.6 Filtrar por camion

**Método**: `GET`  
**URL**: `/api/reportes-falla-critica/camion/:camionId`  
**Ejemplo**: `http://localhost:8080/api/reportes-falla-critica/camion/12`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, SUPERVISOR, ADMIN o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "falla_id": 55,
      "camion_id": 12,
      "conductor_id": 5,
      "descripcion": "Falla en sistema de frenos",
      "created_at": "2026-05-11T09:00:00Z",
      "eliminado": false
    },
    {
      "falla_id": 58,
      "camion_id": 12,
      "conductor_id": 9,
      "descripcion": "Válvula de aire con fuga",
      "created_at": "2026-05-11T12:15:00Z",
      "eliminado": false
    }
  ]
}
```

### 10.2.7 Filtrar por conductor

**Método**: `GET`  
**URL**: `/api/reportes-falla-critica/conductor/:conductorId`  
**Ejemplo**: `http://localhost:8080/api/reportes-falla-critica/conductor/5`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, SUPERVISOR, ADMIN o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "falla_id": 55,
      "camion_id": 12,
      "conductor_id": 5,
      "descripcion": "Falla en sistema de frenos",
      "created_at": "2026-05-11T09:00:00Z",
      "eliminado": false
    },
    {
      "falla_id": 59,
      "camion_id": 18,
      "conductor_id": 5,
      "descripcion": "Temperatura elevada en motor",
      "created_at": "2026-05-11T13:50:00Z",
      "eliminado": false
    }
  ]
}
```

### 10.2.8 Filtrar por rango de fechas

**Método**: `GET`  
**URL**: `/api/reportes-falla-critica/por-fecha?fecha_inicio=&fecha_fin=`  
**Ejemplo**: `http://localhost:8080/api/reportes-falla-critica/por-fecha?fecha_inicio=2026-05-01T00:00:00Z&fecha_fin=2026-05-11T23:59:59Z`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles CONDUCTOR, SUPERVISOR, ADMIN o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "falla_id": 55,
      "camion_id": 12,
      "conductor_id": 5,
      "descripcion": "Falla en sistema de frenos",
      "created_at": "2026-05-11T09:00:00Z",
      "eliminado": false
    },
    {
      "falla_id": 60,
      "camion_id": 21,
      "conductor_id": 9,
      "descripcion": "Sensor de temperatura fuera de rango",
      "created_at": "2026-05-10T18:20:00Z",
      "eliminado": false
    }
  ]
}
```

### 10.3 Seguimientos de Falla Crítica

| Método | URL | Descripción | Roles autorizados |
|---|---|---|---|
| POST | `/api/seguimientos-falla-critica/` | Crear seguimiento para una falla crítica | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/seguimientos-falla-critica/` | Listar seguimientos | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/seguimientos-falla-critica/:id` | Obtener seguimiento por ID | ADMIN, SUPERVISOR, COORDINADOR |
| PUT | `/api/seguimientos-falla-critica/:id` | Actualizar seguimiento | ADMIN, SUPERVISOR, COORDINADOR |
| DELETE | `/api/seguimientos-falla-critica/:id` | Eliminar seguimiento | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/seguimientos-falla-critica/falla/:fallaId` | Filtrar seguimientos por falla (fallaId) | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/seguimientos-falla-critica/por-fecha?fecha_inicio=&fecha_fin=` | Filtrar por rango de fechas (ISO8601) | ADMIN, SUPERVISOR, COORDINADOR |

### 10.3.1 Crear seguimiento

**Método**: `POST`  
**URL**: `/api/seguimientos-falla-critica/`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "falla_id": 7,
  "comentario": "Técnico asignado, en proceso de revisión"
}
```

Response ejemplo 201 Created:

```json
{
  "data": {
    "seguimiento_id": 201,
    "falla_id": 7,
    "comentario": "Técnico asignado, en proceso de revisión",
    "created_at": "2026-05-11T09:15:00Z"
  }
}
```

### 10.3.2 Listar seguimientos

**Método**: `GET`  
**URL**: `/api/seguimientos-falla-critica/`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "seguimiento_id": 201,
      "falla_id": 7,
      "comentario": "Técnico asignado, en proceso de revisión",
      "created_at": "2026-05-11T09:15:00Z"
    },
    {
      "seguimiento_id": 202,
      "falla_id": 7,
      "comentario": "Se reemplazó la pieza dañada y se hizo prueba",
      "created_at": "2026-05-11T11:05:00Z"
    }
  ]
}
```

### 10.3.3 Obtener seguimiento por ID

**Método**: `GET`  
**URL**: `/api/seguimientos-falla-critica/:id`  
**Ejemplo**: `http://localhost:8080/api/seguimientos-falla-critica/201`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": {
    "seguimiento_id": 201,
    "falla_id": 7,
    "comentario": "Técnico asignado, en proceso de revisión",
    "created_at": "2026-05-11T09:15:00Z"
  }
}
```

### 10.3.4 Actualizar seguimiento

**Método**: `PUT`  
**URL**: `/api/seguimientos-falla-critica/:id`  
**Ejemplo**: `http://localhost:8080/api/seguimientos-falla-critica/201`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "falla_id": 7,
  "comentario": "Se realizó ajuste final y quedó operativa"
}
```

#### Response 200 OK

```json
{
  "data": {
    "seguimiento_id": 201,
    "falla_id": 7,
    "comentario": "Se realizó ajuste final y quedó operativa",
    "created_at": "2026-05-11T09:15:00Z"
  }
}
```

### 10.3.5 Eliminar seguimiento

**Método**: `DELETE`  
**URL**: `/api/seguimientos-falla-critica/:id`  
**Ejemplo**: `http://localhost:8080/api/seguimientos-falla-critica/201`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "status": "Seguimiento de falla crítica eliminado exitosamente"
}
```

### 10.3.6 Filtrar seguimientos por falla

**Método**: `GET`  
**URL**: `/api/seguimientos-falla-critica/falla/:fallaId`  
**Ejemplo**: `http://localhost:8080/api/seguimientos-falla-critica/falla/7`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "seguimiento_id": 201,
      "falla_id": 7,
      "comentario": "Técnico asignado, en proceso de revisión",
      "created_at": "2026-05-11T09:15:00Z"
    },
    {
      "seguimiento_id": 202,
      "falla_id": 7,
      "comentario": "Se reemplazó la pieza dañada y se hizo prueba",
      "created_at": "2026-05-11T11:05:00Z"
    }
  ]
}
```

### 10.3.7 Filtrar por rango de fechas

**Método**: `GET`  
**URL**: `/api/seguimientos-falla-critica/por-fecha?fecha_inicio=&fecha_fin=`  
**Ejemplo**: `http://localhost:8080/api/seguimientos-falla-critica/por-fecha?fecha_inicio=2026-05-01T00:00:00Z&fecha_fin=2026-05-11T23:59:59Z`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "seguimiento_id": 201,
      "falla_id": 7,
      "comentario": "Técnico asignado, en proceso de revisión",
      "created_at": "2026-05-11T09:15:00Z"
    },
    {
      "seguimiento_id": 203,
      "falla_id": 9,
      "comentario": "Se dio seguimiento con el taller externo",
      "created_at": "2026-05-10T18:20:00Z"
    }
  ]
}
```

### 10.4 Registros de Mantenimiento

| Método | URL | Descripción | Roles autorizados |
|---|---|---|---|
| POST | `/api/registros-mantenimiento/` | Crear registro de mantenimiento | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/registros-mantenimiento/` | Listar registros | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/registros-mantenimiento/:id` | Obtener registro por ID | ADMIN, SUPERVISOR, COORDINADOR |
| PUT | `/api/registros-mantenimiento/:id` | Actualizar registro | ADMIN, SUPERVISOR, COORDINADOR |
| DELETE | `/api/registros-mantenimiento/:id` | Eliminar registro | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/registros-mantenimiento/alerta/:alerta_id` | Filtrar por alerta | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/registros-mantenimiento/camion/:camion_id` | Filtrar por camión | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/registros-mantenimiento/coordinador/:coordinador_id` | Filtrar por coordinador | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/registros-mantenimiento/fecha?fecha_inicio=&fecha_fin=` | Filtrar por rango de fechas | ADMIN, SUPERVISOR, COORDINADOR |

### 10.4.1 Objeto registro de mantenimiento

```json
{
  "registro_id": 41,
  "alerta_id": 12,
  "camion_id": 10,
  "coordinador_id": 3,
  "mecanico_responsable": "Juan Pérez",
  "fecha_realizada": "2026-05-01T08:00:00Z",
  "kilometraje_mantenimiento": 12450.5,
  "observaciones": "Cambio de aceite y revisión general",
  "created_at": "2026-05-01T08:30:00Z"
}
```

### 10.4.2 Crear registro de mantenimiento

**Método**: `POST`  
**URL**: `/api/registros-mantenimiento/`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "alerta_id": 12,
  "camion_id": 10,
  "coordinador_id": 3,
  "mecanico_responsable": "Juan Pérez",
  "fecha_realizada": "2026-05-01T08:00:00Z",
  "kilometraje_mantenimiento": 12450.5,
  "observaciones": "Cambio de aceite y revisión general"
}
```

#### Response 201 Created

```json
{
  "registro_id": 41,
  "alerta_id": 12,
  "camion_id": 10,
  "coordinador_id": 3,
  "mecanico_responsable": "Juan Pérez",
  "fecha_realizada": "2026-05-01T08:00:00Z",
  "kilometraje_mantenimiento": 12450.5,
  "observaciones": "Cambio de aceite y revisión general",
  "created_at": "2026-05-01T08:30:00Z"
}
```

### 10.4.3 Listar registros de mantenimiento

**Método**: `GET`  
**URL**: `/api/registros-mantenimiento/`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "registro_id": 41,
      "alerta_id": 12,
      "camion_id": 10,
      "coordinador_id": 3,
      "mecanico_responsable": "Juan Pérez",
      "fecha_realizada": "2026-05-01T08:00:00Z",
      "kilometraje_mantenimiento": 12450.5,
      "observaciones": "Cambio de aceite y revisión general",
      "created_at": "2026-05-01T08:30:00Z"
    }
  ]
}
```

### 10.4.4 Obtener registro por ID

**Método**: `GET`  
**URL**: `/api/registros-mantenimiento/:id`  
**Ejemplo**: `http://localhost:8080/api/registros-mantenimiento/41`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "registro_id": 41,
  "alerta_id": 12,
  "camion_id": 10,
  "coordinador_id": 3,
  "mecanico_responsable": "Juan Pérez",
  "fecha_realizada": "2026-05-01T08:00:00Z",
  "kilometraje_mantenimiento": 12450.5,
  "observaciones": "Cambio de aceite y revisión general",
  "created_at": "2026-05-01T08:30:00Z"
}
```

### 10.4.5 Actualizar registro de mantenimiento

**Método**: `PUT`  
**URL**: `/api/registros-mantenimiento/:id`  
**Ejemplo**: `http://localhost:8080/api/registros-mantenimiento/41`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "alerta_id": 12,
  "camion_id": 10,
  "coordinador_id": 3,
  "mecanico_responsable": "Juan Pérez",
  "fecha_realizada": "2026-05-01T09:00:00Z",
  "kilometraje_mantenimiento": 12460.5,
  "observaciones": "Se ajustó el sistema de frenos"
}
```

#### Response 200 OK

```json
{
  "registro_id": 41,
  "alerta_id": 12,
  "camion_id": 10,
  "coordinador_id": 3,
  "mecanico_responsable": "Juan Pérez",
  "fecha_realizada": "2026-05-01T09:00:00Z",
  "kilometraje_mantenimiento": 12460.5,
  "observaciones": "Se ajustó el sistema de frenos",
  "created_at": "2026-05-01T08:30:00Z"
}
```

### 10.4.6 Eliminar registro de mantenimiento

**Método**: `DELETE`  
**URL**: `/api/registros-mantenimiento/:id`  
**Ejemplo**: `http://localhost:8080/api/registros-mantenimiento/41`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "status": "Registro de mantenimiento eliminado exitosamente"
}
```

### 10.4.7 Filtrar por alerta

**Método**: `GET`  
**URL**: `/api/registros-mantenimiento/alerta/:alerta_id`  
**Ejemplo**: `http://localhost:8080/api/registros-mantenimiento/alerta/12`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "registro_id": 41,
  "alerta_id": 12,
  "camion_id": 10,
  "coordinador_id": 3,
  "mecanico_responsable": "Juan Pérez",
  "fecha_realizada": "2026-05-01T08:00:00Z",
  "kilometraje_mantenimiento": 12450.5,
  "observaciones": "Cambio de aceite y revisión general",
  "created_at": "2026-05-01T08:30:00Z"
}
```

### 10.4.8 Filtrar por camión

**Método**: `GET`  
**URL**: `/api/registros-mantenimiento/camion/:camion_id`  
**Ejemplo**: `http://localhost:8080/api/registros-mantenimiento/camion/10`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "registro_id": 41,
      "alerta_id": 12,
      "camion_id": 10,
      "coordinador_id": 3,
      "mecanico_responsable": "Juan Pérez",
      "fecha_realizada": "2026-05-01T08:00:00Z",
      "kilometraje_mantenimiento": 12450.5,
      "observaciones": "Cambio de aceite y revisión general",
      "created_at": "2026-05-01T08:30:00Z"
    }
  ]
}
```

### 10.4.9 Filtrar por coordinador

**Método**: `GET`  
**URL**: `/api/registros-mantenimiento/coordinador/:coordinador_id`  
**Ejemplo**: `http://localhost:8080/api/registros-mantenimiento/coordinador/3`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "registro_id": 41,
      "alerta_id": 12,
      "camion_id": 10,
      "coordinador_id": 3,
      "mecanico_responsable": "Juan Pérez",
      "fecha_realizada": "2026-05-01T08:00:00Z",
      "kilometraje_mantenimiento": 12450.5,
      "observaciones": "Cambio de aceite y revisión general",
      "created_at": "2026-05-01T08:30:00Z"
    }
  ]
}
```

### 10.4.10 Filtrar por rango de fechas

**Método**: `GET`  
**URL**: `/api/registros-mantenimiento/fecha?fecha_inicio=&fecha_fin=`  
**Ejemplo**: `http://localhost:8080/api/registros-mantenimiento/fecha?fecha_inicio=2026-05-01T00:00:00Z&fecha_fin=2026-05-31T23:59:59Z`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "registro_id": 41,
      "alerta_id": 12,
      "camion_id": 10,
      "coordinador_id": 3,
      "mecanico_responsable": "Juan Pérez",
      "fecha_realizada": "2026-05-01T08:00:00Z",
      "kilometraje_mantenimiento": 12450.5,
      "observaciones": "Cambio de aceite y revisión general",
      "created_at": "2026-05-01T08:30:00Z"
    }
  ]
}
```

### 10.5 Reportes de Mantenimiento Generado

| Método | URL | Descripción | Roles autorizados |
|---|---|---|---|
| POST | `/api/reportes-mantenimiento-generado/` | Crear reporte generado | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/reportes-mantenimiento-generado/` | Listar reportes generados | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/reportes-mantenimiento-generado/:id` | Obtener reporte por ID | ADMIN, SUPERVISOR, COORDINADOR |
| PUT | `/api/reportes-mantenimiento-generado/:id` | Actualizar reporte generado | ADMIN, SUPERVISOR, COORDINADOR |
| DELETE | `/api/reportes-mantenimiento-generado/:id` | Eliminar reporte generado | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/reportes-mantenimiento-generado/coordinador/:coordinador_id` | Filtrar por coordinador | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/reportes-mantenimiento-generado/fecha?fecha_inicio=&fecha_fin=` | Filtrar por rango de fechas del reporte | ADMIN, SUPERVISOR, COORDINADOR |
| GET | `/api/reportes-mantenimiento-generado/fecha-generacion?fecha_inicio=&fecha_fin=` | Filtrar por fecha de generación | ADMIN, SUPERVISOR, COORDINADOR |

### 10.5.1 Objeto reporte generado

```json
{
  "reporte_id": 18,
  "coordinador_id": 3,
  "fecha_desde": "2026-05-01T00:00:00Z",
  "fecha_hasta": "2026-05-31T23:59:59Z",
  "observaciones": "Reporte mensual de mantenimiento",
  "created_at": "2026-05-11T08:00:00Z"
}
```

### 10.5.2 Crear reporte generado

**Método**: `POST`  
**URL**: `/api/reportes-mantenimiento-generado/`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "coordinador_id": 3,
  "fecha_desde": "2026-05-01T00:00:00Z",
  "fecha_hasta": "2026-05-31T23:59:59Z",
  "observaciones": "Reporte mensual de mantenimiento"
}
```

#### Response 201 Created

```json
{
  "reporte_id": 18,
  "coordinador_id": 3,
  "fecha_desde": "2026-05-01T00:00:00Z",
  "fecha_hasta": "2026-05-31T23:59:59Z",
  "observaciones": "Reporte mensual de mantenimiento",
  "created_at": "2026-05-11T08:00:00Z"
}
```

### 10.5.3 Listar reportes generados

**Método**: `GET`  
**URL**: `/api/reportes-mantenimiento-generado/`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "reporte_id": 18,
      "coordinador_id": 3,
      "fecha_desde": "2026-05-01T00:00:00Z",
      "fecha_hasta": "2026-05-31T23:59:59Z",
      "observaciones": "Reporte mensual de mantenimiento",
      "created_at": "2026-05-11T08:00:00Z"
    }
  ]
}
```

### 10.5.4 Obtener reporte por ID

**Método**: `GET`  
**URL**: `/api/reportes-mantenimiento-generado/:id`  
**Ejemplo**: `http://localhost:8080/api/reportes-mantenimiento-generado/18`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "reporte_id": 18,
  "coordinador_id": 3,
  "fecha_desde": "2026-05-01T00:00:00Z",
  "fecha_hasta": "2026-05-31T23:59:59Z",
  "observaciones": "Reporte mensual de mantenimiento",
  "created_at": "2026-05-11T08:00:00Z"
}
```

### 10.5.5 Actualizar reporte generado

**Método**: `PUT`  
**URL**: `/api/reportes-mantenimiento-generado/:id`  
**Ejemplo**: `http://localhost:8080/api/reportes-mantenimiento-generado/18`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "coordinador_id": 3,
  "fecha_desde": "2026-05-01T00:00:00Z",
  "fecha_hasta": "2026-06-30T23:59:59Z",
  "observaciones": "Reporte mensual actualizado"
}
```

#### Response 200 OK

```json
{
  "reporte_id": 18,
  "coordinador_id": 3,
  "fecha_desde": "2026-05-01T00:00:00Z",
  "fecha_hasta": "2026-06-30T23:59:59Z",
  "observaciones": "Reporte mensual actualizado",
  "created_at": "2026-05-11T08:00:00Z"
}
```

### 10.5.6 Eliminar reporte generado

**Método**: `DELETE`  
**URL**: `/api/reportes-mantenimiento-generado/:id`  
**Ejemplo**: `http://localhost:8080/api/reportes-mantenimiento-generado/18`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "status": "Reporte de mantenimiento eliminado exitosamente"
}
```

### 10.5.7 Filtrar por coordinador

**Método**: `GET`  
**URL**: `/api/reportes-mantenimiento-generado/coordinador/:coordinador_id`  
**Ejemplo**: `http://localhost:8080/api/reportes-mantenimiento-generado/coordinador/3`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "reporte_id": 18,
      "coordinador_id": 3,
      "fecha_desde": "2026-05-01T00:00:00Z",
      "fecha_hasta": "2026-05-31T23:59:59Z",
      "observaciones": "Reporte mensual de mantenimiento",
      "created_at": "2026-05-11T08:00:00Z"
    }
  ]
}
```

### 10.5.8 Filtrar por rango de fechas

**Método**: `GET`  
**URL**: `/api/reportes-mantenimiento-generado/fecha?fecha_inicio=&fecha_fin=`  
**Ejemplo**: `http://localhost:8080/api/reportes-mantenimiento-generado/fecha?fecha_inicio=2026-05-01T00:00:00Z&fecha_fin=2026-05-31T23:59:59Z`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "reporte_id": 18,
      "coordinador_id": 3,
      "fecha_desde": "2026-05-01T00:00:00Z",
      "fecha_hasta": "2026-05-31T23:59:59Z",
      "observaciones": "Reporte mensual de mantenimiento",
      "created_at": "2026-05-11T08:00:00Z"
    }
  ]
}
```

### 10.5.9 Filtrar por fecha de generación

**Método**: `GET`  
**URL**: `/api/reportes-mantenimiento-generado/fecha-generacion?fecha_inicio=&fecha_fin=`  
**Ejemplo**: `http://localhost:8080/api/reportes-mantenimiento-generado/fecha-generacion?fecha_inicio=2026-05-01T00:00:00Z&fecha_fin=2026-05-31T23:59:59Z`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "reporte_id": 18,
      "coordinador_id": 3,
      "fecha_desde": "2026-05-01T00:00:00Z",
      "fecha_hasta": "2026-05-31T23:59:59Z",
      "observaciones": "Reporte mensual de mantenimiento",
      "created_at": "2026-05-11T08:00:00Z"
    }
  ]
}
```

### Notas comunes

- Todas las rutas protegidas requieren header `Authorization: Bearer <token>`.
- Los parámetros de fecha deben usarse en formato ISO 8601 (`2026-05-01T00:00:00Z`) en `fecha_inicio` y `fecha_fin`.
- En caso de error, la respuesta sigue la estructura estándar definida en la sección "Respuesta de error estandarizada" de este documento.


## 7. Colonias - Operaciones


### 7.1 Resumen de endpoints

| Operación | Método | Endpoint | Autenticación | Rol Requerido |
|-----------|--------|----------|----------------|---------------|
| Listar colonias | GET | `/api/colonia` | No | - |
| Obtener colonia | GET | `/api/colonia/:id` | No | - |
| Crear colonia | POST | `/api/colonia` | JWT requerido | ADMIN |
| Actualizar colonia | PUT | `/api/colonia/:id` | JWT requerido | ADMIN |
| Eliminar colonia | DELETE | `/api/colonia/:id` | JWT requerido | ADMIN |

---

### 7.2 Objeto Colonia

```json
{
  "colonia_id": 1,
  "nombre": "Centro Histórico",
  "descripcion": "Zona central de la ciudad",
  "ciudad": "Ciudad de México",
  "estado": "CDMX",
  "codigo_postal": "06500",
  "latitud": 19.4326,
  "longitud": -99.1332,
  "activa": true,
  "created_at": "2026-01-15T10:00:00Z",
  "updated_at": "2026-05-01T14:30:00Z"
}
```

---

### 7.3 Listar colonias (Público)

**Método**: `GET`  
**URL**: `/api/colonia`  
**Autenticación**: No requerida  

#### Response 200 OK

```json
[
  {
    "colonia_id": 1,
    "nombre": "Centro Histórico",
    "descripcion": "Zona central de la ciudad",
    "ciudad": "Ciudad de México",
    "estado": "CDMX",
    "codigo_postal": "06500",
    "latitud": 19.4326,
    "longitud": -99.1332,
    "activa": true,
    "created_at": "2026-01-15T10:00:00Z",
    "updated_at": "2026-05-01T14:30:00Z"
  },
  {
    "colonia_id": 2,
    "nombre": "Polanco",
    "descripcion": "Zona residencial de lujo",
    "ciudad": "Ciudad de México",
    "estado": "CDMX",
    "codigo_postal": "11560",
    "latitud": 19.4408,
    "longitud": -99.1899,
    "activa": true,
    "created_at": "2026-01-15T10:00:00Z",
    "updated_at": "2026-05-01T14:30:00Z"
  }
]
```

---

### 7.4 Obtener colonia por ID (Público)

**Método**: `GET`  
**URL**: `/api/colonia/:id`  
**Ejemplo**: `http://localhost:8080/api/colonia/1`  
**Autenticación**: No requerida  

#### Response 200 OK

```json
{
  "colonia_id": 1,
  "nombre": "Centro Histórico",
  "descripcion": "Zona central de la ciudad",
  "ciudad": "Ciudad de México",
  "estado": "CDMX",
  "codigo_postal": "06500",
  "latitud": 19.4326,
  "longitud": -99.1332,
  "activa": true,
  "created_at": "2026-01-15T10:00:00Z",
  "updated_at": "2026-05-01T14:30:00Z"
}
```

---

### 7.5 Crear colonia (Administrativo)

**Método**: `POST`  
**URL**: `/api/colonia`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con rol ADMIN)  

#### Request

```json
{
  "nombre": "Nueva Colonia",
  "descripcion": "Descripción de la nueva colonia",
  "ciudad": "Ciudad de México",
  "estado": "CDMX",
  "codigo_postal": "12345",
  "latitud": 19.4500,
  "longitud": -99.1800,
  "activa": true
}
```

**Validaciones:**
- `nombre` requerido (no vacío)
- `ciudad` requerido (no vacío)
- `estado` requerido (no vacío)
- Solo administradores (rol ADMIN) pueden crear colonias

#### Response 201 Created

```json
{
  "colonia_id": 3,
  "nombre": "Nueva Colonia",
  "descripcion": "Descripción de la nueva colonia",
  "ciudad": "Ciudad de México",
  "estado": "CDMX",
  "codigo_postal": "12345",
  "latitud": 19.4500,
  "longitud": -99.1800,
  "activa": true,
  "created_at": "2026-05-04T10:15:00Z",
  "updated_at": "2026-05-04T10:15:00Z"
}
```

---

### 7.6 Actualizar colonia (Administrativo)

**Método**: `PUT`  
**URL**: `/api/colonia/:id`  
**Ejemplo**: `http://localhost:8080/api/colonia/1`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con rol ADMIN)  

#### Request

```json
{
  "nombre": "Centro Histórico Actualizado",
  "descripcion": "Zona central de la ciudad - Actualizada",
  "latitud": 19.4330,
  "activa": true
}
```

**Notas:**
- Solo administradores pueden actualizar colonias
- Todos los campos son opcionales

#### Response 200 OK

Sin contenido (solo status 200)

---

### 7.7 Eliminar colonia (Administrativo)

**Método**: `DELETE`  
**URL**: `/api/colonia/:id`  
**Ejemplo**: `http://localhost:8080/api/colonia/1`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con rol ADMIN)  

#### Response 204 No Content

Sin contenido (solo status 204)

**Notas importantes:**
- Solo administradores pueden eliminar colonias
- La eliminación es permanente
- Los domicilios asociados a esta colonia permanecerán en la base de datos

---

---

---

---

## 9. Incidencias - Operaciones

### 9.1 Resumen de endpoints

| Operación | Método | Endpoint | Autenticación | Roles Requeridos |
|-----------|--------|----------|----------------|------------------|
| Crear incidencia | POST | `/api/incidencias/` | JWT requerido | ADMIN, SUPERVISOR, COORDINADOR |
| Listar incidencias | GET | `/api/incidencias/` | JWT requerido | ADMIN, SUPERVISOR, COORDINADOR |
| Obtener incidencia | GET | `/api/incidencias/:id` | JWT requerido | ADMIN, SUPERVISOR, COORDINADOR |
| Actualizar incidencia | PUT | `/api/incidencias/:id` | JWT requerido | ADMIN, SUPERVISOR, COORDINADOR |
| Eliminar incidencia | DELETE | `/api/incidencias/:id` | JWT requerido | ADMIN, SUPERVISOR, COORDINADOR |
| Por conductor | GET | `/api/incidencias/conductor/:conductor_id` | JWT requerido | ADMIN, SUPERVISOR, COORDINADOR |
| Por punto | GET | `/api/incidencias/punto/:punto_recoleccion_id` | JWT requerido | ADMIN, SUPERVISOR, COORDINADOR |
| Por fecha | GET | `/api/incidencias/fecha?fecha_inicio=&fecha_fin=` | JWT requerido | ADMIN, SUPERVISOR, COORDINADOR |

---

### 9.2 Objeto Incidencia

```json
{
  "incidencia_id": 21,
  "punto_recoleccion_id": 8,
  "conductor_id": 5,
  "descripcion": "Vehículo detenido por bloqueo vial",
  "json_ruta": "[{\"lat\":19.43,\"lng\":-99.13}]",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "eliminado": false,
  "created_at": "2026-04-28T08:35:00Z",
  "updated_at": "2026-04-28T08:35:00Z"
}
```

**Campos:**
- `incidencia_id`: ID único de la incidencia.
- `punto_recoleccion_id`: ID del punto de recolección asociado, opcional.
- `conductor_id`: ID del conductor asociado.
- `descripcion`: Descripción de la incidencia.
- `json_ruta`: Información adicional de ruta en formato JSON, opcional.
- `fecha_reporte`: Fecha y hora del reporte.
- `eliminado`: Indica si la incidencia está marcada como eliminada lógicamente.
- `created_at` / `updated_at`: Fechas de auditoría.

---

### 9.3 Crear incidencia

**Método**: `POST`  
**URL**: `/api/incidencias/`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "punto_recoleccion_id": 8,
  "conductor_id": 5,
  "descripcion": "Vehículo detenido por bloqueo vial",
  "json_ruta": "[{\"lat\":19.43,\"lng\":-99.13}]",
  "fecha_reporte": "2026-04-28T08:30:00Z"
}
```

**Validaciones:**
- `conductor_id` requerido (debe ser numérico)
- `descripcion` requerida (no vacía)
- `fecha_reporte` requerida (formato `YYYY-MM-DD` o ISO 8601)
- `punto_recoleccion_id` y `json_ruta` son opcionales

#### Response 201 Created

```json
{
  "incidencia_id": 21,
  "punto_recoleccion_id": 8,
  "conductor_id": 5,
  "descripcion": "Vehículo detenido por bloqueo vial",
  "json_ruta": "[{\"lat\":19.43,\"lng\":-99.13}]",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "eliminado": false,
  "created_at": "2026-04-28T08:35:00Z",
  "updated_at": "2026-04-28T08:35:00Z"
}
```

---

### 9.4 Listar incidencias

**Método**: `GET`  
**URL**: `/api/incidencias/`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "incidencia_id": 21,
      "punto_recoleccion_id": 8,
      "conductor_id": 5,
      "descripcion": "Vehículo detenido por bloqueo vial",
      "json_ruta": "[{\"lat\":19.43,\"lng\":-99.13}]",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "eliminado": false,
      "created_at": "2026-04-28T08:35:00Z",
      "updated_at": "2026-04-28T08:35:00Z"
    }
  ]
}
```

---

### 9.5 Obtener incidencia por ID

**Método**: `GET`  
**URL**: `/api/incidencias/:id`  
**Ejemplo**: `http://localhost:8080/api/incidencias/21`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "incidencia_id": 21,
  "punto_recoleccion_id": 8,
  "conductor_id": 5,
  "descripcion": "Vehículo detenido por bloqueo vial",
  "json_ruta": "[{\"lat\":19.43,\"lng\":-99.13}]",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "eliminado": false,
  "created_at": "2026-04-28T08:35:00Z",
  "updated_at": "2026-04-28T08:35:00Z"
}
```

---

### 9.6 Actualizar incidencia

**Método**: `PUT`  
**URL**: `/api/incidencias/:id`  
**Ejemplo**: `http://localhost:8080/api/incidencias/21`  
**Headers**: `Authorization: Bearer <token>`, `Content-Type: application/json`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Request

```json
{
  "punto_recoleccion_id": 8,
  "conductor_id": 5,
  "descripcion": "Vehículo liberado luego del bloqueo vial",
  "json_ruta": "[{\"lat\":19.44,\"lng\":-99.12}]",
  "fecha_reporte": "2026-04-28T09:00:00Z"
}
```

**Notas:**
- Los campos se actualizan con los valores enviados.
- `fecha_reporte` es opcional en actualización.

#### Response 200 OK

```json
{
  "incidencia_id": 21,
  "punto_recoleccion_id": 8,
  "conductor_id": 5,
  "descripcion": "Vehículo liberado luego del bloqueo vial",
  "json_ruta": "[{\"lat\":19.44,\"lng\":-99.12}]",
  "fecha_reporte": "2026-04-28T09:00:00Z",
  "eliminado": false,
  "created_at": "2026-04-28T08:35:00Z",
  "updated_at": "2026-04-28T09:00:00Z"
}
```

---

### 9.7 Eliminar incidencia

**Método**: `DELETE`  
**URL**: `/api/incidencias/:id`  
**Ejemplo**: `http://localhost:8080/api/incidencias/21`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "status": "Incidencia marcada como eliminada exitosamente"
}
```

**Notas:**
- La eliminación es lógica: la incidencia se marca como `eliminado = true`.

---

### 9.8 Filtrar por conductor

**Método**: `GET`  
**URL**: `/api/incidencias/conductor/:conductor_id`  
**Ejemplo**: `http://localhost:8080/api/incidencias/conductor/5`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "incidencia_id": 21,
      "punto_recoleccion_id": 8,
      "conductor_id": 5,
      "descripcion": "Vehículo detenido por bloqueo vial",
      "json_ruta": "[{\"lat\":19.43,\"lng\":-99.13}]",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "eliminado": false,
      "created_at": "2026-04-28T08:35:00Z",
      "updated_at": "2026-04-28T08:35:00Z"
    }
  ]
}
```

---

### 9.9 Filtrar por punto de recolección

**Método**: `GET`  
**URL**: `/api/incidencias/punto/:punto_recoleccion_id`  
**Ejemplo**: `http://localhost:8080/api/incidencias/punto/8`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

#### Response 200 OK

```json
{
  "data": [
    {
      "incidencia_id": 21,
      "punto_recoleccion_id": 8,
      "conductor_id": 5,
      "descripcion": "Vehículo detenido por bloqueo vial",
      "json_ruta": "[{\"lat\":19.43,\"lng\":-99.13}]",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "eliminado": false,
      "created_at": "2026-04-28T08:35:00Z",
      "updated_at": "2026-04-28T08:35:00Z"
    }
  ]
}
```

---

### 9.10 Filtrar por fecha

**Método**: `GET`  
**URL**: `/api/incidencias/fecha?fecha_inicio=2026-04-01&fecha_fin=2026-04-30`  
**Headers**: `Authorization: Bearer <token>`  
**Autenticación**: Sí (JWT requerido con roles ADMIN, SUPERVISOR o COORDINADOR)  

**Parámetros de query:**
- `fecha_inicio` (string): Fecha inicial del rango.
- `fecha_fin` (string): Fecha final del rango.

#### Response 200 OK

```json
{
  "data": [
    {
      "incidencia_id": 21,
      "punto_recoleccion_id": 8,
      "conductor_id": 5,
      "descripcion": "Vehículo detenido por bloqueo vial",
      "json_ruta": "[{\"lat\":19.43,\"lng\":-99.13}]",
      "fecha_reporte": "2026-04-28T08:30:00Z",
      "eliminado": false,
      "created_at": "2026-04-28T08:35:00Z",
      "updated_at": "2026-04-28T08:35:00Z"
    }
  ]
}
```

**Notas para frontend**

- Todas las rutas de incidencias requieren token JWT y roles `ADMIN`, `SUPERVISOR`, `COORDINADOR` O `CONDUCTOR`.
- Si un usuario intenta acceder sin token o con un rol no autorizado, recibirá un error `401 Unauthorized` o `403 Forbidden`.
- El token se envía como `Authorization: Bearer <token>` en los headers.

---