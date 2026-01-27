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
-- TABLA: rol
-- =====================
CREATE TABLE IF NOT EXISTS rol (
  role_id SERIAL PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL,
  eliminado BOOLEAN DEFAULT FALSE
);

-- =====================
-- TABLA: colonia
-- =====================
CREATE TABLE IF NOT EXISTS colonia (
  colonia_id SERIAL PRIMARY KEY,
  nombre VARCHAR(255) NOT NULL,
  zona VARCHAR(50),
  created_at TIMESTAMP
);

-- =====================
-- TABLA: usuario
-- =====================
CREATE TABLE IF NOT EXISTS usuario (
  user_id SERIAL PRIMARY KEY,
  nombre VARCHAR(255) NOT NULL,
  alias VARCHAR(100),
  telefono VARCHAR(10),
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role_id INT,
  residencia_id INT,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT fk_usuario_rol FOREIGN KEY (role_id)
    REFERENCES rol(role_id)
);

-- =====================
-- TABLA: domicilio
-- =====================
CREATE TABLE IF NOT EXISTS domicilio (
  domicilio_id SERIAL PRIMARY KEY,
  usuario_id INT,
  alias VARCHAR(100),
  direccion VARCHAR(255),
  colonia_id INT,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT fk_domicilio_usuario FOREIGN KEY (usuario_id)
    REFERENCES usuario(user_id),
  CONSTRAINT fk_domicilio_colonia FOREIGN KEY (colonia_id)
    REFERENCES colonia(colonia_id)
);

-- Constraint idempotente para relación circular usuario<->domicilio
DO $$ 
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'fk_usuario_domicilio'
    ) THEN
        ALTER TABLE usuario ADD CONSTRAINT fk_usuario_domicilio
        FOREIGN KEY (residencia_id) REFERENCES domicilio(domicilio_id);
    END IF;
END $$;

-- =====================
-- TABLA: tipo_camion
-- =====================
CREATE TABLE IF NOT EXISTS tipo_camion (
  tipo_camion_id SERIAL PRIMARY KEY,
  nombre VARCHAR(100),
  descripcion VARCHAR(255),
  created_at TIMESTAMP
);

-- =====================
-- TABLA: camion
-- =====================
CREATE TABLE IF NOT EXISTS camion (
  camion_id SERIAL PRIMARY KEY,
  placa VARCHAR(20) UNIQUE,
  modelo VARCHAR(100),
  tipo_camion_id INT,
  es_rentado BOOLEAN DEFAULT FALSE,
  eliminado BOOLEAN DEFAULT FALSE,
  disponibilidad_id INT DEFAULT 1,
  nombre_disponibilidad VARCHAR(50),
  color_disponibilidad VARCHAR(20),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT fk_camion_tipo FOREIGN KEY (tipo_camion_id)
    REFERENCES tipo_camion(tipo_camion_id)
);

-- =====================
-- TABLA: estado_camion
-- =====================
CREATE TABLE IF NOT EXISTS estado_camion (
  estado_id SERIAL PRIMARY KEY,
  camion_id INT,
  estado VARCHAR(50),
  timestamp TIMESTAMP,
  observaciones TEXT,
  CONSTRAINT fk_estado_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id)
);

-- =====================
-- TABLA: historial_asignacion_camion
-- =====================
CREATE TABLE IF NOT EXISTS historial_asignacion_camion (
  id_historial SERIAL PRIMARY KEY,
  id_chofer INT,
  id_camion INT,
  fecha_asignacion TIMESTAMP,
  fecha_baja TIMESTAMP,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  CONSTRAINT fk_historial_chofer FOREIGN KEY (id_chofer)
    REFERENCES usuario(user_id),
  CONSTRAINT fk_historial_camion FOREIGN KEY (id_camion)
    REFERENCES camion(camion_id)
);

-- =====================
-- TABLA: ruta
-- =====================
CREATE TABLE IF NOT EXISTS ruta (
  ruta_id SERIAL PRIMARY KEY,
  nombre VARCHAR(255),
  descripcion VARCHAR(255),
  json_ruta JSON,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP
);

-- =====================
-- TABLA: punto_recoleccion
-- =====================
CREATE TABLE IF NOT EXISTS punto_recoleccion (
  punto_id SERIAL PRIMARY KEY,
  ruta_id INT,
  cp VARCHAR(20) UNIQUE,
  eliminado BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_punto_ruta FOREIGN KEY (ruta_id)
    REFERENCES ruta(ruta_id)
);

