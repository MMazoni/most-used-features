package output

import (
	"encoding/csv"
	"fmt"
	"github.com/MMazoni/most-used-features/internal/data"
	"os"
)

type CsvOutput struct{}

func (co CsvOutput) GenerateMufOutput(filePath string, data []data.MostAccessedFeatures) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	header := []string{"Path", "Method", "Controller", "Action", "Hits", "Errors"}
	csvWriter.Write(header)

	for _, d := range data {
		row := []string{d.Path, d.Method, d.Controller, d.Action, fmt.Sprintf("%d", d.Access), fmt.Sprintf("%d", d.Error)}
		csvWriter.Write(row)
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		return fmt.Errorf("error writing CSV: %w", err)
	}

	return nil
}

func (co CsvOutput) GenerateCsrfOutput(filePath string, data []data.MostCsrfErrors) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	header := []string{"Path", "Method", "Env", "Controller", "Action", "Errors"}
	csvWriter.Write(header)

	for _, d := range data {
		row := []string{d.Path, d.Method, d.Env, d.Controller, d.Action, fmt.Sprintf("%d", d.Error)}
		csvWriter.Write(row)
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		return fmt.Errorf("error writing CSV: %w", err)
	}

	return nil
}
