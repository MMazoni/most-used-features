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

const outputFile = "csv/most-used-features.csv"

func main() {
    startTime := time.Now()
    dir := input.GetInput()

    fmt.Println(".")
    sheets := make([]data.MostAccessedFeatures, 0)
    err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }

        if !d.IsDir() {

            if !strings.HasSuffix(path, "access_log") {
                return filepath.SkipDir
            }

            file, err := os.Open(path)
            if err != nil {
                return err
            }

            fmt.Println("File processed:", path)
            sheets, err = search.MostUsedFeatures(sheets, file)
            defer file.Close()

            if err != nil {
                return err
            }
        }
        return nil
    })

    csvOutput := output.CsvOutput{}
    err = csvOutput.GenerateOutput(outputFile, sheets)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    elapsedTime := time.Since(startTime)
    seconds := elapsedTime.Seconds()
    fmt.Println(".")
    fmt.Printf("CSV file created successfully in %.2f seconds\n", seconds)
}
