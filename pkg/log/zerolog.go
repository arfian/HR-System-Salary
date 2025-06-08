package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitZeroLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	runLogFile, _ := os.OpenFile(
		"log/app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(runLogFile)

	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
}
