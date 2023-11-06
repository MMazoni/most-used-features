package output

import (
	"github.com/MMazoni/most-used-features/internal/data"
)

type Generator interface {
	GenerateMufOutput(filePath string, data []data.MostAccessedFeatures) error
	GenerateCsrfOutput(filePath string, data []data.MostCsrfErrors) error
}
