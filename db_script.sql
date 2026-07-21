-- =====================
-- DATABASE
-- =====================
CREATE DATABASE proyecto_recolecta;
\c proyecto_recolecta;

-- =====================
-- VERIFICACIÓN DE VERSIÓN
-- =====================
DO $$
DECLARE
    version_num INTEGER;
BEGIN
    SELECT current_setting('server_version_num')::INTEGER INTO version_num;
    RAISE NOTICE 'PostgreSQL version: %', current_setting('server_version');
    IF version_num < 120000 THEN
        RAISE WARNING 'PostgreSQL version < 12 detectada. Algunas funcionalidades pueden no estar disponibles.';
    END IF;
END $$;

-- =====================
-- DOMINIO EMPLEADO
-- =====================

CREATE TABLE IF NOT EXISTS rol (
    id SMALLINT PRIMARY KEY,
    nombre VARCHAR(50) UNIQUE NOT NULL,
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS empleado (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellidos VARCHAR(100) NOT NULL,
    mail VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL,
    desactivado BOOLEAN DEFAULT FALSE,
    rol_id SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS licencia (
  id SERIAL PRIMARY KEY,
  licencia VARCHAR(100) NOT NULL,
  tipo_licencia SMALLINT NOT NULL,
  fecha_vencimiento DATE NOT NULL,
  id_empleado INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================
-- DOMINIO DISPOSITIVOS
-- =====================

CREATE TABLE IF NOT EXISTS dispositivos (
  id SERIAL PRIMARY KEY,
  conductor_id INTEGER NOT NULL UNIQUE,
  mac_address VARCHAR(100) NOT NULL UNIQUE,
  serial_number VARCHAR(100) NOT NULL UNIQUE,
  api_key VARCHAR(255) NOT NULL UNIQUE,
  nombre_dispositivo VARCHAR(100) NULL,
  sistema_operativo VARCHAR(50) DEFAULT 'Android',
  active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

-- =====================
-- DOMINIO CAMION
-- =====================

-- =====================
-- CONTEXT BOUNDARY ASIGNACION
-- =====================

CREATE TABLE IF NOT EXISTS historial_asignacion_camion (
  id_historial SERIAL PRIMARY KEY,
  id_chofer INTEGER NOT NULL,
  id_camion INTEGER NOT NULL,
  fecha_asignacion DATE NOT NULL,
  fecha_baja DATE DEFAULT NULL,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS tipo_camion (
  tipo_camion_id SERIAL PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL,
  descripcion VARCHAR(255) NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS camion (
  id SERIAL PRIMARY KEY,
  placa VARCHAR(20) NOT NULL,
  modelo VARCHAR(50) NOT NULL,
  rentado BOOLEAN DEFAULT FALSE,
  estado VARCHAR(20) NOT NULL,
  tipo_id INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS tipo_mantenimiento (
  tipo_mantenimiento_id SERIAL PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL,
  categoria VARCHAR(50) NOT NULL,
  eliminado BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS alerta_mantenimiento (
  alerta_id SERIAL PRIMARY KEY,
  camion_id INTEGER NOT NULL,
  tipo_mantenimiento_id INTEGER NOT NULL,
  descripcion TEXT NOT NULL,
  observaciones TEXT NULL,
  atendido BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS registro_mantenimiento (
  registro_id SERIAL PRIMARY KEY,
  alerta_id INTEGER DEFAULT NULL,
  camion_id INTEGER NOT NULL,
  coordinador_id INTEGER NOT NULL,
  mecanico_responsable VARCHAR(100) NOT NULL,
  fecha_realizada TIMESTAMP NOT NULL,
  kilometraje_mantenimiento INTEGER NOT NULL,
  observaciones TEXT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

-- =====================
-- DOMINIO RUTA
-- =====================

-- =====================
-- BOUNDARY ASIGNACION RUTA
-- =====================

CREATE TABLE IF NOT EXISTS ruta_camion (
  ruta_camion_id SERIAL PRIMARY KEY,
  ruta_id INTEGER NOT NULL,
  camion_id INTEGER NOT NULL,
  fecha DATE NOT NULL,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS ruta (
  id SERIAL PRIMARY KEY,
  nombre VARCHAR(100) NOT NULL,
  descripcion VARCHAR(255) NOT NULL,
  colonia_id INTEGER NOT NULL,
  json_ruta JSON NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS punto_recoleccion (
  id SERIAL PRIMARY KEY,
  ruta_id INTEGER NOT NULL,
  direccion VARCHAR(255) NOT NULL,
  orden DOUBLE PRECISION DEFAULT 0.0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS relleno_sanitario (
  relleno_id SERIAL PRIMARY KEY,
  nombre VARCHAR(100) NOT NULL,
  direccion VARCHAR(255) NOT NULL,
  es_rentado BOOLEAN DEFAULT FALSE,
  capacidad_toneladas DOUBLE PRECISION NOT NULL,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS estado_camion (
  estado_id SERIAL PRIMARY KEY,
  camion_id INTEGER NOT NULL,
  estado VARCHAR(50) NOT NULL,
  observaciones TEXT NULL,
  timestamp TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS registro_vaciado (
  vaciado_id SERIAL PRIMARY KEY,
  relleno_id INTEGER NOT NULL,
  ruta_camion_id INTEGER NOT NULL,
  hora TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- =====================
-- DOMINIO COLONIA
-- =====================

CREATE TABLE IF NOT EXISTS colonia (
  colonia_id SERIAL PRIMARY KEY,
  nombre VARCHAR(100) NOT NULL,
  zona VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- =====================
-- DOMINIO CIUDADANO
-- =====================

CREATE TABLE IF NOT EXISTS ciudadano (
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) NOT NULL,
  alias VARCHAR(100) NOT NULL,
  password VARCHAR(100) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS domicilio (
  id SERIAL PRIMARY KEY,
  alias VARCHAR(100) NOT NULL,
  calle VARCHAR(100) NOT NULL,
  numero VARCHAR(20) NOT NULL,
  referencia VARCHAR(255) DEFAULT NULL,
  ciudadano_id INTEGER NOT NULL,
  colonia_id INTEGER NOT NULL,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- =====================
-- DOMINIO NOTIFICACIÓN
-- =====================

CREATE TABLE IF NOT EXISTS alerta_usuario (
  alerta_id SERIAL PRIMARY KEY,
  usuario_id INTEGER NOT NULL,
  titulo VARCHAR(150) NOT NULL,
  mensaje TEXT NOT NULL,
  leida BOOLEAN DEFAULT FALSE NOT NULL,
  creado_por INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS aviso (
  id SERIAL PRIMARY KEY,
  enviado_por INTEGER NOT NULL,
  tipo_aviso VARCHAR(50) NOT NULL,
  descripcion VARCHAR(255) NOT NULL,
  entidad_involucrada VARCHAR(100) NOT NULL,
  estado SMALLINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL
);

-- =====================
-- DOMINIO FALLAS (ANOMALIA)
-- =====================

-- Tabla unificada que reemplaza a Anomalia, Incidencia, ReporteConductor,
-- ReporteFallaCritica y SeguimientoFallaCritica. El campo tipo_anomalia
-- indica cuál de esos conceptos representa cada registro.
-- SeguimientoFallaCritica usa anomalia_referencia_id (auto-relación) para
-- apuntar al anomalia_id del REPORTE_FALLA_CRITICA al que da seguimiento.
CREATE TABLE IF NOT EXISTS anomalia (
  anomalia_id SERIAL PRIMARY KEY,
  tipo_anomalia VARCHAR(50) NOT NULL,
  punto_id INTEGER DEFAULT NULL,
  conductor_id INTEGER DEFAULT NULL,
  camion_id INTEGER DEFAULT NULL,
  ruta_id INTEGER DEFAULT NULL,
  anomalia_referencia_id INTEGER DEFAULT NULL,
  descripcion TEXT NOT NULL,
  json_ruta TEXT DEFAULT NULL,
  estado VARCHAR(30) DEFAULT NULL,
  eliminado BOOLEAN DEFAULT FALSE,
  fecha_reporte TIMESTAMP NOT NULL,
  fecha_resolucion TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================
-- FUNCIONES Y TRIGGERS (Para updated_at automático en Postgres)
-- =====================
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT c.table_name
        FROM information_schema.columns c
        JOIN information_schema.tables t ON c.table_name = t.table_name AND c.table_schema = t.table_schema
        WHERE c.column_name = 'updated_at' AND c.table_schema = 'public' AND t.table_type = 'BASE TABLE'
    LOOP
        EXECUTE format('DROP TRIGGER IF EXISTS update_timestamp ON %I', t);
        EXECUTE format('CREATE TRIGGER update_timestamp BEFORE UPDATE ON %I FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column()', t);
    END LOOP;
END $$;

-- =====================
-- MENSAJE DE FINALIZACIÓN
-- =====================
DO $$
DECLARE
    table_count INTEGER;
BEGIN
    SELECT COUNT(*) INTO table_count
    FROM information_schema.tables
    WHERE table_schema = 'public' AND table_type = 'BASE TABLE';

    RAISE NOTICE '=========================================';
    RAISE NOTICE '✅ Script de inicialización de schema completado exitosamente';
    RAISE NOTICE 'Base de datos: proyecto_recolecta';
    RAISE NOTICE 'Tablas creadas/verificadas: %', table_count;
    RAISE NOTICE '=========================================';
END $$;
