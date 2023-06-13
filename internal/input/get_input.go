package input

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func GetInput() (string, string) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter the directory log path(default: /var/log/httpd): ")
    inputPath, _ := reader.ReadString('\n')
    inputPath = strings.TrimSpace(inputPath)
    if inputPath == "" {
        inputPath = "/var/log/httpd"
    }

    fmt.Print("Enter the output path(default: /tmp): ")
    outputPath, _ := reader.ReadString('\n')
    outputPath = strings.TrimSpace(outputPath)
    if outputPath == "" {
        outputPath = "/var/log/httpd"
    }

    return inputPath, outputPath
}
