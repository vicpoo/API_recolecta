-- =====================
-- DATABASE
-- =====================
CREATE DATABASE proyecto_recolecta;

-- =====================
-- TABLA: rol
-- =====================
CREATE TABLE rol (
  role_id SERIAL PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL,
  eliminado BOOLEAN DEFAULT FALSE
);

-- =====================
-- TABLA: colonia
-- =====================
CREATE TABLE colonia (
  colonia_id SERIAL PRIMARY KEY,
  nombre VARCHAR(255) NOT NULL,
  zona VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================
-- TABLA: usuario
-- =====================
CREATE TABLE usuario (
  user_id SERIAL PRIMARY KEY,
  nombre VARCHAR(255) NOT NULL,
  alias VARCHAR(100),
  telefono VARCHAR(10),
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role_id INT,
  residencia_id INT,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_usuario_rol FOREIGN KEY (role_id)
    REFERENCES rol(role_id)
);

-- =====================
-- TABLA: domicilio
-- =====================
CREATE TABLE domicilio (
  domicilio_id SERIAL PRIMARY KEY,
  usuario_id INT,
  alias VARCHAR(100),
  direccion VARCHAR(255),
  colonia_id INT,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_domicilio_usuario FOREIGN KEY (usuario_id)
    REFERENCES usuario(user_id),
  CONSTRAINT fk_domicilio_colonia FOREIGN KEY (colonia_id)
    REFERENCES colonia(colonia_id)
);

ALTER TABLE usuario
ADD CONSTRAINT fk_usuario_domicilio
FOREIGN KEY (residencia_id)
REFERENCES domicilio(domicilio_id);

-- =====================
-- TABLA: tipo_camion
-- =====================
CREATE TABLE tipo_camion (
  tipo_camion_id SERIAL PRIMARY KEY,
  nombre VARCHAR(100),
  descripcion VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================
-- TABLA: camion
-- =====================
CREATE TABLE camion (
  camion_id SERIAL PRIMARY KEY,
  placa VARCHAR(20) UNIQUE,
  modelo VARCHAR(100),
  tipo_camion_id INT,
  es_rentado BOOLEAN DEFAULT FALSE,
  eliminado BOOLEAN DEFAULT FALSE,
  disponibilidad_id INT DEFAULT 1,
  nombre_disponibilidad VARCHAR(50),
  color_disponibilidad VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_camion_tipo FOREIGN KEY (tipo_camion_id)
    REFERENCES tipo_camion(tipo_camion_id)
);

-- =====================
-- TABLA: estado_camion
-- =====================
CREATE TABLE estado_camion (
  estado_id SERIAL PRIMARY KEY,
  camion_id INT,
  estado VARCHAR(50),
  timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  observaciones TEXT,
  CONSTRAINT fk_estado_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id)
);

-- =====================
-- TABLA: historial_asignacion_camion
-- =====================
CREATE TABLE historial_asignacion_camion (
  id_historial SERIAL PRIMARY KEY,
  id_chofer INT,
  id_camion INT,
  fecha_asignacion TIMESTAMP,
  fecha_baja TIMESTAMP,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_historial_chofer FOREIGN KEY (id_chofer)
    REFERENCES usuario(user_id),
  CONSTRAINT fk_historial_camion FOREIGN KEY (id_camion)
    REFERENCES camion(camion_id)
);

-- =====================
-- TABLA: ruta
-- =====================
CREATE TABLE ruta (
  ruta_id SERIAL PRIMARY KEY,
  nombre VARCHAR(255),
  descripcion VARCHAR(255),
  json_ruta JSON,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================
-- TABLA: punto_recoleccion
-- =====================
CREATE TABLE punto_recoleccion (
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
CREATE TABLE ruta_camion (
  ruta_camion_id SERIAL PRIMARY KEY,
  ruta_id INT,
  camion_id INT,
  fecha DATE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  eliminado BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_ruta_camion_ruta FOREIGN KEY (ruta_id)
    REFERENCES ruta(ruta_id),
  CONSTRAINT fk_ruta_camion_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id)
);

-- =====================
-- TABLA: tipo_mantenimiento
-- =====================
CREATE TABLE tipo_mantenimiento (
  tipo_mantenimiento_id SERIAL PRIMARY KEY,
  nombre VARCHAR(100),
  categoria VARCHAR(20),
  eliminado BOOLEAN DEFAULT FALSE
);

-- =====================
-- TABLA: alerta_mantenimiento
-- =====================
CREATE TABLE alerta_mantenimiento (
  alerta_id SERIAL PRIMARY KEY,
  camion_id INT,
  tipo_mantenimiento_id INT,
  descripcion VARCHAR(255),
  observaciones TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  atendido BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_alerta_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id),
  CONSTRAINT fk_alerta_tipo FOREIGN KEY (tipo_mantenimiento_id)
    REFERENCES tipo_mantenimiento(tipo_mantenimiento_id)
);

-- =====================
-- TABLA: registro_mantenimiento
-- =====================
CREATE TABLE registro_mantenimiento (
  registro_id SERIAL PRIMARY KEY,
  alerta_id INT,
  camion_id INT,
  coordinador_id INT,
  mecanico_responsable VARCHAR(255),
  fecha_realizada TIMESTAMP,
  kilometraje_mantenimiento DOUBLE PRECISION,
  observaciones TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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
CREATE TABLE incidencia (
  incidencia_id SERIAL PRIMARY KEY,
  punto_recoleccion_id INT,
  conductor_id INT,
  descripcion VARCHAR(255),
  json_ruta JSON,
  fecha_reporte TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  eliminado BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_incidencia_punto FOREIGN KEY (punto_recoleccion_id)
    REFERENCES punto_recoleccion(punto_id),
  CONSTRAINT fk_incidencia_conductor FOREIGN KEY (conductor_id)
    REFERENCES usuario(user_id)
);

-- =====================
-- TABLA: reporte_falla_critica
-- =====================
CREATE TABLE reporte_falla_critica (
  falla_id SERIAL PRIMARY KEY,
  camion_id INT,
  conductor_id INT,
  descripcion VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  eliminado BOOLEAN DEFAULT FALSE,
  CONSTRAINT fk_falla_camion FOREIGN KEY (camion_id)
    REFERENCES camion(camion_id),
  CONSTRAINT fk_falla_conductor FOREIGN KEY (conductor_id)
    REFERENCES usuario(user_id)
);

-- =====================
-- TABLA: seguimiento_falla_critica
-- =====================
CREATE TABLE seguimiento_falla_critica (
  seguimiento_id SERIAL PRIMARY KEY,
  falla_id INT,
  comentario TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_seguimiento_falla FOREIGN KEY (falla_id)
    REFERENCES reporte_falla_critica(falla_id)
);

-- =====================
-- TABLA: anomalia
-- =====================
CREATE TABLE anomalia (
  anomalia_id SERIAL PRIMARY KEY,
  punto_id INT,
  tipo_anomalia VARCHAR(50),
  descripcion TEXT,
  fecha_reporte TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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
CREATE TABLE relleno_sanitario (
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
CREATE TABLE registro_vaciado (
  vaciado_id SERIAL PRIMARY KEY,
  relleno_id INT,
  ruta_camion_id INT,
  hora TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_vaciado_relleno FOREIGN KEY (relleno_id)
    REFERENCES relleno_sanitario(relleno_id),
  CONSTRAINT fk_vaciado_ruta_camion FOREIGN KEY (ruta_camion_id)
    REFERENCES ruta_camion(ruta_camion_id)
);

-- =====================
-- TABLA: notificacion
-- =====================
CREATE TABLE notificacion (
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
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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
CREATE TABLE reporte_conductor (
  reporte_id SERIAL PRIMARY KEY,
  conductor_id INT,
  camion_id INT,
  ruta_id INT,
  descripcion VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
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
CREATE TABLE reporte_mantenimiento_generado (
  reporte_id SERIAL PRIMARY KEY,
  coordinador_id INT,
  fecha_desde TIMESTAMP,
  fecha_hasta TIMESTAMP,
  observaciones VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_reporte_mantenimiento_coordinador FOREIGN KEY (coordinador_id)
    REFERENCES usuario(user_id)
);

-- =====================
-- TABLA: aviso_general
-- =====================
CREATE TABLE aviso_general (
  aviso_id SERIAL PRIMARY KEY,
  titulo VARCHAR(50),
  mensaje TEXT,
  activo BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================
-- TABLA: alerta_usuario
-- =====================
CREATE TABLE alerta_usuario (
  alerta_id SERIAL PRIMARY KEY,
  titulo VARCHAR(50),
  mensaje TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================
-- INSERTS: rol
-- =====================
INSERT INTO rol (nombre, eliminado) VALUES
-- Datos correctos
('Administrador', false),
('Coordinador', false),
-- Datos basura
(NULL, true),
('RolInvalido1234567890123456789012345678901234567890', true);

-- =====================
-- INSERTS: colonia
-- =====================
INSERT INTO colonia (nombre, zona, created_at) VALUES
-- Datos correctos
('Centro Histórico', 'Norte', CURRENT_TIMESTAMP),
('Jardines del Valle', 'Sur', CURRENT_TIMESTAMP),
-- Datos basura
('', NULL, CURRENT_TIMESTAMP),
('ColoniaMuyLarga1234567890123456789012345678901234567890123456789012345678901234567890', 'ZonaMuyLarga123456789012345678901234567890', CURRENT_TIMESTAMP);

-- =====================
-- INSERTS: usuario (sin residencia_id por ahora)
-- =====================
INSERT INTO usuario (nombre, alias, telefono, email, password, role_id, eliminado) VALUES
-- Datos correctos
('Juan Pérez', 'Juancho', '5512345678', 'juan.perez@email.com', 'hashed_password_123', 1, false),
('María García', 'Mary', '5587654321', 'maria.garcia@email.com', 'hashed_password_456', 2, false),
-- Datos basura
('', NULL, '123', 'invalido@', 'pwd', NULL, true),
('NombreMuyLargo1234567890123456789012345678901234567890', 'AliasMuyLargo1234567890123456789012345678901234567890', '12345678901', 'emailinvalido@@dominio.com', 'p', 999, true);

-- =====================
-- INSERTS: domicilio
-- =====================
INSERT INTO domicilio (usuario_id, alias, direccion, colonia_id, eliminado) VALUES
-- Datos correctos
(1, 'Casa', 'Calle Principal 123', 1, false),
(2, 'Oficina', 'Avenida Central 456', 2, false),
-- Datos basura
(NULL, '', '', 999, true),
(999, 'AliasMuyLargo1234567890123456789012345678901234567890', 'DirecciónMuyLarga123456789012345678901234567890123456789012345678901234567890', 999, true);

-- =====================
-- Actualizar usuarios con residencia_id
-- =====================
UPDATE usuario SET residencia_id = 1 WHERE user_id = 1;
UPDATE usuario SET residencia_id = 2 WHERE user_id = 2;

-- =====================
-- INSERTS: tipo_camion
-- =====================
INSERT INTO tipo_camion (nombre, descripcion, created_at) VALUES
-- Datos correctos
('Compactador', 'Camión para basura compactada', CURRENT_TIMESTAMP),
('Volteo', 'Camión de volteo para carga pesada', CURRENT_TIMESTAMP),
-- Datos basura
('', NULL, CURRENT_TIMESTAMP),
('TipoMuyLargo1234567890123456789012345678901234567890', 'DescripciónMuyLarga1234567890123456789012345678901234567890123456789012345678901234567890', CURRENT_TIMESTAMP);

-- =====================
-- INSERTS: camion
-- =====================
INSERT INTO camion (placa, modelo, tipo_camion_id, es_rentado, eliminado, nombre_disponibilidad, color_disponibilidad) VALUES
-- Datos correctos
('ABC123', 'Kenworth T880', 1, false, false, 'Disponible', 'verde'),
('XYZ789', 'Freightliner M2', 2, true, false, 'En ruta', 'amarillo'),
-- Datos basura
('', NULL, 999, true, true, NULL, NULL),
('PLACAMUYLARGA1234567', 'ModeloMuyLargo1234567890123456789012345678901234567890', 999, true, true, 'NombreDisponibilidadMuyLargo12345678901234567890', 'ColorMuyLargo1234567890');

-- =====================
-- INSERTS: estado_camion
-- =====================
INSERT INTO estado_camion (camion_id, estado, observaciones) VALUES
-- Datos correctos
(1, 'Operativo', 'En perfecto estado'),
(2, 'En mantenimiento', 'Cambio de aceite programado'),
-- Datos basura
(999, '', NULL),
(999, 'EstadoMuyLargo1234567890123456789012345678901234567890', 'ObservaciónMuyLarga12345678901234567890123456789012345678901234567890123456789012345678901234567890');

-- =====================
-- INSERTS: historial_asignacion_camion
-- =====================
INSERT INTO historial_asignacion_camion (id_chofer, id_camion, fecha_asignacion, fecha_baja, eliminado) VALUES
-- Datos correctos
(1, 1, '2024-01-15 08:00:00', NULL, false),
(2, 2, '2024-01-20 09:00:00', '2024-01-25 17:00:00', false),
-- Datos basura
(999, 999, NULL, NULL, true),
(999, 999, '2024-13-45 25:61:00', '2024-13-45 25:61:00', true);

-- =====================
-- INSERTS: ruta
-- =====================
INSERT INTO ruta (nombre, descripcion, json_ruta, eliminado) VALUES
-- Datos correctos
('Ruta Norte', 'Recolección zona norte', '{"puntos": [1,2,3]}', false),
('Ruta Sur', 'Recolección zona sur', '{"puntos": [4,5,6]}', false),
-- Datos basura
('', NULL, '{}', true),
('NombreMuyLargo1234567890123456789012345678901234567890', 'DescripciónMuyLarga123456789012345678901234567890123456789012345678901234567890', '{"invalid": "json"', true);

-- =====================
-- INSERTS: punto_recoleccion
-- =====================
INSERT INTO punto_recoleccion (ruta_id, cp, eliminado) VALUES
-- Datos correctos
(1, '01000', false),
(2, '02000', false),
-- Datos basura
(999, '', true),
(999, 'CP12345678901234567890', true);

-- =====================
-- INSERTS: ruta_camion
-- =====================
INSERT INTO ruta_camion (ruta_id, camion_id, fecha, eliminado) VALUES
-- Datos correctos
(1, 1, '2024-01-26', false),
(2, 2, '2024-01-25', false),
-- Datos basura
(999, 999, NULL, true),
(999, 999, '2024-13-45', true);

-- =====================
-- INSERTS: tipo_mantenimiento
-- =====================
INSERT INTO tipo_mantenimiento (nombre, categoria, eliminado) VALUES
-- Datos correctos
('Cambio de aceite', 'Preventivo', false),
('Reparación de frenos', 'Correctivo', false),
-- Datos basura
('', '', true),
('NombreMuyLargo1234567890123456789012345678901234567890', 'CategoriaMuyLarga1234567890', true);

-- =====================
-- INSERTS: alerta_mantenimiento
-- =====================
INSERT INTO alerta_mantenimiento (camion_id, tipo_mantenimiento_id, descripcion, observaciones, atendido) VALUES
-- Datos correctos
(1, 1, 'Aceite bajo nivel', 'Verificar fuga posible', false),
(2, 2, 'Frenos desgastados', 'Pastillas al 10%', true),
-- Datos basura
(999, 999, '', NULL, true),
(999, 999, 'DescripciónMuyLarga12345678901234567890123456789012345678901234567890123456789012345678901234567890', 'ObservaciónMuyLarga' || REPEAT('x', 1000), true);

-- =====================
-- INSERTS: registro_mantenimiento
-- =====================
INSERT INTO registro_mantenimiento (alerta_id, camion_id, coordinador_id, mecanico_responsable, fecha_realizada, kilometraje_mantenimiento, observaciones) VALUES
-- Datos correctos
(1, 1, 2, 'José Martínez', '2024-01-20 10:00:00', 50000.5, 'Cambio completo de aceite y filtro'),
(2, 2, 1, 'Carlos López', '2024-01-22 14:30:00', 75000.0, 'Cambio de pastillas y discos'),
-- Datos basura
(999, 999, 999, '', NULL, -1000.0, NULL),
(999, 999, 999, 'NombreMuyLargo1234567890123456789012345678901234567890', '2024-13-45 25:61:00', 9999999999.99, REPEAT('x', 1000));

-- =====================
-- INSERTS: incidencia
-- =====================
INSERT INTO incidencia (punto_recoleccion_id, conductor_id, descripcion, json_ruta, eliminado) VALUES
-- Datos correctos
(1, 1, 'Contenedor bloqueado', '{"ubicacion": "Calle Principal"}', false),
(2, 2, 'Tráfico pesado', '{"alternativa": "Calle Alterna"}', false),
-- Datos basura
(999, 999, '', NULL, true),
(999, 999, 'DescripciónMuyLarga12345678901234567890123456789012345678901234567890123456789012345678901234567890', '{"invalid": "json"', true);

-- =====================
-- INSERTS: reporte_falla_critica
-- =====================
INSERT INTO reporte_falla_critica (camion_id, conductor_id, descripcion, eliminado) VALUES
-- Datos correctos
(1, 1, 'Fuga de combustible grave', false),
(2, 2, 'Falla en sistema de frenos', false),
-- Datos basura
(999, 999, '', true),
(999, 999, 'DescripciónMuyLarga12345678901234567890123456789012345678901234567890123456789012345678901234567890', true);

-- =====================
-- INSERTS: seguimiento_falla_critica
-- =====================
INSERT INTO seguimiento_falla_critica (falla_id, comentario) VALUES
-- Datos correctos
(1, 'Enviado equipo de soporte'),
(2, 'Camión remolcado al taller'),
-- Datos basura
(999, ''),
(999, REPEAT('x', 1000));

-- =====================
-- INSERTS: anomalia
-- =====================
INSERT INTO anomalia (punto_id, tipo_anomalia, descripcion, estado, id_chofer_id) VALUES
-- Datos correctos
(1, 'Accidente', 'Colisión con vehículo', 'Reportado', 1),
(2, 'Obstrucción', 'Árbol caído en ruta', 'Resuelto', 2),
-- Datos basura
(999, '', '', NULL, 999),
(999, 'TipoMuyLargo1234567890123456789012345678901234567890', REPEAT('x', 1000), 'EstadoMuyLargo12345678901234567890', 999);

-- =====================
-- INSERTS: relleno_sanitario
-- =====================
INSERT INTO relleno_sanitario (nombre, direccion, es_rentado, eliminado, capacidad_toneladas) VALUES
-- Datos correctos
('Relleno Norte', 'Carretera Norte Km 10', false, false, 5000.00),
('Relleno Sur', 'Autopista Sur Km 15', true, false, 7500.50),
-- Datos basura
('', '', true, true, -1000.00),
('NombreMuyLargo1234567890123456789012345678901234567890', 'DirecciónMuyLarga123456789012345678901234567890123456789012345678901234567890', true, true, 9999999999.99);

-- =====================
-- INSERTS: registro_vaciado
-- =====================
INSERT INTO registro_vaciado (relleno_id, ruta_camion_id, hora) VALUES
-- Datos correctos
(1, 1, '2024-01-26 10:30:00'),
(2, 2, '2024-01-25 16:45:00'),
-- Datos basura
(999, 999, NULL),
(999, 999, '2024-13-45 25:61:00');

-- =====================
-- INSERTS: notificacion
-- =====================
INSERT INTO notificacion (usuario_id, tipo, titulo, mensaje, activa, id_camion_relacionado, creado_por) VALUES
-- Datos correctos
(1, 'Mantenimiento', 'Recordatorio', 'Cambio de aceite pendiente', true, 1, 2),
(2, 'Falla', 'Alerta', 'Reporte de falla crítica', true, 2, 1),
-- Datos basura
(999, '', '', '', true, 999, 999),
(999, 'TipoMuyLargo1234567890123456789012345678901234567890', 'TituloMuyLargo12345678901234567890123456789012345678901234567890', REPEAT('x', 1000), true, 999, 999);

-- =====================
-- INSERTS: reporte_conductor
-- =====================
INSERT INTO reporte_conductor (conductor_id, camion_id, ruta_id, descripcion) VALUES
-- Datos correctos
(1, 1, 1, 'Ruta completada sin incidentes'),
(2, 2, 2, 'Retraso por tráfico pesado'),
-- Datos basura
(999, 999, 999, ''),
(999, 999, 999, 'DescripciónMuyLarga12345678901234567890123456789012345678901234567890123456789012345678901234567890');

-- =====================
-- INSERTS: reporte_mantenimiento_generado
-- =====================
INSERT INTO reporte_mantenimiento_generado (coordinador_id, fecha_desde, fecha_hasta, observaciones) VALUES
-- Datos correctos
(1, '2024-01-01 00:00:00', '2024-01-31 23:59:59', 'Mantenimiento preventivo mensual'),
(2, '2024-02-01 00:00:00', '2024-02-28 23:59:59', 'Revisión post-vacaciones'),
-- Datos basura
(999, NULL, NULL, ''),
(999, '2024-13-45 25:61:00', '2024-13-45 25:61:00', 'ObservacionesMuyLargas123456789012345678901234567890123456789012345678901234567890');

-- =====================
-- INSERTS: aviso_general
-- =====================
INSERT INTO aviso_general (titulo, mensaje, activo) VALUES
-- Datos correctos
('Mantenimiento Programado', 'El próximo lunes habrá mantenimiento general', true),
('Feriado', 'No habrá recolección el día festivo', true),
-- Datos basura
('', '', false),
('TituloMuyLargo12345678901234567890123456789012345678901234567890', REPEAT('x', 1000), false);

-- =====================
-- INSERTS: alerta_usuario
-- =====================
INSERT INTO alerta_usuario (titulo, mensaje) VALUES
-- Datos correctos
('Bienvenido', 'Bienvenido al sistema de recolección'),
('Actualización', 'El sistema se actualizará esta noche'),
-- Datos basura
('', ''),
('TituloMuyLargo12345678901234567890123456789012345678901234567890', REPEAT('x', 1000));