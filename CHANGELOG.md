# Changelog

> Todos los cambios importantes en este proyecto están documentados aquí.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/) y el proyecto sigue [Versionado Semántico](https://semver.org/lang/es/).

Referencias de como usarlo: [Guia del Changelog](./CHANGELOG.md#-guía-del-changelog)

---

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