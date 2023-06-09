package input

import (
    "log"
    "os"
)

const filePath = "/var/log/httpd/sbx-us_access_log"

func GetInput() *os.File {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }

    return file
}