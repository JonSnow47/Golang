/*
 * Revision History:
 *     Initial: 2018/05/02        Chen Yanchen
 */

package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type httpError struct {
	Code    int
	Key     string `json:"error"`
	Message string `json:"message"`
}

func newHTTPError(code int, key, message string) *httpError {
	return &httpError{
		Code:    code,
		Key:     key,
		Message: message,
	}
}

func (e *httpError) Error() string {
	return e.Key + ": " + e.Message
}

// httpErrorHandler customize echo's HTTP error handler.
func httpErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		key  = "ServerError"
		msg  string
	)

	if he, ok := err.(*httpError); ok {
		code = he.Code
		key = he.Key
		msg = he.Message
	} else if config.Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, newHTTPError(code, key, msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
