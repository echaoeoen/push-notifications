package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/oeoen/push-notifications/helper/errorp"
)

func responseError(c echo.Context, err error) error {
	e, ok := err.(*errorp.NotificationError)
	if ok {
		return c.JSON(int(e.Status()), e)
	}
	return err
}
