package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
)

type Csv struct {
	Path string
	Method string
	Access int
}

func main() {
	filePath := "access.log"
// 	notFeatures := ["fonts", "js", "css", "assets", "img", "favicon"]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sheets := make([]Csv, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		path, method := getPathAndMethodOfLogLine(line, "HTTP/1.1\"")

		if !isTheCorrectPath(path) {
			break
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
			sheets = append(sheets, Csv{
				Path: path,
				Method: method,
				Access: 1,
			})
		}
	}
	fmt.Println(sheets)

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
		return true
	}
	return false
}

func formatPath(path string) string {
	formatted := strings.Split(path, "?")[0]

	re := regexp.MustCompile(`/\d+$`)
	if lastIndex := re.FindStringIndex(formatted); lastIndex != nil {
		return formatted[:lastIndex[0]]
	}
	return formatted
}