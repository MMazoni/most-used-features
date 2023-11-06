package search

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/MMazoni/most-used-features/internal/data"
)

func MostCsrfErrors(sheets []data.MostCsrfErrors, timestampObj data.TimestampFilename, file *os.File) ([]data.MostCsrfErrors, data.TimestampFilename, error) {
	scanner := bufio.NewScanner(file)
	allowedMethods := "GET POST PUT PATCH DELETE"

	for scanner.Scan() {
		line := scanner.Text()
		path, method, env := getCsrfWordsLog(line)

		if !strings.Contains(allowedMethods, method) {
			continue
		}
		controller, action := GetControllerAndActionFromPath(path)

		found := false
		for i, sheet := range sheets {
			if sheet.Path == path && sheet.Method == method {
				sheets[i].Error++
				found = true
				break
			}
		}
		if !found {
			sheets = append(sheets, data.MostCsrfErrors{
				Path:       path,
				Method:     method,
				Env:        env,
				Controller: controller,
				Action:     action,
				Error:      1,
			})
		}
	}

	err := scanner.Err()
	return sheets, timestampObj, err
}

func getCsrfWordsLog(line string) (string, string, string) {
	methRegex := regexp.MustCompile(`\[meth\s+([\w-]+)\]`)
	uriRegex := regexp.MustCompile(`\[uri\s+([^\]]*)\]`)
	envRegex := regexp.MustCompile(`\[env\s+([\w-]+)\]`)

	method := methRegex.FindStringSubmatch(line)[1]
	uriMatch := uriRegex.FindStringSubmatch(line)[1]

	env := envRegex.FindStringSubmatch(line)[1]

	return formatUri(uriMatch), method, env
}

func formatUri(uriMatch string) string {
	var uri string
	if len(uriMatch) > 0 {
		uri = uriMatch
		if strings.Contains(uri, "?") {
			uri = strings.Split(uri, "?")[0]
		}
	}
	pattern := regexp.MustCompile(`\d+$`)
	match := pattern.FindStringSubmatchIndex(uri)

	if match != nil {
		uri = uri[:match[0]]
	}

	return uri
}
