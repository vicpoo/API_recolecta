// helpers.go
package infrastructure

import (
	"fmt"
	"time"
)

func parseFecha(fechaStr string) (time.Time, error) {
	formatos := []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02",
	}
	
	for _, formato := range formatos {
		if fecha, err := time.Parse(formato, fechaStr); err == nil {
			return fecha, nil
		}
	}
	
	return time.Time{}, fmt.Errorf("formato de fecha no reconocido")
}