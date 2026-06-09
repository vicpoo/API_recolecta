# API Recolecta — Entidades para Frontend

Este documento reúne solo las entidades que por ahora se integrarán en frontend: login, anomalías y alertas de mantenimiento.

## Base URL

- Local: `http://localhost:8080`
- Prefijo base usado en la API: `/api`

## Autenticación

- Esquema principal: `Authorization: Bearer <token>`
- Varias rutas usan `JWTAuthMiddleware()`.

## Convenciones generales

- `Content-Type: application/json` para `POST`, `PUT`, `PATCH`.
- Fechas en ISO 8601.

### Respuesta de error estandarizada

Todas las respuestas de error usan la misma estructura JSON devuelta por las utilidades en `src/core/error_handler.go`:

```json
{
  "error": {
    "code": "ERROR_CODE",
    "message": "Mensaje legible para frontend",
    "details": { /* opcional, libre */ }
  }
}
```

Campos importantes:
- `code`: código nominal (string) que permite mapear lógica en frontend.
- `message`: texto legible para mostrar al usuario o para logging.
- `details`: objeto opcional con información contextual (por ejemplo `identifier` o `error` con el mensaje interno).

### Códigos de error disponibles (constantes)

| Código | HTTP | Descripción |
|---|---:|---|
| `VALIDATION_ERROR` | 400 | Validación de entrada fallida (campos/binding) |
| `INVALID_INPUT` | 400 | Parámetro de entrada inválido (p. ej. id no numérico) |
| `BAD_REQUEST` | 400 | Solicitud incorrecta |
| `NOT_FOUND` | 404 | Recurso no encontrado (detalle: `identifier`) |
| `UNAUTHORIZED` | 401 | Autenticación requerida o inválida |
| `FORBIDDEN` | 403 | Acceso denegado (rol insuficiente) |
| `CONFLICT` | 409 | Conflicto con datos existentes |
| `INTERNAL_SERVER_ERROR` | 500 | Error interno del servidor (detalle: `error`) |
| `DATABASE_ERROR` | 500 | Error de base de datos (detalle: `error`) |
| `OPERATION_FAILED` | 500 | Operación fallida |

### Mapeo de helpers (comportamiento)

- `RespondValidationError(...)` → HTTP 400, `VALIDATION_ERROR`, `details` contiene el mapa de errores de validación.
- `RespondInvalidInput(...)` → HTTP 400, `INVALID_INPUT`.
- `RespondBadRequest(...)` → HTTP 400, `BAD_REQUEST`.
- `RespondNotFound(resource, id)` → HTTP 404, `NOT_FOUND`, `details: { "identifier": id }`.
- `RespondConflict(...)` → HTTP 409, `CONFLICT`.
- `RespondInternalServerError(..., err)` → HTTP 500, `INTERNAL_SERVER_ERROR`, `details: { "error": err.Error() }` cuando aplica.
- `RespondDatabaseError(..., err)` → HTTP 500, `DATABASE_ERROR`, `details: { "error": err.Error() }`.

### Ejemplos de respuestas de error

- Not found (recurso no existe):

```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Anomalía no encontrado",
    "details": { "identifier": "123" }
  }
}
```

- Error interno con detalle de excepción:

```json
{
  "error": {
    "code": "INTERNAL_SERVER_ERROR",
    "message": "Error al obtener las anomalías",
    "details": { "error": "sql: no rows in result set" }
  }
}
```

---

## 1. Login

### 1.1 Login de empleado

- Método: POST
- URL: `/api/empleados/login`
- Ejemplo completo: `http://localhost:8080/api/empleados/login`
- Headers: `Content-Type: application/json`
- Auth: No

**Request**

```json
{
  "mail_or_username": "admin@recolecta.com",
  "password": "123456"
}
```

**Response 200**

```json
{
  "message": "login correcto",
  "token": "jwt...",
  "data": {
    "id": 1,
    "nombre": "Luis",
    "apellidos": "Pérez",
    "mail": "admin@recolecta.com",
    "username": "luis.perez",
    "desactivado": false,
    "rol_id": 2,
    "created_at": "2026-05-01T12:00:00Z",
    "updated_at": "2026-05-01T12:00:00Z"
  }
}
```

**Notas para frontend**

- Guardar el `token` y enviarlo como `Authorization: Bearer <token>`.
- El `rol_id` viaja dentro del JWT y permite decidir qué vistas mostrar.
- El campo `data.password` no se expone en la respuesta.

---

## 2. Empleados

Base: `/api/empleados`

