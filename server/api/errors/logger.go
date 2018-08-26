package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type data struct {
	UUID    string `json:"uuid"`
	Message string `json:"message"`
}

func newDataFromError(err Error) data {
	return data{
		UUID:    uuid.NewV4().String(),
		Message: err.Error(),
	}
}

// Middleware catch err from handlers or middlewares.
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				var err error
				switch r := r.(type) {
				case error:
					err = r
				default:
					err = fmt.Errorf("%v", r)
				}

				errorUUID := uuid.NewV4().String()
				c.Set("error_uuid", errorUUID)

				logrus.WithField("error_uuid", errorUUID).
					Errorf("%+v", errors.WithStack(err))
				c.JSON(http.StatusInternalServerError, data{
					UUID:    errorUUID,
					Message: http.StatusText(http.StatusInternalServerError),
				})
			}
		}()

		if err := next(c); err != nil {
			if ferr, ok := err.(Error); ok {
				data := newDataFromError(ferr)
				logrus.WithField("error_uuid", data.UUID).
					Warnf("%+v", ferr.SourceError)
				c.Set("error_uuid", data.UUID)

				switch ferr.Type {
				case Data:
					return c.JSON(http.StatusBadRequest, data)
				case Action:
					return c.JSON(http.StatusForbidden, data)
				case NotFound:
					return c.JSON(http.StatusNotFound, data)
				case Remote:
					return c.JSON(http.StatusServiceUnavailable, data)
				case Authorization:
					return c.JSON(http.StatusUnauthorized, data)
				}
			}

			errorUUID := uuid.NewV4().String()
			c.Set("error_uuid", errorUUID)

			if ferr, ok := err.(*echo.HTTPError); ok {
				e := logrus.WithField("error_uuid", errorUUID)
				if ferr.Code == 500 {
					e.Errorf("%+v", errors.WithStack(err))
				} else {
					e.Warnf("%+v", errors.WithStack(err))
				}
				return c.JSON(ferr.Code, data{
					UUID:    errorUUID,
					Message: fmt.Sprintf("%v", ferr.Message),
				})
			}

			logrus.WithField("error_uuid", errorUUID).Errorf("%+v", err)
			return c.JSON(http.StatusInternalServerError, data{
				UUID:    errorUUID,
				Message: http.StatusText(http.StatusInternalServerError),
			})
		}
		return nil
	}
}
