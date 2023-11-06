package commands

import (
	"fmt"
	"github.com/MMazoni/most-used-features/internal/data"
	"github.com/MMazoni/most-used-features/internal/input"
	"github.com/MMazoni/most-used-features/internal/output"
	"github.com/MMazoni/most-used-features/internal/search"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CsrfCommand(cmd *cobra.Command, args []string) {
	dir, outputDir := input.GetInput()
	startTime := time.Now()

	// log the errors
	logFile, err := os.Create("logs/csrf_error.log")
	if err != nil {
		fmt.Println("Failed to create log file:", err)
		return
	}
	defer logFile.Close()

	// Set the log output to the file
	log.SetOutput(logFile)

	defer func() {
		if r := recover(); r != nil {
			errMsg := fmt.Sprintf("Panic occurred: %v", r)
			log.Println(errMsg)
		}
	}()

	fmt.Println(".")
	sheets := make([]data.MostCsrfErrors, 0)
	timestampFilename := data.TimestampFilename{}
	err = filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.Contains(path, "csrf.log") {

			file, err := os.Open(path)
			if err != nil {
				return err
			}

			fmt.Println("File processed:", path)
			sheets, timestampFilename, err = search.MostCsrfErrors(sheets, timestampFilename, file)
			defer file.Close()

			if err != nil {
				return err
			}
		}
		return nil
	})

	//    firstDate, lastDate := formatDate(timestampFilename)
	outputFile := fmt.Sprintf("%s/features-%s.csv", outputDir, time.Now().Format("200601021504"))
	csvOutput := output.CsvOutput{}
	err = csvOutput.GenerateCsrfOutput(outputFile, sheets)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	elapsedTime := time.Since(startTime)
	seconds := elapsedTime.Seconds()
	fmt.Println(".")
	fmt.Printf("CSV file created successfully in %.4f seconds\n", seconds)
}