| Método | URL | Descripción | Auth |
|---|---|---|---|
| POST | `/api/empleados` | Crear empleado | Sí, `ADMIN` |
| GET | `/api/empleados` | Listar empleados | Sí, `ADMIN` |
| GET | `/api/empleados/:id` | Obtener empleado por id | Sí, `ADMIN` |
| PATCH | `/api/empleados/:id` | Actualizar empleado | Sí, `ADMIN` |
| DELETE | `/api/empleados/:id` | Eliminar empleado | Sí, `ADMIN` |
| POST | `/api/empleados/login` | Login de empleado | No |

**Objeto de empleado**

```json
{
  "id": 1,
  "nombre": "Luis",
  "apellidos": "Pérez",
  "mail": "admin@recolecta.com",
  "username": "luis.perez",
  "desactivado": false,
  "rol_id": 2,
  "created_at": "2026-05-01T12:00:00Z",
  "updated_at": "2026-05-01T12:00:00Z"
}
```

### 2.1 Crear empleado

- Método: POST
- URL: `/api/empleados`
- Headers: `Authorization: Bearer <token>`, `Content-Type: application/json`
- Auth: `ADMIN`

**Request**

```json
{
  "nombre": "Luis",
  "apellidos": "Pérez",
  "mail": "admin@recolecta.com",
  "username": "luis.perez",
  "password": "123456",
  "rol_id": 2,
  "desactivado": false
}
```

**Response 201**

```json
{
  "message": "empleado creado correctamente",
  "data": {
    "id": 1,
    "nombre": "Luis",
    "apellidos": "Pérez",
    "mail": "admin@recolecta.com",
    "username": "luis.perez",
    "desactivado": false,
    "rol_id": 2,
    "created_at": "2026-05-01T12:00:00Z",
    "updated_at": "2026-05-01T12:00:00Z"
  }
}
```

### 2.2 Listar empleados

- Método: GET
- URL: `/api/empleados`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`

**Response 200**

```json
{
  "data": [
    {
      "id": 1,
      "nombre": "Luis",
      "apellidos": "Pérez",
      "mail": "admin@recolecta.com",
      "username": "luis.perez",
      "desactivado": false,
      "rol_id": 2,
      "created_at": "2026-05-01T12:00:00Z",
      "updated_at": "2026-05-01T12:00:00Z"
    }
  ]
}
```

### 2.3 Obtener empleado por id

- Método: GET
- URL: `/api/empleados/:id`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`

**Response 200**

```json
{
  "data": {
    "id": 1,
    "nombre": "Luis",
    "apellidos": "Pérez",
    "mail": "admin@recolecta.com",
    "username": "luis.perez",
    "desactivado": false,
    "rol_id": 2,
    "created_at": "2026-05-01T12:00:00Z",
    "updated_at": "2026-05-01T12:00:00Z"
  }
}
```

### 2.4 Actualizar empleado

- Método: PATCH
- URL: `/api/empleados/:id`
- Headers: `Authorization: Bearer <token>`, `Content-Type: application/json`
- Auth: `ADMIN`

**Request**

```json
{
  "nombre": "Luis Alberto",
  "mail": "luis.alberto@recolecta.com",
  "password": "nuevo123",
  "rol_id": 3,
  "desactivado": false
}
```

**Response 200**

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
    "rol_id": 3,
    "created_at": "2026-05-01T12:00:00Z",
    "updated_at": "2026-05-01T12:05:00Z"
  }
}
```

### 2.5 Eliminar empleado

- Método: DELETE
- URL: `/api/empleados/:id`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`

**Response 200**

```json
{ "message": "empleado eliminado correctamente" }
```

**Notas para frontend**

- El borrado es lógico: el registro se marca con `deleted_at` y se oculta de listados y búsquedas.
- Crear, listar, consultar, actualizar y eliminar requieren token JWT y rol `ADMIN`.
- Login sigue disponible sin autenticación en `/api/empleados/login`.

---

## 3. Anomalías

Base: `/api/anomalias`

| Método | URL | Descripción | Auth |
|---|---|---|---|
| POST | `/api/anomalias/` | Crear anomalía | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/anomalias/` | Listar | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/anomalias/:id` | Obtener | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| PUT | `/api/anomalias/:id` | Actualizar | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| DELETE | `/api/anomalias/:id` | Eliminar | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/anomalias/punto/:puntoId` | Por punto | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/anomalias/chofer/:choferId` | Por chofer | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/anomalias/estado` | Filtro por `?estado=` | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/anomalias/tipo` | Filtro por `?tipo_anomalia=` | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/anomalias/por-fecha` | Filtro por `?fecha_inicio=&fecha_fin=` | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |

**Objeto de anomalía**

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

### 3.1 Crear anomalía

- Método: POST
- URL: `/api/anomalias/`
- Headers: `Authorization: Bearer <token>`, `Content-Type: application/json`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`

**Request**

