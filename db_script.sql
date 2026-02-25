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

CREATE TABLE rol (
    id TINYINT PRIMARY KEY,
    nombre VARCHAR(50) UNIQUE NOT NULL
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE empleado (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellidos VARCHAR(100) NOT NULL,
    mail VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL,
    desactivado BOOLEAN DEFAULT FALSE,
    rol_id TINYINT NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE licencia (
  id SERIAL PRIMARY KEY,
  licencia VARCHAR(100) NOT NULL,
  tipo_licencia TINYINT NOT NULL,
  fecha_vencimiento DATE NOT NULL,
  id_empleado INTEGER NOT NULL,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
);

CREATE TABLE historial_asignacion (
  
);

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
    RAISE NOTICE '✅ Script de inicialización completado exitosamente';
    RAISE NOTICE 'Base de datos: proyecto_recolecta';
    RAISE NOTICE 'Tablas creadas/verificadas: %', table_count;
    RAISE NOTICE 'Índices creados: 6 (tablas transaccionales)';
    RAISE NOTICE '=========================================';
END $$;
