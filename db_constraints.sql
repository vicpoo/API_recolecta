\c proyecto_recolecta;

-- =====================
-- CONSTRAINTS
-- =====================

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'uq_rol_nombre'
    ) THEN
        ALTER TABLE rol ADD CONSTRAINT uq_rol_nombre UNIQUE (nombre);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_rol_empleado'
    ) THEN
        ALTER TABLE empleado ADD CONSTRAINT fk_rol FOREIGN KEY (rol_id) REFERENCES rol(id);
    END IF;
    
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'uq_mail_empleado'
    ) THEN
        ALTER TABLE empleado ADD CONSTRAINT uq_mail_empleado UNIQUE (mail);
    END IF;
    
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'uq_username_empleado'
    ) THEN
        ALTER TABLE empleado ADD CONSTRAINT uq_username_empleado UNIQUE (username);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_empleado_licencia'
    ) THEN
        ALTER TABLE licencia ADD CONSTRAINT fk_empleado_licencia FOREIGN KEY (id_empleado) REFERENCES empleado(id);
    END IF;
    
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'uq_licencia'
    ) THEN
        ALTER TABLE licencia ADD CONSTRAINT uq_licencia UNIQUE (licencia);
    END IF;
    
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_fecha_vencimiento'
    ) THEN
        ALTER TABLE licencia ADD CONSTRAINT chk_fecha_vencimiento CHECK (fecha_vencimiento > CURRENT_DATE);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_empleado_licencia'
    ) THEN
        ALTER TABLE licencia ADD CONSTRAINT fk_empleado_licencia FOREIGN KEY (id_empleado) REFERENCES empleado(id);
    END IF;
END $$;


DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_chofer_historial'
    ) THEN
        ALTER TABLE historial_asignacion_camion ADD CONSTRAINT fk_chofer_historial FOREIGN KEY (id_chofer) REFERENCES empleado(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_camion_historial'
    ) THEN
        ALTER TABLE historial_asignacion_camion ADD CONSTRAINT fk_camion_historial FOREIGN KEY (id_camion) REFERENCES camion(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'uq_chofer_camion_historial'
    ) THEN
        ALTER TABLE historial_asignacion_camion ADD CONSTRAINT uq_chofer_camion_historial UNIQUE (id_chofer, id_camion, fecha_asignacion);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_fecha_asignacion'
    ) THEN
        ALTER TABLE historial_asignacion_camion ADD CONSTRAINT chk_fecha_asignacion CHECK (fecha_asignacion <= CURRENT_DATE);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_fecha_baja'
    ) THEN
        ALTER TABLE historial_asignacion_camion ADD CONSTRAINT chk_fecha_baja CHECK (fecha_baja IS NULL OR fecha_baja >= fecha_asignacion);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'uq_placa_camion'
    ) THEN
        ALTER TABLE camion ADD CONSTRAINT uq_placa_camion UNIQUE (placa);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'fk_tipo_camion'
    ) THEN
        ALTER TABLE camion ADD CONSTRAINT fk_tipo_camion FOREIGN KEY (tipo_id) REFERENCES tipo_camion(id);
    END IF;
END $$;

DO $$
BEGIN
    -- Alerta Mantenimiento Constraints
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_camion_alerta_mantenimiento'
    ) THEN
        ALTER TABLE alerta_mantenimiento ADD CONSTRAINT fk_camion_alerta_mantenimiento FOREIGN KEY (camion_id) REFERENCES camion(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_tipo_mantenimiento_alerta'
    ) THEN
        ALTER TABLE alerta_mantenimiento ADD CONSTRAINT fk_tipo_mantenimiento_alerta FOREIGN KEY (tipo_mantenimiento_id) REFERENCES tipo_mantenimiento(tipo_mantenimiento_id);
    END IF;
END $$;

DO $$
BEGIN
    -- Registro Mantenimiento Constraints
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_camion_afectado'
    ) THEN
        ALTER TABLE registro_mantenimiento ADD CONSTRAINT fk_camion_afectado FOREIGN KEY (camion_id) REFERENCES camion(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_alerta_mantenimiento_registro'
    ) THEN
        ALTER TABLE registro_mantenimiento ADD CONSTRAINT fk_alerta_mantenimiento_registro FOREIGN KEY (alerta_id) REFERENCES alerta_mantenimiento(alerta_id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_kilometraje_mantenimiento'
    ) THEN
        ALTER TABLE registro_mantenimiento ADD CONSTRAINT chk_kilometraje_mantenimiento CHECK (kilometraje_mantenimiento >= 0);
    END IF;
END $$;