```json
{ "punto_id": 8, "tipo_anomalia": "bloqueo", "descripcion": "Calle cerrada por obras", "fecha_reporte": "2026-04-28T08:30:00Z", "estado": "PENDIENTE", "id_chofer_id": 5 }
```

**Response 201**

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

### 3.2 Listar anomalías

- Método: GET
- URL: `/api/anomalias/`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

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

### 3.3 Obtener anomalía por id

- Método: GET
- URL: `/api/anomalias/:id`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

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

### 3.4 Actualizar anomalía

- Método: PUT
- URL: `/api/anomalias/:id`
- Headers: `Authorization: Bearer <token>`, `Content-Type: application/json`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`

**Request**

```json
{ "tipo_anomalia": "bloqueo", "descripcion": "Calle despejada", "estado": "RESUELTA", "id_chofer_id": 5 }
```

**Response 200**

```json
{
  "anomalia_id": 21,
  "punto_id": 8,
  "tipo_anomalia": "bloqueo",
  "descripcion": "Calle despejada",
  "fecha_reporte": "2026-04-28T08:30:00Z",
  "estado": "RESUELTA",
  "fecha_resolucion": "2026-04-28T09:00:00Z",
  "id_chofer_id": 5
}
```

### 3.5 Eliminar anomalía

- Método: DELETE
- URL: `/api/anomalias/:id`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{ "message": "Anomalía eliminada exitosamente" }
```

### 3.6 Filtrar por punto

- Método: GET
- URL: `/api/anomalias/punto/:puntoId`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

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

### 3.7 Filtrar por chofer

- Método: GET
- URL: `/api/anomalias/chofer/:choferId`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

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

### 3.8 Filtrar por estado

- Método: GET
- URL: `/api/anomalias/estado?estado=PENDIENTE`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

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

### 3.9 Filtrar por tipo

- Método: GET
- URL: `/api/anomalias/tipo?tipo_anomalia=bloqueo`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

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

### 3.10 Filtrar por fecha

- Método: GET
- URL: `/api/anomalias/por-fecha?fecha_inicio=2026-04-01&fecha_fin=2026-04-30`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

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

**Notas para frontend**

- Todas las rutas de anomalías requieren token JWT y roles `ADMIN`, `SUPERVISOR` o `COORDINADOR`.
- Si un usuario intenta acceder sin token o con un rol no autorizado, recibirá un error `401 Unauthorized` o `403 Forbidden`.
- El token se envía como `Authorization: Bearer <token>` en los headers.

---

## 4. Alertas de Mantenimiento

Base: `/api/alertas-mantenimiento`

| Método | URL | Descripción | Auth |
|---|---|---|---|
| POST | `/api/alertas-mantenimiento/` | Crear alerta | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/alertas-mantenimiento/` | Listar | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/alertas-mantenimiento/:id` | Obtener | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| PUT | `/api/alertas-mantenimiento/:id` | Actualizar | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| DELETE | `/api/alertas-mantenimiento/:id` | Eliminar | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/alertas-mantenimiento/pendientes` | Pendientes | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/alertas-mantenimiento/atendidas` | Atendidas | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/alertas-mantenimiento/camion/:camion_id` | Por camión | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/alertas-mantenimiento/tipo/:tipo_id` | Por tipo | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| GET | `/api/alertas-mantenimiento/fecha` | Por fecha | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |
| PATCH | `/api/alertas-mantenimiento/:id/atender` | Marcar atendida | Sí, `ADMIN`, `SUPERVISOR`, `COORDINADOR` |

**Objeto de alerta de mantenimiento**

```json
{
  "alerta_id": 12,
  "camion_id": 10,
  "tipo_mantenimiento_id": 1,
  "descripcion": "Mantenimiento pendiente",
  "observaciones": "Revisar frenos",
  "created_at": "2026-05-01T08:00:00Z",
  "atendido": false
}
```

### 4.1 Crear alerta de mantenimiento

- Método: POST
- URL: `/api/alertas-mantenimiento/`
- Headers: `Authorization: Bearer <token>`, `Content-Type: application/json`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`

**Request**

```json
{ "camion_id": 10, "tipo_mantenimiento_id": 1, "descripcion": "Mantenimiento pendiente", "observaciones": "Revisar frenos" }
```

**Response 201**

```json
{
  "alerta_id": 12,
  "camion_id": 10,
  "tipo_mantenimiento_id": 1,
  "descripcion": "Mantenimiento pendiente",
  "observaciones": "Revisar frenos",
  "created_at": "2026-05-01T08:00:00Z",
  "atendido": false
}
```

### 4.2 Listar alertas de mantenimiento

- Método: GET
- URL: `/api/alertas-mantenimiento/`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{
  "data": [
    {
      "alerta_id": 12,
      "camion_id": 10,
      "tipo_mantenimiento_id": 1,
      "descripcion": "Mantenimiento pendiente",
      "observaciones": "Revisar frenos",
      "created_at": "2026-05-01T08:00:00Z",
      "atendido": false
    }
  ]
}
```

