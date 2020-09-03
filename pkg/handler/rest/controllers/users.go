package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oeoen/push-notifications/helper"
	"github.com/oeoen/push-notifications/pkg/notification"
)

func PutFCMToken(m notification.Manager, cn notification.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		application, username := c.Param("application"), c.Param("username")

		body := &notification.FCMToken{}
		if err := c.Bind(body); err != nil {
			cn.Logger().Error(err)
			return err
		}
		err := m.StorageManager().SetFCMToken(c.Request().Context(), application, username, *body)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusOK, helper.Response{
			Message: "update success",
		})
	}
}

func GetFCMToken(m notification.Manager, cn notification.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		application, username := c.Param("application"), c.Param("username")
		data, err := m.StorageManager().GetFCMToken(c.Request().Context(), application, username)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusOK, data)
	}
}
