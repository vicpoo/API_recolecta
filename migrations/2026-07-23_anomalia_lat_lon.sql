-- Agrega coordenadas GPS a la tabla anomalia. Insumo para el algoritmo
-- genetico de rutas (AG, todavia no implementado): cuando exista, necesita
-- saber DONDE ocurrio el bloqueo/incidente para decidir que arista del grafo
-- de rutas modificar (block_edge / inflate_weight, ver accion_sugerida).
--
-- No basta con punto_id: muchos reportes (ej. "calle bloqueada" de un
-- ciudadano o conductor) ocurren en un punto arbitrario de una ruta, no
-- necesariamente en un punto_recoleccion ya registrado. Por eso son
-- columnas propias en anomalia, no una referencia a otra tabla.
--
-- Mismo nombre/tipo que usa Rutas.PuntoRecoleccion (lat/lon) para mantener
-- la convencion ya establecida en el resto del backend. Nullable: todavia
-- ningun cliente (web/app) captura y envia ubicacion.
--
-- Correr manualmente contra la BD que ya esta corriendo:
--
--   docker exec -i postgres_db psql -U <DB_USER> -d <DB_NAME> < migrations/2026-07-23_anomalia_lat_lon.sql
--
-- Es idempotente: se puede correr varias veces sin romper nada gracias a
-- IF NOT EXISTS.

ALTER TABLE anomalia
  ADD COLUMN IF NOT EXISTS lat DOUBLE PRECISION,
  ADD COLUMN IF NOT EXISTS lon DOUBLE PRECISION;

COMMENT ON COLUMN anomalia.lat IS
  'Latitud de donde ocurrio la anomalia. Nullable hasta que los clientes capturen ubicacion. Insumo futuro del algoritmo genetico de rutas.';
COMMENT ON COLUMN anomalia.lon IS
  'Longitud de donde ocurrio la anomalia. Ver comentario de anomalia.lat.';
