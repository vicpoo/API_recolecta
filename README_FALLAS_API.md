# API Recolecta — Flujo de Fallas (Anomalías)

Este documento describe el nuevo flujo del módulo **Fallas**, que reemplaza las
cinco entidades anteriores (`Anomalia`, `Incidencia`, `ReporteConductor`,
`ReporteFallaCritica`, `SeguimientoFallaCritica`) por una **única entidad
unificada `Anomalia`**, diferenciada mediante el campo `tipo_anomalia`.

> ⚠️ Ver la sección [Estado actual / pendientes](#estado-actual--pendientes)
> al final: el módulo está en migración y **el proyecto no compila
> todavía**. Este documento describe el contrato objetivo (el que ya
> respetan el dominio, los casos de uso y el repositorio) y señala
> explícitamente qué controladores faltan por alinear.

## Base URL

- Local: `http://localhost:8080`
- Prefijo: `/api/anomalias`

## Autenticación y roles

Todas las rutas de `/api/anomalias` requieren JWT y uno de estos roles:

- `ADMIN`
- `SUPERVISOR`
- `COORDINADOR`

```
Authorization: Bearer <token>
```

Sin token válido → `401 UNAUTHORIZED`. Con token pero rol no permitido →
`403 FORBIDDEN`.

---

## Por qué se unificó el modelo

Antes existían 5 tablas/entidades independientes con CRUD casi idéntico
(`Incidencia`, `ReporteConductor`, `ReporteFallaCritica`,
`SeguimientoFallaCritica`, y la propia `Anomalia`). Ahora todo vive en una
sola tabla `anomalia`, y el campo `tipo_anomalia` indica qué "concepto"
representa cada fila:

| `tipo_anomalia`             | Reemplaza a                | Notas |
|---|---|---|
| `ANOMALIA`                  | Anomalia (general)          | Caso genérico |
| `INCIDENCIA`                | Incidencia                  | Reportada en un punto de recolección |
| `REPORTE_CONDUCTOR`         | ReporteConductor             | Reportada por un conductor |
| `REPORTE_FALLA_CRITICA`     | ReporteFallaCritica           | Falla grave de camión/ruta |
| `SEGUIMIENTO_FALLA_CRITICA` | SeguimientoFallaCritica        | Comentario/seguimiento sobre un `REPORTE_FALLA_CRITICA` existente |

`SEGUIMIENTO_FALLA_CRITICA` es un caso especial: **no** crea un nuevo
incidente, sino que usa `anomalia_referencia_id` para apuntar al
`anomalia_id` del `REPORTE_FALLA_CRITICA` al que da seguimiento (auto
relación dentro de la misma tabla `anomalia`).

---

## Modelo de datos: `Anomalia`

```json
{
  "anomalia_id": 12,
  "tipo_anomalia": "REPORTE_FALLA_CRITICA",
  "punto_id": 5,
  "conductor_id": 3,
  "camion_id": 2,
  "ruta_id": 7,
  "anomalia_referencia_id": null,
  "descripcion": "Fuga de aceite en el motor",
  "json_ruta": "",
  "estado": "PENDIENTE",
  "eliminado": false,
  "fecha_reporte": "2026-07-14T08:30:00Z",
  "fecha_resolucion": null,
  "created_at": "2026-07-14T08:31:02Z",
  "updated_at": "2026-07-14T08:31:02Z"
}
```

| Campo | Tipo | Notas |
|---|---|---|
| `anomalia_id` | int32 | PK, autogenerado |
| `tipo_anomalia` | string (enum) | Uno de los 5 valores de la tabla anterior |
| `punto_id` | int32 \| null | Punto de recolección relacionado |
| `conductor_id` | int32 \| null | Conductor relacionado |
| `camion_id` | int32 \| null | Camión relacionado |
| `ruta_id` | int32 \| null | Ruta relacionada |
| `anomalia_referencia_id` | int32 \| null | Solo se usa en `SEGUIMIENTO_FALLA_CRITICA`: apunta a otra `anomalia_id` |
| `descripcion` | string | Requerido |
| `json_ruta` | string | Opcional, texto libre (JSON serializado de una ruta/ubicación) |
| `estado` | string | `PENDIENTE`, `EN_PROCESO`, `RESUELTA`, u otro definido por el cliente |
| `eliminado` | bool | Borrado lógico |
| `fecha_reporte` | datetime | Requerido; si se omite al crear, se usa la fecha/hora actual |
| `fecha_resolucion` | datetime \| null | Se llena al resolver |
| `created_at` / `updated_at` | datetime | Gestionados por el servidor |

El **borrado es lógico**: `DELETE` no borra la fila, marca
`eliminado = true`. Esto es intencional porque una fila
`SEGUIMIENTO_FALLA_CRITICA` puede referenciar a otra vía
`anomalia_referencia_id`, y borrar físicamente rompería esa relación.
`GetAll` y todos los `GetBy*` filtran siempre `eliminado = false`.

---

## Respuesta de error estandarizada

Todas las rutas del backend (no solo Fallas) usan el mismo formato de error:

```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Anomalía no encontrado",
    "details": { "identifier": "12" }
  }
}
```

| Código | HTTP | Cuándo ocurre en este módulo |
|---|---:|---|
| `VALIDATION_ERROR` | 400 | Falta un campo requerido en el body, o falta un query param requerido (`estado`, `tipo_anomalia`, `fecha_inicio`/`fecha_fin`) |
| `BAD_REQUEST` | 400 | Formato de fecha inválido (no ISO 8601 / no reconocido) |
| `INVALID_INPUT` | 400 | El `id` de la URL no es un entero válido |
| `NOT_FOUND` | 404 | El `anomalia_id` no existe (ver aviso abajo) |
| `INTERNAL_SERVER_ERROR` | 500 | Error no controlado (incluye, hoy, los casos que deberían ser 404) |

---

## Endpoints

### 1. Crear anomalía

`POST /api/anomalias/`

**Body** (contrato objetivo — `entities.CreateAnomaliaRequest`):

```json
{
  "tipo_anomalia": "REPORTE_FALLA_CRITICA",
  "punto_id": null,
  "conductor_id": 3,
  "camion_id": 2,
  "ruta_id": 7,
  "anomalia_referencia_id": null,
  "descripcion": "Fuga de aceite en el motor",
  "json_ruta": "",
  "estado": "PENDIENTE",
  "fecha_reporte": "2026-07-14T08:30:00Z"
}
```

- Requeridos: `tipo_anomalia`, `descripcion`.
- `fecha_reporte` es opcional; si se omite/está vacía, el servidor la fija
  con la hora actual.
- `tipo_anomalia` debe ser uno de los 5 valores del enum
  (`EsTipoAnomaliaValido` en `tipo_anomalia.go`).

**Respuesta 201** — el objeto `Anomalia` creado (ver modelo de datos arriba).

**Errores:**
- `400 VALIDATION_ERROR` — falta `tipo_anomalia` o `descripcion`.
- `400 BAD_REQUEST` — `fecha_reporte` con formato no reconocido.

---

### 2. Listar todas

`GET /api/anomalias/`

**Respuesta 200:**

```json
{ "data": [ { "anomalia_id": 1, "...": "..." } ] }
```

Solo devuelve filas con `eliminado = false`, ordenadas por
`fecha_reporte DESC`.

---

### 3. Obtener por ID

`GET /api/anomalias/{id}`

**Respuesta 200:** objeto `Anomalia`.
**Errores:** `400 INVALID_INPUT` (id no numérico) · `404 NOT_FOUND` (no existe).

---

### 4. Actualizar

`PUT /api/anomalias/{id}`

**Body** (contrato objetivo — `entities.UpdateAnomaliaRequest`):

```json
{
  "tipo_anomalia": "REPORTE_FALLA_CRITICA",
  "punto_id": null,
  "conductor_id": 3,
  "camion_id": 2,
  "ruta_id": 7,
  "anomalia_referencia_id": null,
  "descripcion": "Fuga de aceite reparada",
  "json_ruta": "",
  "estado": "RESUELTA",
  "eliminado": false,
  "fecha_reporte": "2026-07-14T08:30:00Z",
  "fecha_resolucion": "2026-07-15T10:00:00Z"
}
```

- Requeridos: `tipo_anomalia`, `descripcion`, `fecha_reporte`.
- `fecha_resolucion` es opcional (string ISO 8601 o `null`).

**Respuesta 200:** objeto `Anomalia` actualizado.
**Errores:** `400 INVALID_INPUT` (id no numérico) · `400 VALIDATION_ERROR` /
`400 BAD_REQUEST` (body inválido o fechas mal formateadas) ·
`404 NOT_FOUND` (no existe).

---

### 5. Eliminar (borrado lógico)

`DELETE /api/anomalias/{id}`

**Respuesta 200:**

```json
{ "message": "Anomalía eliminada exitosamente" }
```

**Errores:** `400 INVALID_INPUT` (id no numérico) · `404 NOT_FOUND` (no existe).

---

### 6. Por punto de recolección

`GET /api/anomalias/punto/{puntoId}`

**Respuesta 200:** arreglo de `Anomalia` (sin envolver en `data`).
**Errores:** `400 INVALID_INPUT` (id no numérico).

---

### 7. Por conductor/chofer

`GET /api/anomalias/chofer/{choferId}`

Filtra por `conductor_id`. **Respuesta 200:** arreglo de `Anomalia`.
**Errores:** `400 INVALID_INPUT` (id no numérico).

---

### 8. Por estado

`GET /api/anomalias/estado?estado=PENDIENTE`

**Respuesta 200:** arreglo de `Anomalia`.
**Errores:** `400 VALIDATION_ERROR` (falta `estado`).

---

### 9. Por tipo de anomalía

`GET /api/anomalias/tipo?tipo_anomalia=REPORTE_FALLA_CRITICA`

**Respuesta 200:** arreglo de `Anomalia`.
**Errores:** `400 VALIDATION_ERROR` (falta `tipo_anomalia`).

---

### 10. Por rango de fechas

`GET /api/anomalias/por-fecha?fecha_inicio=2026-07-01&fecha_fin=2026-07-15`

**Respuesta 200:** arreglo de `Anomalia`.
**Errores:** `400 VALIDATION_ERROR` (falta `fecha_inicio` o `fecha_fin`).

---

## Nota sobre formas de respuesta inconsistentes

Hoy la forma de la respuesta 200 **no es uniforme** entre endpoints:

| Endpoint | Forma real de la respuesta |
|---|---|
| Crear / Obtener por ID / Actualizar | objeto `Anomalia` "pelado" |
| Listar todas | `{ "data": [...] }` |
| Eliminar | `{ "message": "..." }` |
| Por punto / chofer / estado / tipo / fecha | arreglo `[...]` "pelado" |

Los comentarios `@Success` de Swagger dicen que todas devuelven
`AnomaliaResponse` / `AnomaliaListResponse` (con `success`, `message`,
`data`, `code`), pero el código real (`core.RespondOK`, `core.RespondCreated`,
`c.JSON(...)`) no arma ese envoltorio. Si el frontend/mobile va a
consumir esto, conviene decidir una única forma y alinear código + swagger.

---

## Estado actual / pendientes

Al momento de escribir este documento, **`go build ./...` falla** porque el
módulo está a medio migrar. Antes de dar por cerrado el flujo nuevo, falta:

1. **Eliminar los archivos legado** de `Incidencia`, `ReporteConductor`,
   `ReporteFallaCritica` y `SeguimientoFallaCritica` en
   `src/Fallas/infrastructure/` (controllers, routers, dependencias,
   repositorios postgres) y sus registros en `dependencies.go` (raíz). Sus
   casos de uso (`src/Fallas/application/*UseCase.go`) y entidades
   (`src/Fallas/domain/entities/*.go`) ya fueron borrados, pero los
   controladores que los usaban siguen presentes y no compilan
   (`undefined: application.CreateIncidenciaUseCase`, etc.).

2. **Arreglar `CreateAnomaliaController.go` y `UpdateAnomaliaController.go`**:
   siguen usando un struct de request con `IDChoferID` y llaman a
   `entities.NewAnomaliaConPunto(...)`, ninguno de los cuales existe ya en
   `entities.Anomalia` (que usa `ConductorID`, no `IDChoferID`). Lo más
   simple es hacerlos `ShouldBindJSON` directo sobre
   `entities.CreateAnomaliaRequest` / `entities.UpdateAnomaliaRequest`
   (definidos en `anomalia_swagger.go`), que ya calzan con los campos reales
   de la entidad.

3. **Renombrar `GetAnomaliasByChoferIDUseCase`** (referenciado en
   `dependencies_anomalia.go` y `GetAnomaliasByChoferIDController.go`) a algo
   consistente con el repositorio, que expone `GetByConductorID` — el
   archivo `GetAnomaliasByChoferIDUseCase.go` no existe actualmente en
   `src/Fallas/application/`.

4. **Corregir la detección de "no encontrado"**: `PostgresAnomaliaRepository`
   devuelve errores como `fmt.Errorf("anomalía con ID %d no encontrada", id)`,
   pero `GetAnomaliaByIdController`, `UpdateAnomaliaController` y
   `DeleteAnomaliaController` comparan contra el string literal
   `"anomalia not found"` (que nunca se produce). Resultado: un ID
   inexistente hoy devolvería `500 INTERNAL_SERVER_ERROR` en vez de
   `404 NOT_FOUND`. Conviene exportar un error tipado
   (p. ej. `var ErrAnomaliaNotFound = errors.New(...)`) y compararlo con
   `errors.Is`.

5. El módulo **Mantenimiento** completo también fue borrado en el mismo
   cambio (`src/Mantenimiento/...` y su registro en `dependencies.go`); si
   eso fue intencional, falta quitar cualquier referencia residual, aunque
   no se encontró ninguna en el resto del código.
