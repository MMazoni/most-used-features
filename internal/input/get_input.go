package input

import (
    "log"
    "os"
)

func GetInput() *os.File {
    filePath := "access.log"
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }

    return file
}