-- =====================
-- TABLA: ruta_camion
-- =====================
CREATE TABLE IF NOT EXISTS ruta_camion (
  ruta_camion_id SERIAL PRIMARY KEY,
  ruta_id INT,
  camion_id INT,
  fecha DATE,
  created_at TIMESTAMP,
  eliminado BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_ruta_camion_ruta FOREIGN KEY (ruta_id)
    REFERENCES ruta(ruta_id),
  CONSTRAINT fk_ruta_camion_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id)
);

-- =====================
-- TABLA: tipo_mantenimiento
-- =====================
CREATE TABLE IF NOT EXISTS tipo_mantenimiento (
  tipo_mantenimiento_id SERIAL PRIMARY KEY,
  nombre VARCHAR(100),
  categoria VARCHAR(20),
  eliminado BOOLEAN DEFAULT FALSE
);

-- =====================
-- TABLA: alerta_mantenimiento
-- =====================
CREATE TABLE IF NOT EXISTS alerta_mantenimiento (
  alerta_id SERIAL PRIMARY KEY,
  camion_id INT,
  tipo_mantenimiento_id INT,
  descripcion VARCHAR(255),
  observaciones TEXT,
  created_at TIMESTAMP,
  atendido BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_alerta_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id),
  CONSTRAINT fk_alerta_tipo FOREIGN KEY (tipo_mantenimiento_id)
    REFERENCES tipo_mantenimiento(tipo_mantenimiento_id)
);

-- =====================
-- TABLA: registro_mantenimiento
-- =====================
CREATE TABLE IF NOT EXISTS registro_mantenimiento (
  registro_id SERIAL PRIMARY KEY,
  alerta_id INT,
  camion_id INT,
  coordinador_id INT,
  mecanico_responsable VARCHAR(255),
  fecha_realizada TIMESTAMP,
  kilometraje_mantenimiento DOUBLE PRECISION,
  observaciones TEXT,
  created_at TIMESTAMP,
  CONSTRAINT fk_registro_alerta FOREIGN KEY (alerta_id)
    REFERENCES alerta_mantenimiento(alerta_id),
  CONSTRAINT fk_registro_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id),
  CONSTRAINT fk_registro_coordinador FOREIGN KEY (coordinador_id)
    REFERENCES usuario(user_id)
);

-- =====================
-- TABLA: incidencia
-- =====================
CREATE TABLE IF NOT EXISTS incidencia (
  incidencia_id SERIAL PRIMARY KEY,
  punto_recoleccion_id INT,
  conductor_id INT,
  descripcion VARCHAR(255),
  json_ruta JSON,
  fecha_reporte TIMESTAMP,
  eliminado BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_incidencia_punto FOREIGN KEY (punto_recoleccion_id)
    REFERENCES punto_recoleccion(punto_id),
  CONSTRAINT fk_incidencia_conductor FOREIGN KEY (conductor_id)
    REFERENCES usuario(user_id)
);

-- =====================
-- TABLA: reporte_falla_critica
-- =====================
CREATE TABLE IF NOT EXISTS reporte_falla_critica (
  falla_id SERIAL PRIMARY KEY,
  camion_id INT,
  conductor_id INT,
  descripcion VARCHAR(255),
  created_at TIMESTAMP,
  eliminado BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_falla_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id),
  CONSTRAINT fk_falla_conductor FOREIGN KEY (conductor_id)
    REFERENCES usuario(user_id)
);

-- =====================
-- TABLA: seguimiento_falla_critica
-- =====================
CREATE TABLE IF NOT EXISTS seguimiento_falla_critica (
  seguimiento_id SERIAL PRIMARY KEY,
  falla_id INT,
  comentario TEXT,
  created_at TIMESTAMP,
  CONSTRAINT fk_seguimiento_falla FOREIGN KEY (falla_id)
    REFERENCES reporte_falla_critica(falla_id)
);

-- =====================
-- TABLA: anomalia
-- =====================
CREATE TABLE IF NOT EXISTS anomalia (
  anomalia_id SERIAL PRIMARY KEY,
  punto_id INT,
  tipo_anomalia VARCHAR(50),
  descripcion TEXT,
  fecha_reporte TIMESTAMP,
  estado VARCHAR(30),
  fecha_resolucion TIMESTAMP,
  id_chofer_id INT,
  CONSTRAINT fk_anomalia_punto FOREIGN KEY (punto_id)
    REFERENCES punto_recoleccion(punto_id),
  CONSTRAINT fk_anomalia_chofer FOREIGN KEY (id_chofer_id)
    REFERENCES usuario(user_id)
);

-- =====================
-- TABLA: relleno_sanitario
-- =====================
CREATE TABLE IF NOT EXISTS relleno_sanitario (
  relleno_id SERIAL PRIMARY KEY,
  nombre VARCHAR(255),
  direccion VARCHAR(255),
  es_rentado BOOLEAN DEFAULT FALSE,
  eliminado BOOLEAN DEFAULT FALSE,
  capacidad_toneladas DECIMAL(10,2)
);

