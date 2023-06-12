package input

import (
    "bufio"
    "fmt"
    "log"
    "strings"
    "os"
)

func GetInput() *os.File {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter the log file path: ")
    filePath, _ := reader.ReadString('\n')
    filePath = strings.TrimSpace(filePath)

    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }

    return file
}
