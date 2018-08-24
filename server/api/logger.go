package api

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// NewLoggerConverter returns an Echo logger converter to Logrus.
func NewLoggerConverter() echo.Logger { return loggerConverter{} }

type loggerConverter struct{}

func (l loggerConverter) SetOutput(out io.Writer) { logrus.SetOutput(out) }
func (l loggerConverter) Output() io.Writer       { return os.Stdout }
func (l loggerConverter) SetLevel(level log.Lvl) {
	switch level {
	case log.DEBUG:
		logrus.SetLevel(logrus.DebugLevel)
	case log.INFO:
		logrus.SetLevel(logrus.InfoLevel)
	case log.WARN:
		logrus.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		logrus.SetLevel(logrus.ErrorLevel)
	}
}
func (l loggerConverter) Level() log.Lvl                       { return log.DEBUG }
func (l loggerConverter) SetPrefix(string)                     {}
func (l loggerConverter) Prefix() string                       { return "" }
func (l loggerConverter) Print(args ...interface{})            { logrus.Print(args...) }
func (l loggerConverter) Printf(s string, args ...interface{}) { logrus.Printf(s, args...) }
func (l loggerConverter) Printj(arg log.JSON)                  { logrus.Print(arg) }
func (l loggerConverter) Debug(args ...interface{})            { logrus.Debug(args...) }
func (l loggerConverter) Debugf(s string, args ...interface{}) { logrus.Debugf(s, args...) }
func (l loggerConverter) Debugj(arg log.JSON)                  { logrus.Debugln(arg) }
func (l loggerConverter) Info(args ...interface{})             { logrus.Info(args...) }
func (l loggerConverter) Infof(s string, args ...interface{})  { logrus.Infof(s, args...) }
func (l loggerConverter) Infoj(arg log.JSON)                   { logrus.Infoln(arg) }
func (l loggerConverter) Warn(args ...interface{})             { logrus.Warn(args...) }
func (l loggerConverter) Warnf(s string, args ...interface{})  { logrus.Warnf(s, args...) }
func (l loggerConverter) Warnj(arg log.JSON)                   { logrus.Warnln(arg) }
func (l loggerConverter) Error(args ...interface{})            { logrus.Error(args...) }
func (l loggerConverter) Errorf(s string, args ...interface{}) { logrus.Errorf(s, args...) }
func (l loggerConverter) Errorj(arg log.JSON)                  { logrus.Errorln(arg) }
func (l loggerConverter) Fatal(args ...interface{})            { logrus.Fatal(args...) }
func (l loggerConverter) Fatalj(arg log.JSON)                  { logrus.Fatalln(arg) }
func (l loggerConverter) Fatalf(s string, args ...interface{}) { logrus.Fatalf(s, args...) }
func (l loggerConverter) Panic(args ...interface{})            { logrus.Panic(args...) }
func (l loggerConverter) Panicj(arg log.JSON)                  { logrus.Panicln(arg) }
func (l loggerConverter) Panicf(s string, args ...interface{}) { logrus.Panicf(s, args...) }

// RequestLogger middleware for Echo.
func RequestLogger(dev bool) echo.MiddlewareFunc {
	co := color.New()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req, res := c.Request(), c.Response()

			// execute next handlers and calculate latency
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			latency := stop.Sub(start)

			fields := logrus.Fields{
				"remote_ip":     req.RemoteAddr,
				"latency_human": latency.String(),
			}

			// if an error uuid exists in context add it to fields
			errorUUID := c.Get("error_uuid")
			if errorUUID != nil {
				fields["error_uuid"] = errorUUID
			}

			status, method, uri := res.Status, req.Method, req.URL
			if !dev {
				fields["uri"] = uri
				fields["method"] = method
				fields["user_agent"] = req.UserAgent()

				// calculate bytes in
				if cl := req.Header.Get(echo.HeaderContentLength); cl != "" {
					b, _ := strconv.ParseInt(cl, 10, 64)
					fields["bytes_in"] = b
				}

				fields["bytes_out"] = res.Size
				fields["latency"] = latency.Nanoseconds() / 1000
				fields["referer"] = req.Referer()
				fields["host"] = req.Host
				fields["path"] = c.Path()
				fields["status"] = status
			}

			// create colored status for message
			s := co.Green(status)
			switch {
			case status >= 500:
				s = co.Red(status)
			case status >= 400:
				s = co.Yellow(status)
			case status >= 300:
				s = co.Cyan(status)
			}

			// log info of the request
			if status == 500 {
				logrus.WithFields(fields).Error(fmt.Sprintf("%s [%s] %s", method, s, uri))
			} else {
				logrus.WithFields(fields).Info(fmt.Sprintf("%s [%s] %s", method, s, uri))
			}

			return nil
		}
	}
}
