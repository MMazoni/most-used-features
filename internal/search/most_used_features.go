package search

import (
    "bufio"
    "fmt"
    "github.com/MMazoni/most-used-features/internal/data"
    "os"
    "regexp"
    "strconv"
    "strings"
    "time"
    "unicode"
)

func MostUsedFeatures(sheets []data.MostAccessedFeatures, timestampObj data.TimestampFilename, file *os.File) ([]data.MostAccessedFeatures, data.TimestampFilename, error) {
    scanner := bufio.NewScanner(file)
    allowedMethods := "GET POST PUT PATCH DELETE"

    for scanner.Scan() {
        line := scanner.Text()
        line = strings.ReplaceAll(line, "//", "/")
        path, method, code, shouldSkip := getWordsOfLogLine(line, "HTTP/")

        if shouldSkip {
            continue
        }
        if !strings.Contains(allowedMethods, method) || !strings.Contains(path, "/") {
            continue
        }
        controller, action := getControllerAndActionFromPath(path)
//        timestampObj, _ = getTimestamp(date, timestampObj)

        found := false
        for i, sheet := range sheets {
            if sheet.Path == path && sheet.Method == method {
                sheets[i].Access++
                if code >= 400 && code < 600 {
                    sheets[i].Error++
                }
                found = true
                break
            }
        }
        if !found {
            errorCount := 0
            if code >= 400 && code < 600 {
                errorCount = 1
            }
            sheets = append(sheets, data.MostAccessedFeatures{
                Path:   path,
                Method: method,
                Controller: controller,
                Action: action,
                Access: 1,
                Error: errorCount,
            })
        }
    }

    err := scanner.Err()
    return sheets, timestampObj, err

}

func getWordsOfLogLine(line string, pattern string) (string, string, int, bool) {
    index := strings.Index(line, pattern)
    if index == -1 {
        return "", "", 0, true
    }

    beforeSubstring := line[:index]
    afterSubstring := line[index:]
    words := strings.Fields(beforeSubstring)
    path := formatPath(words[len(words)-1])
    if !isTheCorrectPath(path) {
        return "", "", 0, true
    }

//    date = words[len(words)-4][1:12]


    code, err := strconv.Atoi(strings.Fields(afterSubstring)[1])
    if err != nil {
        fmt.Println("Failed to convert string to int:", err)
        code = 0
    }
    if len(words) > 1 {
        return path, words[len(words)-2][1:], code, false
    } else if len(words) == 1 {
        return "", words[0], code, false
    }

    return "", "", 0, true
}

func isTheCorrectPath(path string) bool {
    prefixes := []string{
        "/fonts", "/js", "/css", "/assets", "/img", "/favicon", "/manifest.json",
        "/ads.txt", "/robot", "/image", "/apple-touch-icon", "/RepairQ-",
        "/.git", "/Core","/boaform", "/GponForm", "/_profiler", "/.env", "/system",
        "/HNAP1", "/client", "/upl", "/geoip", "/1.php", "/bundle", "/file",
        "/sqlmanager", "/db", "/php", "/mysql", "/sql", "/admin", "/_phpmyadmin",
        "/phpMyAdmin", "/MyAdmin", "/administrator", "/PMA", "/1php", "/pma",
        "/wp", "/program", "/vendor", "/geoserver", "/hudson", "/boaform",
        "/cgi-bin", "/.git", "/Telerik", "/gate", "/debug", "/sitemap", "/live",
        "/back", "/dev", "/core", "/source", "/rest", "/script", "/laravel",
        "/shared", "/private", "/app", "/env", "/docker", "/cp", "/cms",
        "/local", "/front", "/config", "/video", "http", "/dvr", "/axis",
        "/cn", "/druid", "/old", "/aws", "/blogs", "/v2", "/s/", "/ecp/",
        "/telescope", "/mgmt", "/sendgrid", "/manage", "/doc", "/owa", "/manager",
        "/metrics", "/conf", "/library", "/audio", "/storage", "/base",
        "/protected", "/newsite", "/www", "/sites", "/database", "/ec2",
        "/muieblackcat", "/shell", "/dashboard", "/download", "/supp",
        "/root", "/test", "/temp", "/tools", "/server", "/.docker",
        "/.s3", "/.vscode", "/alpha", "/beta", "/bootstrap", "/demo",
        "/home", "/manual", "/services", "/apache", "/inf", "/deploy",
        "/forum", "/console", "/web", "/File", "/channel", "/sys",
        "/.jupyter", "/twitter", "/acme", "/anaconda", "/agora",
        "/babel", "/backend", "/back-end", "/backup", "/blob", "bookchain",
        "/blue", "/box", "/build", "/cardea", "/cron", "/dataset", "/custom",
        "/delivery", "/dist", "/django", "/example", "/e2e", "/engine",
        "/dotfiles", "/favs", "/exapi", "/fastlane", "/Final", "/final",
        "/fixture", "/flask", "/gists", "/html", "/icon", "/src", "/static",
        "/style", "/stat", "/theme", "/unsplash", "/unix", "/ubuntu",
        "/vue", "/symfony", "/cred", "/linux", "/node", "/ops", "/picture",
        "/prisma", "/public", "/Socket", "/var", "/.wp",
    }
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

func getControllerAndActionFromPath(path string) (string, string) {
    var controller = ""
    if path == "/" || path == "//" {
        return "Index", "Index"
    }

    parts := strings.Split(path[1:], "/")

    if strings.HasPrefix(path, "/api/") || strings.HasPrefix(path, "/ajax/") {
        controller = strings.Join([]string{capitalize(parts[0]), " ", capitalize(parts[1]) }, "")
        if len(parts) == 2 {
            return controller, "Index"
        }
        return controller, capitalize(parts[2])
    }

    controller = capitalize(parts[0])

    if len(parts) == 1 {
        return controller, "Index"
    }
    return controller, capitalize(parts[1])
}

func capitalize(str string) string {
    if len(str) < 2 {
        return ""
    }
    runes := []rune(str)
    runes[0] = unicode.ToUpper(runes[0])
    return string(runes)
}

func getTimestamp(dateString string, timestamp data.TimestampFilename) (data.TimestampFilename, error) {
    parsedTime, err := time.Parse("02/Jan/2006", dateString)
    if err != nil {
        fmt.Println("Failed to parse the date:", err)
        return data.TimestampFilename{}, err
    }
    if timestamp.FirstHitDate.IsZero() && timestamp.LastHitDate.IsZero() {
        timestamp.FirstHitDate = parsedTime
        timestamp.LastHitDate = parsedTime
        return timestamp, nil
    }
    if parsedTime.Before(timestamp.FirstHitDate) {
        timestamp.FirstHitDate = parsedTime
    }
    if parsedTime.After(timestamp.LastHitDate) {
        timestamp.LastHitDate = parsedTime
    }

    return timestamp, nil
}
