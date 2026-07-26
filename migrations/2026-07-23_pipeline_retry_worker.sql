-- Agrega el contador de reintentos que usa el PipelineRetryWorker (ver
-- src/Fallas/infrastructure/pipeline_retry_worker.go) para reintentar de
-- forma acotada el pipeline modelo_reportes -> clasificador_reportes cuando
-- una anomalia quedo en estado_pipeline = 'error' (microservicio caido,
-- timeout, etc.) o 'procesando' abandonado (el proceso que la tomo se
-- reinicio a la mitad, p. ej. Air en dev).
--
-- Correr manualmente contra la BD que ya esta corriendo:
--
--   docker exec -i postgres_db psql -U <DB_USER> -d <DB_NAME> < migrations/2026-07-23_pipeline_retry_worker.sql
--
-- Es idempotente: se puede correr varias veces sin romper nada gracias a
-- IF NOT EXISTS.

ALTER TABLE anomalia
  ADD COLUMN IF NOT EXISTS pipeline_intentos INTEGER NOT NULL DEFAULT 0;

COMMENT ON COLUMN anomalia.pipeline_intentos IS
  'Cuantas veces se intento correr el pipeline sobre esta anomalia. Al llegar a MaxIntentosPipeline (application.ProcesarPipelineAnomaliaUseCase) el worker deja de reintentarla sola y queda para revision manual (visible via pipeline_error en /api/anomalias).';

COMMENT ON COLUMN anomalia.estado_pipeline IS
  'pendiente | procesando (reclamada, en vuelo) | rechazado (fraude segun modelo_reportes) | clasificado (paso por clasificador_reportes) | error (algun microservicio fallo/timeout, reintentable hasta pipeline_intentos < MaxIntentosPipeline)';
