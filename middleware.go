package elogrus

import (
	"io"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// Logrus : implement Logger
type Logrus struct {
	*logrus.Logger
}

// Logger ...
//var Logger *logrus.Logger

// GetEchoLogger for e.Logger
func GetEchoLogger(Logger *logrus.Logger) Logrus {
	return Logrus{Logger}
}

// Level returns logger level
func (l Logrus) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	default:
		l.Panic("Invalid level")
	}

	return log.OFF
}

// SetHeader is a stub to satisfy interface
// It's controlled by Logger
func (l Logrus) SetHeader(_ string) {}

// SetPrefix It's controlled by Logger
func (l Logrus) SetPrefix(s string) {}

// Prefix It's controlled by Logger
func (l Logrus) Prefix() string {
	return ""
}

// SetLevel set level to logger from given log.Lvl
func (l Logrus) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		l.Logger.SetLevel(logrus.DebugLevel)
	case log.WARN:
		l.Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		l.Logger.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		l.Logger.SetLevel(logrus.InfoLevel)
	default:
		l.Panic("Invalid level")
	}
}

// Output logger output func
func (l Logrus) Output() io.Writer {
	return l.Out
}

// SetOutput change output, default os.Stdout
func (l Logrus) SetOutput(w io.Writer) {
	l.Logger.SetOutput(w)
}

// Printj print json log
func (l Logrus) Printj(j log.JSON) {
	msg := extractMsg(j)
	l.Logger.WithFields(logrus.Fields(j)).Print(msg)
}

// Debugj debug json log
func (l Logrus) Debugj(j log.JSON) {
	msg := extractMsg(j)
	l.Logger.WithFields(logrus.Fields(j)).Debug(msg)
}

// Infoj info json log
func (l Logrus) Infoj(j log.JSON) {
	msg := extractMsg(j)
	l.Logger.WithFields(logrus.Fields(j)).Info(msg)
}

// Warnj warning json log
func (l Logrus) Warnj(j log.JSON) {
	msg := extractMsg(j)
	l.Logger.WithFields(logrus.Fields(j)).Warn(msg)
}

// Errorj error json log
func (l Logrus) Errorj(j log.JSON) {
	msg := extractMsg(j)
	l.Logger.WithFields(logrus.Fields(j)).Error(msg)
}

// Fatalj fatal json log
func (l Logrus) Fatalj(j log.JSON) {
	msg := extractMsg(j)
	l.Logger.WithFields(logrus.Fields(j)).Fatal(msg)
}

// Panicj panic json log
func (l Logrus) Panicj(j log.JSON) {
	msg := extractMsg(j)
	l.Logger.WithFields(logrus.Fields(j)).Panic(msg)
}

// Print string log
func (l Logrus) Print(i ...interface{}) {
	l.Logger.Print(i...)
}

// Debug string log
func (l Logrus) Debug(i ...interface{}) {
	l.Logger.Debug(i...)
}

// Info string log
func (l Logrus) Info(i ...interface{}) {
	l.Logger.Info(i...)
}

// Warn string log
func (l Logrus) Warn(i ...interface{}) {
	l.Logger.Warn(i...)
}

// Error string log
func (l Logrus) Error(i ...interface{}) {
	l.Logger.Error(i...)
}

// Fatal string log
func (l Logrus) Fatal(i ...interface{}) {
	l.Logger.Fatal(i...)
}

// Panic string log
func (l Logrus) Panic(i ...interface{}) {
	l.Logger.Panic(i...)
}

func logrusMiddlewareHandler(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	if err := next(c); err != nil {
		c.Error(err)
	}
	stop := time.Now()

	p := req.URL.Path

	bytesIn := req.Header.Get(echo.HeaderContentLength)

	c.Logger().Infoj(map[string]interface{}{
		"time_rfc3339":  time.Now().Format(time.RFC3339),
		"remote_ip":     c.RealIP(),
		"host":          req.Host,
		"uri":           req.RequestURI,
		"method":        req.Method,
		"path":          p,
		"referer":       req.Referer(),
		"user_agent":    req.UserAgent(),
		"status":        res.Status,
		"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		"latency_human": stop.Sub(start).String(),
		"bytes_in":      bytesIn,
		"bytes_out":     strconv.FormatInt(res.Size, 10),
		Msg:             "Handled request",
	})

	return nil
}

func logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return logrusMiddlewareHandler(c, next)
	}
}

// Hook is a function to process middleware.
func Hook() echo.MiddlewareFunc {
	return logger
}