### 4.3 Obtener alerta por id

- Método: GET
- URL: `/api/alertas-mantenimiento/:id`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{
  "alerta_id": 12,
  "camion_id": 10,
  "tipo_mantenimiento_id": 1,
  "descripcion": "Mantenimiento pendiente",
  "observaciones": "Revisar frenos",
  "created_at": "2026-05-01T08:00:00Z",
  "atendido": false
}
```

### 4.4 Actualizar alerta de mantenimiento

- Método: PUT
- URL: `/api/alertas-mantenimiento/:id`
- Headers: `Authorization: Bearer <token>`, `Content-Type: application/json`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`

**Request**

```json
{ "camion_id": 10, "tipo_mantenimiento_id": 1, "descripcion": "Mantenimiento actualizado", "observaciones": "Revisar frenos" }
```

**Response 200**

```json
{
  "alerta_id": 12,
  "camion_id": 10,
  "tipo_mantenimiento_id": 1,
  "descripcion": "Mantenimiento actualizado",
  "observaciones": "Revisar frenos",
  "created_at": "2026-05-01T08:00:00Z",
  "atendido": false
}
```

### 4.5 Eliminar alerta de mantenimiento

- Método: DELETE
- URL: `/api/alertas-mantenimiento/:id`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{ "message": "Alerta eliminada exitosamente" }
```

### 4.6 Alertas pendientes

- Método: GET
- URL: `/api/alertas-mantenimiento/pendientes`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{
  "data": [
    {
      "alerta_id": 12,
      "camion_id": 10,
      "tipo_mantenimiento_id": 1,
      "descripcion": "Mantenimiento pendiente",
      "observaciones": "Revisar frenos",
      "created_at": "2026-05-01T08:00:00Z",
      "atendido": false
    }
  ]
}
```

### 4.7 Alertas atendidas

- Método: GET
- URL: `/api/alertas-mantenimiento/atendidas`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{
  "data": [
    {
      "alerta_id": 13,
      "camion_id": 10,
      "tipo_mantenimiento_id": 2,
      "descripcion": "Mantenimiento atendido",
      "observaciones": "Cambio de filtros",
      "created_at": "2026-05-01T09:00:00Z",
      "atendido": true
    }
  ]
}
```

### 4.8 Filtrar por camión

- Método: GET
- URL: `/api/alertas-mantenimiento/camion/:camion_id`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{
  "data": [
    {
      "alerta_id": 12,
      "camion_id": 10,
      "tipo_mantenimiento_id": 1,
      "descripcion": "Mantenimiento pendiente",
      "observaciones": "Revisar frenos",
      "created_at": "2026-05-01T08:00:00Z",
      "atendido": false
    }
  ]
}
```

### 4.9 Filtrar por tipo

- Método: GET
- URL: `/api/alertas-mantenimiento/tipo/:tipo_id`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{
  "data": [
    {
      "alerta_id": 12,
      "camion_id": 10,
      "tipo_mantenimiento_id": 1,
      "descripcion": "Mantenimiento pendiente",
      "observaciones": "Revisar frenos",
      "created_at": "2026-05-01T08:00:00Z",
      "atendido": false
    }
  ]
}
```

### 4.10 Filtrar por fecha

- Método: GET
- URL: `/api/alertas-mantenimiento/fecha?fecha_inicio=2026-05-01&fecha_fin=2026-05-31`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{
  "data": [
    {
      "alerta_id": 12,
      "camion_id": 10,
      "tipo_mantenimiento_id": 1,
      "descripcion": "Mantenimiento pendiente",
      "observaciones": "Revisar frenos",
      "created_at": "2026-05-01T08:00:00Z",
      "atendido": false
    }
  ]
}
```

### 4.11 Marcar atendida

- Método: PATCH
- URL: `/api/alertas-mantenimiento/:id/atender`
- Headers: `Authorization: Bearer <token>`
- Auth: `ADMIN`, `SUPERVISOR`, `COORDINADOR`
- Response 200:

```json
{ "message": "Alerta marcada como atendida" }
```

**Notas para frontend**

- Todas las rutas de alertas de mantenimiento requieren token JWT y roles `ADMIN`, `SUPERVISOR` o `COORDINADOR`.
- Si un usuario intenta acceder sin token o con un rol no autorizado, recibirá un error `401 Unauthorized` o `403 Forbidden`.
- El token se envía como `Authorization: Bearer <token>` en los headers.