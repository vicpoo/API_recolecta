// tipo_anomalia.go
package entities

// TipoAnomalia es un alias de string: permite usar constantes con nombre
// (enum) sin cambiar el tipo real de la columna tipo_anomalia en BD.
type TipoAnomalia = string

const (
	TipoAnomaliaGeneral                 TipoAnomalia = "ANOMALIA"
	TipoAnomaliaIncidencia              TipoAnomalia = "INCIDENCIA"
	TipoAnomaliaReporteConductor        TipoAnomalia = "REPORTE_CONDUCTOR"
	TipoAnomaliaReporteFallaCritica     TipoAnomalia = "REPORTE_FALLA_CRITICA"
	TipoAnomaliaSeguimientoFallaCritica TipoAnomalia = "SEGUIMIENTO_FALLA_CRITICA"
)

// TiposAnomaliaValidos contiene todos los valores permitidos para tipo_anomalia.
var TiposAnomaliaValidos = []TipoAnomalia{
	TipoAnomaliaGeneral,
	TipoAnomaliaIncidencia,
	TipoAnomaliaReporteConductor,
	TipoAnomaliaReporteFallaCritica,
	TipoAnomaliaSeguimientoFallaCritica,
}

// EsTipoAnomaliaValido verifica si el tipo pertenece al enum soportado.
func EsTipoAnomaliaValido(tipo string) bool {
	for _, t := range TiposAnomaliaValidos {
		if t == tipo {
			return true
		}
	}
	return false
}
