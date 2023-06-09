package output

import (
    "github.com/MMazoni/most-used-features/internal/data"
)

type Generator interface {
    GenerateOutput( filePath string, data []data.MostAccessedFeatures) error
}
