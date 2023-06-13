package input

import (
    "bufio"
    "fmt"
    "strings"
    "os"
)

func GetInput() string {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter the directory log path: ")
    dirPath, _ := reader.ReadString('\n')
    dirPath = strings.TrimSpace(dirPath)

    return dirPath
}