-- =====================
-- TABLA: registro_vaciado
-- =====================
CREATE TABLE IF NOT EXISTS registro_vaciado (
  vaciado_id SERIAL PRIMARY KEY,
  relleno_id INT,
  ruta_camion_id INT,
  hora TIMESTAMP,
  CONSTRAINT fk_vaciado_relleno FOREIGN KEY (relleno_id)
    REFERENCES relleno_sanitario(relleno_id),
  CONSTRAINT fk_vaciado_ruta_camion FOREIGN KEY (ruta_camion_id)
    REFERENCES ruta_camion(ruta_camion_id)
);

-- =====================
-- TABLA: notificacion
-- =====================
CREATE TABLE IF NOT EXISTS notificacion (
  notificacion_id SERIAL PRIMARY KEY,
  usuario_id INT,
  tipo VARCHAR(50),
  titulo VARCHAR(100),
  mensaje TEXT,
  activa BOOLEAN DEFAULT TRUE,
  id_camion_relacionado INT,
  id_falla_relacionado INT,
  id_mantenimiento_relacionado INT,
  creado_por INT,
  created_at TIMESTAMP,
  CONSTRAINT fk_notif_usuario FOREIGN KEY (usuario_id)
    REFERENCES usuario(user_id),
  CONSTRAINT fk_notif_camion FOREIGN KEY (id_camion_relacionado)
    REFERENCES camion(camion_id),
  CONSTRAINT fk_notif_falla FOREIGN KEY (id_falla_relacionado)
    REFERENCES reporte_falla_critica(falla_id),
  CONSTRAINT fk_notif_mantenimiento FOREIGN KEY (id_mantenimiento_relacionado)
    REFERENCES registro_mantenimiento(registro_id),
  CONSTRAINT fk_notif_creador FOREIGN KEY (creado_por)
    REFERENCES usuario(user_id)
);

-- =====================
-- TABLA: reporte_conductor
-- =====================
CREATE TABLE IF NOT EXISTS reporte_conductor (
  reporte_id SERIAL PRIMARY KEY,
  conductor_id INT,
  camion_id INT,
  ruta_id INT,
  descripcion VARCHAR(255),
  created_at TIMESTAMP,
  CONSTRAINT fk_reporte_conductor FOREIGN KEY (conductor_id)
    REFERENCES usuario(user_id),
  CONSTRAINT fk_reporte_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id),
  CONSTRAINT fk_reporte_ruta FOREIGN KEY (ruta_id)
    REFERENCES ruta(ruta_id)
);

-- =====================
-- TABLA: reporte_mantenimiento_generado
-- =====================
CREATE TABLE IF NOT EXISTS reporte_mantenimiento_generado (
  reporte_id SERIAL PRIMARY KEY,
  coordinador_id INT,
  fecha_desde TIMESTAMP,
  fecha_hasta TIMESTAMP,
  observaciones VARCHAR(255),
  created_at TIMESTAMP,
  CONSTRAINT fk_reporte_mantenimiento_coordinador FOREIGN KEY (coordinador_id)
    REFERENCES usuario(user_id)
);

-- =====================
-- TABLA: aviso_general
-- =====================
CREATE TABLE IF NOT EXISTS aviso_general (
  aviso_id SERIAL PRIMARY KEY,
  titulo VARCHAR(50),
  mensaje TEXT,
  activo BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP
);

-- =====================
-- TABLA: alerta_usuario
-- =====================
CREATE TABLE IF NOT EXISTS alerta_usuario (
  alerta_id SERIAL PRIMARY KEY,
  titulo VARCHAR(50),
  mensaje TEXT,
  created_at TIMESTAMP
);

-- =====================
-- ÍNDICES (Solo tablas transaccionales de alto volumen)
-- =====================

-- notificacion: tabla de eventos/logs, crecimiento continuo
CREATE INDEX IF NOT EXISTS idx_notificacion_usuario_id ON notificacion(usuario_id);
CREATE INDEX IF NOT EXISTS idx_notificacion_created_at ON notificacion(created_at DESC);

-- incidencia: reportes diarios, alta escritura
CREATE INDEX IF NOT EXISTS idx_incidencia_conductor_id ON incidencia(conductor_id);
CREATE INDEX IF NOT EXISTS idx_incidencia_fecha_reporte ON incidencia(fecha_reporte DESC);

-- estado_camion: tracking continuo
CREATE INDEX IF NOT EXISTS idx_estado_camion_camion_id ON estado_camion(camion_id);
CREATE INDEX IF NOT EXISTS idx_estado_camion_timestamp ON estado_camion(timestamp DESC);

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
