# Changelog

> Todos los cambios importantes en este proyecto están documentados aquí.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/) y el proyecto sigue [Versionado Semántico](https://semver.org/lang/es/).

Referencias de como usarlo: [Guia del Changelog](./CHANGELOG.md#-guía-del-changelog)

---

# [0.10.0-alpha] - 2026-04-15
## Rodrigo Mijangos [Issue #14](https://github.com/RodrigoMijangos/recolecta_web/issues/14)
### Added
- Pruebas unitarias para repositorio Redis de trazabilidad de eventos (`event_deduplication`, `event_trace`).
- Pruebas unitarias para repositorio Redis de sesiones realtime admin (`ws:upgrade`, `ws:session`, `server_epoch`).

### Changed
- Se fortalece validación de regresiones en flujos de dedupe/traza y lifecycle de sesión realtime.

# [0.9.0-alpha] - 2026-04-15
## Rodrigo Mijangos [Issue #12](https://github.com/RodrigoMijangos/recolecta_web/issues/12)
### Added
- Endpoints de consulta de trazas de eventos: `GET /api/notifications/events/traces/:event_id` y `GET /api/notifications/events/traces/truck/:truck_id`.
- Endpoint de consulta de sesión realtime: `GET /api/realtime/ws/sessions/:session_id`.
- Lectura tipada de trazas y sesiones desde Redis para observabilidad operativa.

### Changed
- El router de notificaciones/realtime expone capacidades de inspección para soporte y auditoría.

# [0.8.0-alpha] - 2026-04-15
## Rodrigo Mijangos [Issue #11](https://github.com/RodrigoMijangos/recolecta_web/issues/11)
### Added
- Emisión de token exclusivo de upgrade websocket para administrador (`/api/realtime/ws/upgrade-token`).
- Consumo one-time del token de upgrade y creación de sesión realtime (`/api/realtime/ws/sessions/consume`).
- Endpoints de heartbeat y cierre de sesión realtime para administrador.

### Changed
- Se incorpora repositorio Redis para `realtime:server_epoch:current`, `ws:upgrade:*` y `ws:session:*`.

# [0.7.0-alpha] - 2026-04-15
## Rodrigo Mijangos [Issue #10](https://github.com/RodrigoMijangos/recolecta_web/issues/10)
### Added
- Caso de uso para procesar eventos `TruckStateEvent` con deduplicación por hash.
- Persistencia de deduplicación y trazabilidad de eventos en Redis (`event_deduplication:*`, `event_trace:*`).
- Endpoint `POST /api/notifications/events/truck-state` para orquestar resolución de regla por estado.

### Changed
- El enrutador de notificaciones integra el flujo de orquestación de eventos con reglas dinámicas.

# [0.6.0-alpha] - 2026-04-15
## Rodrigo Mijangos [Issue #9](https://github.com/RodrigoMijangos/recolecta_web/issues/9)
### Added
- Motor dinámico de reglas de notificación en Redis por `state_code`.
- Nuevos endpoints para administrar reglas: listar, consultar por estado, crear/actualizar y eliminar.
- Control de versionado de reglas en Redis con clave global `rules:version`.

### Changed
- El módulo de notificaciones incorpora un repositorio dedicado para `rules:state:{state_code}`.
- Se agrega validación de payload para reglas (estado, acción y radio mínimo).

# [0.5.0-alpha] - 2026-04-15
## Rodrigo Mijangos [Issue #53](https://github.com/RodrigoMijangos/recolecta_web/issues/53)
### Added
- Contratos base de eventos de estado de camión (`TruckStateEvent`) con control de versión (`v1`).
- Contrato base para token exclusivo de upgrade websocket de administrador (`AdminWSUpgradeTokenClaim`).
- Catálogo inicial de estados operativos/críticos y acciones de orquestación para notificaciones.

### Changed
- Se define una frontera clara server-owned para payloads de movilidad y canal realtime de administración.
- Se alinean dependencias del módulo de notificaciones en `go.mod`/`go.sum` para compilación consistente de FCM y Redis.

# [0.4.0-alpha] - 2026-03-26
## Rodrigo Mijangos [Issue #X](https://github.com/RodrigoMijangos/recolecta_web/issues/X)
### Added
- Implementación de entidades para modelado de datos a utilizar en PostgreSQL conforme a ciudadanos.
- Ruta e implementación de lógica para registrar ciudadanos.
- Flexibilidad para registrar ciudadanos con o sin geolocalización, tokens FCM son obligatorios.
- Repositorio de ciudadanos usando PostgreSQL.
- Caso de uso para actualizar coordenadas.
- Creación del controlador de ciudadanos con endpoint para registrar ciudadanos y actualizar coordenadas.

### Changed
- Alta en las dependencias de la infraestructura.
- Configuración de archivo main para incializar Redis y cerrarlo al finalizar la ejecución.

# 0.3.0-alpha - 2026-03-25
## Rodrigo Mijangos [Issue #24](https://github.com/RodrigoMijangos/recolecta_web/issues/24)
### Added
- Implementación de Caso de uso y Endopoint para mandar notificaciones a FCM.
- Implementación de Caso de uso y Endopoint para guardar logs de notificaciones push.
- Logging de mensajes de error en el servicio de notificaciones push.

# 0.2.0-alpha - 2026-03-24
## Rodrigo Mijangos [Issue #23](https://github.com/RodrigoMijangos/recolecta_web/issues/23)
### Added
- Implementación de entidad para logs de notificaciones push.
- Implementación de entidad para envio de notificaciones push.
- Repositorio de logs de notificaciones push usando Redis.
- Servicio de notificaciones push con lógica de negocio para enviar notificaciones y registrar logs.
- Servicio de almacenamiento para logs de notificaiones push.

### Changed
- Corrección de fecha de implementación del changelog en la versión 0.1.0-alpha.

# 0.1.0-alpha - 2026-03-24
## Rodrigo Mijangos [Issue #X](https://github.com/RodrigoMijangos/recolecta_web/issues/X)
### Added
- Archivo de configuración para variables de entorno.
- Entidades de ciudadano para almacenar información de geolocalización y manejo de tokens FCM.
- Puerto de almacenamiento agnóstico al servicio para guardar tokens FCM y geolocalización.
- Repositorio de operaciones para manejar la lógica de negocio relacionada con geolocalización y tokens FCM de ciudadanos.
- Confifuración de cliente Redis en formato de inyección de dependencias.
---

[Volver Arriba](#-changelog)

## 📖 Guía del Changelog

### 🎯 Cómo Leerlo

Cada versión está dividida en **categorías** que te ayudan a identificar qué tipo de cambios se hicieron:

| Categoría | Significa | Ejemplo |
|-----------|-----------|---------|
| **Added** | Nuevas funcionalidades | Nueva página de login |
| **Changed** | Cambios en funcionalidad existente | Refactor de componentes |
| **Deprecated** | Features que pronto desaparecerán | Método antiguo que será reemplazado |
| **Removed** | Código o archivos removidos | Componentes deprecados |
| **Fixed** | Bug fixes | Corrección de error en validación |
| **Security** | Parches de seguridad | Actualización de dependencias críticas |

### 🏗️ Cómo Mantenerlo

Cada vez que hagas cambios importantes, **debes actualizar el changelog** ANTES de hacer el commit:

#### En Desarrollo (rama activa)

```markdown
## [Unreleased]

### Added
- Nueva funcionalidad X

### Fixed
- Bug en el componente Y
```

#### 📝 Guía de Traducción: Commits → Changelog

Usa esta tabla para decidir si un commit debe ir al changelog y cómo categorizarlo:

| Tipo Commit | ¿Va al Changelog? | Categoría | Ejemplo |
|-------------|-------------------|-----------|---------|
| `feat:` | ✅ Sí | **Added** | `feat: agregar notificaciones FCM` → `Added: Sistema de notificaciones FCM` |
| `fix:` | ✅ Sí | **Fixed** | `fix: corregir cálculo de radio` → `Fixed: Cálculo de radio en geolocalización` |
| `perf:` | ✅ Sí | **Changed** | `perf: optimizar consultas Redis` → `Changed: Optimización de consultas geoespaciales` |
| `refactor:` | ⚠️ Solo si es significativo | **Changed** | `refactor: reestructurar módulo rutas` → `Changed: Reestructuración de módulo de rutas` |
| `docs:` | ⚠️ Solo si es importante | **Added/Changed** | `docs: agregar guía Redis` → `Added: Documentación de schema Redis` |
| `chore:` | ❌ No (generalmente) | - | `chore: actualizar deps` → (no va al changelog) |
| `test:` | ❌ No | - | `test: agregar tests unitarios` → (no va al changelog) |
| `style:` | ❌ No | - | `style: formatear código` → (no va al changelog) |
| `build:` | ❌ No | - | `build: actualizar Dockerfile` → (no va al changelog) |
| `ci:` | ❌ No | - | `ci: configurar GitHub Actions` → (no va al changelog) |

**Reglas:**
- Si el cambio **afecta al usuario o desarrollador**, va al changelog
- Si es solo interno/mantenimiento, NO va
- Traduce commits técnicos a lenguaje claro para el changelog

#### ✅ Al Hacer Release

1. **Reemplaza `[Unreleased]` con la versión** en formato `X.Y.Z`
2. **Añade la fecha** en formato `YYYY-MM-DD`
3. **Crea un nuevo tag** en Git

```bash
# Ejemplo:
git tag -a v0.2.0 -m "Release version 0.2.0"
git push origin v0.2.0
```

---

## 📊 Sistema de Versionado (Versionado Semántico)

Usamos **SemVer**: `MAJOR.MINOR.PATCH(-prerelease)(+metadata)`

### Formato: X.Y.Z

```
0.1.0
├── 0 = MAJOR (cambios incompatibles)
├── 1 = MINOR (nuevas funcionalidades)
└── 0 = PATCH (bug fixes)
```

### 📈 Reglas de Versionado

| Cambio | Incrementa | Ejemplo |
|--------|-----------|---------|
| Bug fixes y mejoras pequeñas | PATCH | 0.1.0 → 0.1.1 |
| Nuevas funcionalidades | MINOR | 0.1.0 → 0.2.0 |
| Cambios incompatibles | MAJOR | 0.1.0 → 1.0.0 |

### 🔤 Estados Especiales (Prerelease)

Para versiones en desarrollo, usamos sufijos:

```
0.1.0-alpha    → Versión muy temprana, inestable
0.1.0-beta     → Más estable pero en pruebas
0.1.0-rc.1     → Release Candidate (casi lista)
1.0.0          → Versión estable final
```

### 📋 Hoja de Referencia Rápida

```bash
# Versión actual
git describe --tags

# Ver todos los tags
git tag -l

# Crear nuevo tag (cuando hagas release)
git tag -a v0.2.0 -m "Release version 0.2.0"

# Ver cambios desde último tag
git log $(git describe --tags --abbrev=0)..HEAD --oneline
```

---

## 💡 Consejos para Desarrolladores

### ✍️ Al Hacer Cambios

1. **Trabaja en tu rama** (ej: `feature/nueva-funcionalidad`)
2. **Actualiza el changelog** en la sección `[Unreleased]`
3. **Sé descriptivo** pero conciso:
   - ✅ `Added: Modal de confirmación en validación de rutas`
   - ❌ `fixed stuff`

### 🔍 Antes de hacer un Pull Request

```bash
# Verifica que el changelog esté actualizado
git diff main -- CHANGELOG.md

# Lee tu changelog
cat CHANGELOG.md
```

### 📦 Al Hacer Release (Solo para Admin)

```bash
# 1. Actualizar versión en package.json (frontend)
# 2. Reemplazar [Unreleased] en CHANGELOG.md
# 3. Hacer commit
git commit -am "chore: release v0.2.0"

# 4. Crear tag
git tag -a v0.2.0 -m "Release version 0.2.0"

# 5. Hacer push
git push origin main
git push origin v0.2.0
```
