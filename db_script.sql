CREATE DATABASE IF NOT EXISTS proyecto_recolecta;
USE proyecto_recolecta;

-- =====================
-- TABLA: rol
-- =====================
CREATE TABLE rol (
  role_id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(50) NOT NULL,
  eliminado BOOLEAN DEFAULT FALSE
) ENGINE=InnoDB;

-- =====================
-- TABLA: colonia
-- =====================
CREATE TABLE colonia (
  colonia_id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(255) NOT NULL,
  zona VARCHAR(50),
  created_at DATETIME
) ENGINE=InnoDB;

-- =====================
-- TABLA: usuario
-- =====================
CREATE TABLE usuario (
  user_id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(255) NOT NULL,
  alias VARCHAR(100),
  telefono VARCHAR(10),
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role_id INT,
  residencia_id INT,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (role_id) REFERENCES rol(role_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: domicilio
-- =====================
CREATE TABLE domicilio (
  domicilio_id INT AUTO_INCREMENT PRIMARY KEY,
  usuario_id INT,
  alias VARCHAR(100),
  direccion VARCHAR(255),
  colonia_id INT,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (usuario_id) REFERENCES usuario(user_id),
  FOREIGN KEY (colonia_id) REFERENCES colonia(colonia_id)
) ENGINE=InnoDB;

ALTER TABLE usuario
ADD CONSTRAINT fk_usuario_domicilio
FOREIGN KEY (residencia_id) REFERENCES domicilio(domicilio_id);

-- =====================
-- TABLA: tipo_camion
-- =====================
CREATE TABLE tipo_camion (
  tipo_camion_id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(100),
  descripcion VARCHAR(255),
  created_at DATETIME
) ENGINE=InnoDB;

-- =====================
-- TABLA: camion
-- =====================
CREATE TABLE camion (
  camion_id INT AUTO_INCREMENT PRIMARY KEY,
  placa VARCHAR(20) UNIQUE,
  modelo VARCHAR(100),
  tipo_camion_id INT,
  es_rentado BOOLEAN DEFAULT FALSE,
  eliminado BOOLEAN DEFAULT FALSE,
  disponibilidad_id INT DEFAULT 1,
  nombre_disponibilidad VARCHAR(50),
  color_disponibilidad VARCHAR(20),
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (tipo_camion_id) REFERENCES tipo_camion(tipo_camion_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: estado_camion
-- =====================
CREATE TABLE estado_camion (
  estado_id INT AUTO_INCREMENT PRIMARY KEY,
  camion_id INT,
  estado VARCHAR(50),
  timestamp DATETIME,
  observaciones TEXT,
  FOREIGN KEY (camion_id) REFERENCES camion(camion_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: historial_asignacion_camion
-- =====================
CREATE TABLE historial_asignacion_camion (
  id_historial INT AUTO_INCREMENT PRIMARY KEY,
  id_chofer INT,
  id_camion INT,
  fecha_asignacion DATETIME,
  fecha_baja DATETIME,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at DATETIME,
  updated_at DATETIME,
  FOREIGN KEY (id_chofer) REFERENCES usuario(user_id),
  FOREIGN KEY (id_camion) REFERENCES camion(camion_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: ruta
-- =====================
CREATE TABLE ruta (
  ruta_id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(255),
  descripcion VARCHAR(255),
  json_ruta TEXT,
  eliminado BOOLEAN DEFAULT FALSE,
  created_at DATETIME
) ENGINE=InnoDB;

-- =====================
-- TABLA: punto_recoleccion
-- =====================
CREATE TABLE punto_recoleccion (
  punto_id INT AUTO_INCREMENT PRIMARY KEY,
  ruta_id INT,
  cp VARCHAR(20) UNIQUE,
  eliminado BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (ruta_id) REFERENCES ruta(ruta_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: ruta_camion
-- =====================
CREATE TABLE ruta_camion (
  ruta_camion_id INT AUTO_INCREMENT PRIMARY KEY,
  ruta_id INT,
  camion_id INT,
  fecha DATE,
  created_at DATETIME,
  eliminado BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (ruta_id) REFERENCES ruta(ruta_id),
  FOREIGN KEY (camion_id) REFERENCES camion(camion_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: tipo_mantenimiento
-- =====================
CREATE TABLE tipo_mantenimiento (
  tipo_mantenimiento_id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(100),
  categoria VARCHAR(20),
  eliminado BOOLEAN DEFAULT FALSE
) ENGINE=InnoDB;

-- =====================
-- TABLA: alerta_mantenimiento
-- =====================
CREATE TABLE alerta_mantenimiento (
  alerta_id INT AUTO_INCREMENT PRIMARY KEY,
  camion_id INT,
  tipo_mantenimiento_id INT,
  descripcion VARCHAR(255),
  observaciones TEXT,
  created_at DATETIME,
  atendido BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (camion_id) REFERENCES camion(camion_id),
  FOREIGN KEY (tipo_mantenimiento_id) REFERENCES tipo_mantenimiento(tipo_mantenimiento_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: registro_mantenimiento
-- =====================
CREATE TABLE registro_mantenimiento (
  registro_id INT AUTO_INCREMENT PRIMARY KEY,
  alerta_id INT,
  camion_id INT,
  coordinador_id INT,
  mecanico_responsable VARCHAR(255),
  fecha_realizada DATETIME,
  kilometraje_mantenimiento FLOAT,
  observaciones TEXT,
  created_at DATETIME,
  FOREIGN KEY (alerta_id) REFERENCES alerta_mantenimiento(alerta_id),
  FOREIGN KEY (camion_id) REFERENCES camion(camion_id),
  FOREIGN KEY (coordinador_id) REFERENCES usuario(user_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: incidencia
-- =====================
CREATE TABLE incidencia (
  incidencia_id INT AUTO_INCREMENT PRIMARY KEY,
  punto_recoleccion_id INT,
  conductor_id INT,
  descripcion VARCHAR(255),
  json_ruta TEXT,
  fecha_reporte DATETIME,
  eliminado BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (punto_recoleccion_id) REFERENCES punto_recoleccion(punto_id),
  FOREIGN KEY (conductor_id) REFERENCES usuario(user_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: reporte_falla_critica
-- =====================
CREATE TABLE reporte_falla_critica (
  falla_id INT AUTO_INCREMENT PRIMARY KEY,
  camion_id INT,
  conductor_id INT,
  descripcion VARCHAR(255),
  created_at DATETIME,
  eliminado BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (camion_id) REFERENCES camion(camion_id),
  FOREIGN KEY (conductor_id) REFERENCES usuario(user_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: seguimiento_falla_critica
-- =====================
CREATE TABLE seguimiento_falla_critica (
  seguimiento_id INT AUTO_INCREMENT PRIMARY KEY,
  falla_id INT,
  comentario TEXT,
  created_at DATETIME,
  FOREIGN KEY (falla_id) REFERENCES reporte_falla_critica(falla_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: anomalia
-- =====================
CREATE TABLE anomalia (
  anomalia_id INT AUTO_INCREMENT PRIMARY KEY,
  punto_id INT,
  tipo_anomalia VARCHAR(50),
  descripcion TEXT,
  fecha_reporte DATETIME,
  estado VARCHAR(30),
  fecha_resolucion DATETIME,
  id_chofer_id INT,
  FOREIGN KEY (punto_id) REFERENCES punto_recoleccion(punto_id),
  FOREIGN KEY (id_chofer_id) REFERENCES usuario(user_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: relleno_sanitario
-- =====================
CREATE TABLE relleno_sanitario (
  relleno_id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(255),
  direccion VARCHAR(255),
  es_rentado BOOLEAN DEFAULT FALSE,
  eliminado BOOLEAN DEFAULT FALSE,
  capacidad_toneladas DECIMAL(10,2)
) ENGINE=InnoDB;

-- =====================
-- TABLA: registro_vaciado
-- =====================
CREATE TABLE registro_vaciado (
  vaciado_id INT AUTO_INCREMENT PRIMARY KEY,
  relleno_id INT,
  ruta_camion_id INT,
  hora DATETIME,
  FOREIGN KEY (relleno_id) REFERENCES relleno_sanitario(relleno_id),
  FOREIGN KEY (ruta_camion_id) REFERENCES ruta_camion(ruta_camion_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: notificacion
-- =====================
CREATE TABLE notificacion (
  notificacion_id INT AUTO_INCREMENT PRIMARY KEY,
  usuario_id INT,
  tipo VARCHAR(50),
  titulo VARCHAR(100),
  mensaje TEXT,
  activa BOOLEAN DEFAULT TRUE,
  id_camion_relacionado INT,
  id_falla_relacionado INT,
  id_mantenimiento_relacionado INT,
  creado_por INT,
  created_at DATETIME,
  FOREIGN KEY (usuario_id) REFERENCES usuario(user_id),
  FOREIGN KEY (id_camion_relacionado) REFERENCES camion(camion_id),
  FOREIGN KEY (id_falla_relacionado) REFERENCES reporte_falla_critica(falla_id),
  FOREIGN KEY (id_mantenimiento_relacionado) REFERENCES registro_mantenimiento(registro_id),
  FOREIGN KEY (creado_por) REFERENCES usuario(user_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: reporte_conductor
-- =====================
CREATE TABLE reporte_conductor (
  reporte_id INT AUTO_INCREMENT PRIMARY KEY,
  conductor_id INT,
  camion_id INT,
  ruta_id INT,
  descripcion VARCHAR(255),
  created_at DATETIME,
  FOREIGN KEY (conductor_id) REFERENCES usuario(user_id),
  FOREIGN KEY (camion_id) REFERENCES camion(camion_id),
  FOREIGN KEY (ruta_id) REFERENCES ruta(ruta_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: reporte_mantenimiento_generado
-- =====================
CREATE TABLE reporte_mantenimiento_generado (
  reporte_id INT AUTO_INCREMENT PRIMARY KEY,
  coordinador_id INT,
  fecha_desde DATETIME,
  fecha_hasta DATETIME,
  observaciones VARCHAR(255),
  created_at DATETIME,
  FOREIGN KEY (coordinador_id) REFERENCES usuario(user_id)
) ENGINE=InnoDB;

-- =====================
-- TABLA: aviso_general
-- =====================
CREATE TABLE aviso_general (
  aviso_id INT AUTO_INCREMENT PRIMARY KEY,
  titulo VARCHAR(50),
  mensaje TEXT,
  activo BOOLEAN DEFAULT TRUE,
  created_at DATETIME
) ENGINE=InnoDB;

-- =====================
-- TABLA: alerta_usuario
-- =====================
CREATE TABLE alerta_usuario (
  alerta_id INT AUTO_INCREMENT PRIMARY KEY,
  titulo VARCHAR(50),
  mensaje TEXT,
  created_at DATETIME
  ) ENGINE=InnoDB;