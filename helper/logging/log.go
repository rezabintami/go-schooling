package logging

import (
	"flag"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Logger struct {
	Log *log.Logger
}

func NewLogger() Logger {
	logger := log.New()
	logger.Out = os.Stdout

	return Logger{
		Log: logger,
	}
}

func Logging(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	f, err := os.OpenFile("api.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("cannot open 'testlogfile', (%s)", err.Error())
		flag.Usage()
		os.Exit(-1)
	}
	log.SetOutput(f)

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func LogEntry(c echo.Context, request, response interface{}) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	f, err := os.OpenFile("api.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("cannot open 'testlogfile', (%s)", err.Error())
		flag.Usage()
		os.Exit(-1)
	}
	log.SetOutput(f)

	return log.WithFields(log.Fields{
		"at":       time.Now().Format("2006-01-02 15:04:05"),
		"method":   c.Request().Method,
		"uri":      c.Request().URL.String(),
		"ip":       c.Request().RemoteAddr,
		"request":  request,
		"response": response,
	})
}
