package main

import (
    "fmt"
    "log"
    "github.com/MMazoni/most-used-features/internal/input"
    "github.com/MMazoni/most-used-features/internal/output"
    "github.com/MMazoni/most-used-features/internal/search"
)

const outputFile = "csv/most-used-features.csv"

func main() {

    // open file - get input
    file := input.GetInput()

    // search line by line the most accessed features
    sheets, err := search.MostUsedFeatures(file)
    defer file.Close()
    
    if err != nil {
        log.Fatal(err)
    }

    // generate the output
    csvOutput := output.CsvOutput{}
    err = csvOutput.GenerateOutput(outputFile, sheets)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("CSV file created successfully")
}
