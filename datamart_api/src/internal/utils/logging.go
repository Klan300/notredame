package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Debug(message string, arguments ...interface{}) {
	if strings.ToLower(Config.Logging.Level) == "debug" {
		if arguments != nil {
			message = fmt.Sprintf(message, arguments...)
		}
		log.Println(message)
	}
}

func Error(message string, arguments ...interface{}) {
	if arguments != nil {
		message = fmt.Sprintf(message, arguments...)
	}
	log.Panicln(message)
}

func LoadLoggingConfig() {

	var outputs []io.Writer

	if Config.Logging.Stdout {
		outputs = append(outputs, os.Stdout)
	}

	if Config.Logging.Dirname != "" {

		rootDirname, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Panicf("[logging.go] %v\n", err)
		}

		logDirname := filepath.Join(rootDirname, Config.Logging.Dirname)
		var dirPermissions os.FileMode = 0755
		os.MkdirAll(logDirname, dirPermissions)

		fileName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))
		filePath := filepath.Join(logDirname, fileName)
		fileOperations := os.O_CREATE | os.O_APPEND | os.O_WRONLY

		var filePermissions os.FileMode = 0644
		file, err := os.OpenFile(filePath, fileOperations, filePermissions)

		if err != nil {
			log.Panicf("[logging.go] %v\n", err)
		}

		outputs = append(outputs, file)
	}

	log.SetOutput(io.MultiWriter(outputs...))

}
