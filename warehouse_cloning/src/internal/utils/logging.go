package utils

import (
    "fmt"
    "io"
    "log"
    "os"
    "time"
    "strings"
)

func Debug(message string, arguments ...interface{}){
    if strings.ToLower(Config.Logging.Level) == "debug" {
        if arguments != nil {
            message = fmt.Sprintf(message, arguments...)
        }
        log.Println(message)
    }
}

func Error(message string, arguments ...interface{}){
    if arguments != nil {
        message = fmt.Sprintf(message, arguments...)
    }
    log.Panicln(message)
}

func LoadLoggingConfig() {

    var outputs []io.Writer;

    if Config.Logging.Stdout {
        outputs = append(outputs, os.Stdout)
    }

    if Config.Logging.Dirname != "" {

        fileName        := fmt.Sprintf("%s/%s.log", Config.Logging.Dirname, time.Now().Format("2006-01-02"))
        fileOperations  := os.O_CREATE | os.O_APPEND | os.O_WRONLY

        var filePermissions os.FileMode
        filePermissions = 0644

        file, err := os.OpenFile(fileName, fileOperations, filePermissions)
        if err != nil {
           log.Panicf("[log.go] %s\n", err)
        }
        outputs = append(outputs, file)
    }

    log.SetOutput(io.MultiWriter(outputs...))
}
