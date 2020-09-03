package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oeoen/push-notifications/helper"
	"github.com/oeoen/push-notifications/pkg/notification"
)

var NotificationParam = map[string]string{
	"application": "application",
	"username":    "username",
	"lastId":      "last_id",
	"id":          "id",
}

func GetNotificationNext(m notification.Manager, cn notification.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		application, username, lastId := c.Param(NotificationParam["application"]), c.Param(NotificationParam["username"]), c.Param(NotificationParam["lastId"])
		notifications, err := m.StorageManager().FetchNotification(c.Request().Context(), application, username, [3]string{"id", ">", lastId})
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusOK, notifications)
	}
}

func GetNewestNotification(m notification.Manager, cn notification.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		application, username, lastId := c.Param(NotificationParam["application"]), c.Param(NotificationParam["username"]), c.Param(NotificationParam["lastId"])
		var par = [][3]string{}
		if lastId != "" {
			par = append(par, [3]string{"id", ">", lastId})
		}
		notifications, err := m.StorageManager().FetchNotification(c.Request().Context(), application, username, par...)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusOK, notifications)
	}
}

func SendNotification(m notification.Manager, cn notification.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		application, username := c.Param(NotificationParam["application"]), c.Param(NotificationParam["username"])
		content := &notification.Content{}
		if err := c.Bind(content); err != nil {
			cn.Logger().Error(err)
			return err
		}
		err := m.SendNotification(c.Request().Context(), application, username, *content)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusAccepted, helper.Response{
			Message: "Notification sended",
		})
	}
}

func ReadNotification(m notification.Manager, cn notification.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		application, username, id := c.Param(NotificationParam["application"]), c.Param(NotificationParam["username"]), c.Param(NotificationParam["id"])

		err := m.StorageManager().ReadNotification(c.Request().Context(), application, username, id)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusAccepted, helper.Response{
			Message: "Notification sended",
		})
	}
}
