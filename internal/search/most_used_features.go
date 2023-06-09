package search

import (
    "bufio"
    "github.com/MMazoni/most-used-features/internal/data"
    "os"
    "regexp"
    "strings"
)

func MostUsedFeatures(file *os.File) ([]data.MostAccessedFeatures, error) {
    sheets := make([]data.MostAccessedFeatures, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        path, method := getPathAndMethodOfLogLine(line, "HTTP/1.1\"")

        if !isTheCorrectPath(path) {
            continue
        }

        found := false
        for i, sheet := range sheets {
            if sheet.Path == path && sheet.Method == method {
                sheets[i].Access++
                found = true
                break
            }
        }
        if !found {
            sheets = append(sheets, data.MostAccessedFeatures{
                Path:   path,
                Method: method,
                Access: 1,
            })
        }
    }

    err := scanner.Err()
    return sheets, err

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
    prefixes := []string{"/fonts", "/js", "/css", "/assets", "/img", "/favicon"}
    for _, prefix := range prefixes {
        if strings.HasPrefix(path, prefix) {
            return false
        }
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
