-- Agrega ciudadano_id a anomalia: hasta ahora la tabla solo registraba quien
-- reporto cuando era un conductor (conductor_id); un reporte de ciudadano
-- quedaba sin ningun dato que lo ligara de vuelta a quien lo mando. Sin esto
-- no hay forma de que un ciudadano liste ("mis reportes") ni elimine su
-- propio reporte -- ambos ya confirmados como requisito del proyecto.
--
-- Correr manualmente contra la BD que ya esta corriendo:
--
--   docker exec -i postgres_db psql -U <DB_USER> -d <DB_NAME> < migrations/2026-07-23_anomalia_ciudadano_id.sql
--
-- Es idempotente: se puede correr varias veces sin romper nada gracias a
-- IF NOT EXISTS.

ALTER TABLE anomalia
  ADD COLUMN IF NOT EXISTS ciudadano_id INTEGER DEFAULT NULL;

COMMENT ON COLUMN anomalia.ciudadano_id IS
  'Quien reporto, cuando fue un ciudadano (nunca junto con conductor_id: una anomalia la reporta un conductor o un ciudadano, no los dos). Se deriva del JWT al crear (ver CreateAnomaliaController.go), nunca se confia en lo que mande el cliente en el body salvo cuando quien crea es staff.';
