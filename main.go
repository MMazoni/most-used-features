package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "regexp"
    "encoding/csv"
)

type MostAccessedFeatures struct {
    Path string
    Method string
    Access int
}

func main() {
    filePath := "access.log"

    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    sheets := make([]MostAccessedFeatures, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        path, method := getPathAndMethodOfLogLine(line, "HTTP/1.1\"")

        if !isTheCorrectPath(path) {
            continue
        }

        found := false
        for i := range sheets {
            if sheets[i].Path == path && sheets[i].Method == method {
                sheets[i].Access++
                found = true
                break
            }
        }
        if !found {
            sheets = append(sheets, MostAccessedFeatures{
                Path: path,
                Method: method,
                Access: 1,
            })
        }
    }
    generateCsvFile(sheets)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func getPathAndMethodOfLogLine(line string, pattern string) (string, string) {
    index := strings.Index(line, pattern)
    if index == -1 {
        return "", ""
    }

    substring := line[:index]
    words := strings.Fields(substring)
    if len(words) > 1 {
        return formatPath(words[len(words)-1]), words[len(words)-2][1:]
    } else if len(words) == 1 {
        return "", words[0]
    }

    return "", ""
 }

func isTheCorrectPath(path string) bool {
    pattern := `/(fonts|js|css|assets|img|favicon)`
    regExp := regexp.MustCompile(pattern)
    if regExp.MatchString(path) {
        return false
    }
    return true
}

func formatPath(path string) string {
    formatted := strings.Split(path, "?")[0]

    re := regexp.MustCompile(`/\d+$`)
    if lastIndex := re.FindStringIndex(formatted); lastIndex != nil {
        return formatted[:lastIndex[0]]
    }
    return formatted
}

func generateCsvFile(data []MostAccessedFeatures) {
    file, err := os.Create("most-used-features.csv")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()
    csvWriter := csv.NewWriter(file)

    header := []string{"Path", "Method", "Hits"}
    csvWriter.Write(header)
    for _, d := range data {
        row := []string{d.Path, d.Method, fmt.Sprintf("%d", d.Access)}
        csvWriter.Write(row)
    }

    csvWriter.Flush()

    if err := csvWriter.Error(); err != nil {
        fmt.Println("Error writing CSV:", err)
        return
    }
    fmt.Println("CSV file created successfully")
}
