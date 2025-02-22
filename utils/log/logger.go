package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	LogFile *os.File
	Logger  zerolog.Logger
)

func Init() error {
	err := initMultiWriter()
	if err != nil {
		return err
	}
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	return nil
}

func initMultiWriter() error {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", os.ModePerm)
		if err != nil {
			return err
		}
	}
	LogFile, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	multi := zerolog.MultiLevelWriter(LogFile, os.Stdout)
	Logger = zerolog.New(multi).With().Timestamp().Logger()
	return nil
}

func Cleanup() {
	if LogFile != nil {
		LogFile.Close()
	}
}