DO $$
BEGIN
    -- Ruta Camion Constraints
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_camion_asignado_ruta'
    ) THEN
        ALTER TABLE ruta_camion ADD CONSTRAINT fk_camion_asignado_ruta FOREIGN KEY (camion_id) REFERENCES camion(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_fecha_asignacion_ruta'
    ) THEN
        ALTER TABLE ruta_camion ADD CONSTRAINT chk_fecha_asignacion_ruta CHECK (fecha <= CURRENT_DATE);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_ruta_asignada'
    ) THEN
        ALTER TABLE ruta_camion ADD CONSTRAINT fk_ruta_asignada FOREIGN KEY (ruta_id) REFERENCES ruta(id);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_colonia_ruta'
    ) THEN
        ALTER TABLE ruta ADD CONSTRAINT fk_colonia_ruta FOREIGN KEY (colonia_id) REFERENCES colonia(colonia_id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_nombre_ruta'
    ) THEN
        ALTER TABLE ruta ADD CONSTRAINT chk_nombre_ruta CHECK (nombre <> '');
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'uq_nombre_ruta'
    ) THEN
        ALTER TABLE ruta ADD CONSTRAINT uq_nombre_ruta UNIQUE (nombre);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_ruta_id_punto_recoleccion'
    ) THEN
        ALTER TABLE punto_recoleccion ADD CONSTRAINT fk_ruta_id_punto_recoleccion FOREIGN KEY (ruta_id) REFERENCES ruta(id);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'uq_email_ciudadano'
    ) THEN
        ALTER TABLE ciudadano ADD CONSTRAINT uq_email_ciudadano UNIQUE (email);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'uq_alias_ciudadano'
    ) THEN
        ALTER TABLE ciudadano ADD CONSTRAINT uq_alias_ciudadano UNIQUE (alias);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_email_ciudadano'
    ) THEN
        ALTER TABLE ciudadano ADD CONSTRAINT chk_email_ciudadano CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_ciudadano_domicilio'
    ) THEN
        ALTER TABLE domicilio ADD CONSTRAINT fk_ciudadano_domicilio FOREIGN KEY (ciudadano_id) REFERENCES ciudadano(id);
    END IF;

    if NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_colonia_domicilio'
    ) THEN
        ALTER TABLE domicilio ADD CONSTRAINT fk_colonia_domicilio FOREIGN KEY (colonia_id) REFERENCES colonia(colonia_id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_alias_domicilio'
    ) THEN
        ALTER TABLE domicilio ADD CONSTRAINT chk_alias_domicilio CHECK (alias <> '');
    END IF;

END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_camion_estado'
    ) THEN
        ALTER TABLE estado_camion ADD CONSTRAINT fk_camion_estado FOREIGN KEY (camion_id) REFERENCES camion(id);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_relleno_vaciado'
    ) THEN
        ALTER TABLE registro_vaciado ADD CONSTRAINT fk_relleno_vaciado FOREIGN KEY (relleno_id) REFERENCES relleno_sanitario(relleno_id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_ruta_camion_vaciado'
    ) THEN
        ALTER TABLE registro_vaciado ADD CONSTRAINT fk_ruta_camion_vaciado FOREIGN KEY (ruta_camion_id) REFERENCES ruta_camion(ruta_camion_id);
    END IF;
END $$;

DO $$
BEGIN
    -- Anomalia Constraints
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_punto_anomalia'
    ) THEN
        ALTER TABLE anomalia ADD CONSTRAINT fk_punto_anomalia FOREIGN KEY (punto_id) REFERENCES punto_recoleccion(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_conductor_anomalia'
    ) THEN
        ALTER TABLE anomalia ADD CONSTRAINT fk_conductor_anomalia FOREIGN KEY (conductor_id) REFERENCES empleado(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_camion_anomalia'
    ) THEN
        ALTER TABLE anomalia ADD CONSTRAINT fk_camion_anomalia FOREIGN KEY (camion_id) REFERENCES camion(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_ruta_anomalia'
    ) THEN
        ALTER TABLE anomalia ADD CONSTRAINT fk_ruta_anomalia FOREIGN KEY (ruta_id) REFERENCES ruta(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints
        WHERE constraint_name = 'fk_referencia_anomalia'
    ) THEN
        ALTER TABLE anomalia ADD CONSTRAINT fk_referencia_anomalia FOREIGN KEY (anomalia_referencia_id) REFERENCES anomalia(anomalia_id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_tipo_anomalia'
    ) THEN
        ALTER TABLE anomalia ADD CONSTRAINT chk_tipo_anomalia CHECK (tipo_anomalia IN ('ANOMALIA', 'INCIDENCIA', 'REPORTE_CONDUCTOR', 'REPORTE_FALLA_CRITICA', 'SEGUIMIENTO_FALLA_CRITICA'));
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_estado_anomalia'
    ) THEN
        ALTER TABLE anomalia ADD CONSTRAINT chk_estado_anomalia CHECK (estado IS NULL OR estado IN ('PENDIENTE', 'EN_PROCESO', 'RESUELTA'));
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_descripcion_anomalia'
    ) THEN
        ALTER TABLE anomalia ADD CONSTRAINT chk_descripcion_anomalia CHECK (descripcion <> '');
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.check_constraints
        WHERE constraint_name = 'chk_fecha_resolucion_anomalia'
    ) THEN
        ALTER TABLE anomalia ADD CONSTRAINT chk_fecha_resolucion_anomalia CHECK (fecha_resolucion IS NULL OR fecha_resolucion >= fecha_reporte);
    END IF;
END $$;

