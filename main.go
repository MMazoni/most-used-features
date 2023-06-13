package main

import (
    "fmt"
    "github.com/MMazoni/most-used-features/internal/data"
    "github.com/MMazoni/most-used-features/internal/input"
    "github.com/MMazoni/most-used-features/internal/output"
    "github.com/MMazoni/most-used-features/internal/search"
    "os"
    "path/filepath"
    "strings"
    "time"
)

func main() {
    dir, outputDir := input.GetInput()
    startTime := time.Now()

    fmt.Println(".")
    sheets := make([]data.MostAccessedFeatures, 0)
    timestampFilename := data.TimestampFilename{}
    err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }

        if !d.IsDir() && strings.Contains(path, "access_log") {

            file, err := os.Open(path)
            if err != nil {
                return err
            }

            fmt.Println("File processed:", path)
            sheets, timestampFilename, err = search.MostUsedFeatures(sheets, timestampFilename, file)
            defer file.Close()

            if err != nil {
                return err
            }
        }
        return nil
    })

    firstDate, lastDate := formatDate(timestampFilename)
    outputFile := fmt.Sprintf("%s/features%s-%s.csv", outputDir, firstDate, lastDate)
    csvOutput := output.CsvOutput{}
    err = csvOutput.GenerateOutput(outputFile, sheets)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    elapsedTime := time.Since(startTime)
    seconds := elapsedTime.Seconds()
    fmt.Println(".")
    fmt.Printf("CSV file created successfully in %.4f seconds\n", seconds)
}

func formatDate(timestamp data.TimestampFilename) (string, string) {
    layout := "20060102"
    firstDateString := timestamp.FirstHitDate.Format(layout)
    lastDateString := timestamp.LastHitDate.Format(layout)

    return firstDateString, lastDateString
}
