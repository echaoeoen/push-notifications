package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oeoen/push-notifications/helper"
	"github.com/oeoen/push-notifications/pkg/police"
)

func UpsertPolicyRole(m police.Manager, cn police.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		body := &police.RBAC{}
		if err := c.Bind(body); err != nil {
			cn.Logger().Error(err)
			return err
		}
		tenant := parseParam(c, "tenant")
		body.Tenant = tenant
		if err := m.UpsertRole(c.Request().Context(), body); err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusCreated, body)
	}
}

func GetPolicyRole(m police.Manager, cn police.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		tenant, policy := parseParam(c, "tenant"), parseParam(c, "policy")
		acl, err := m.GetRolePolicy(c.Request().Context(), tenant, policy)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusOK, acl)
	}
}

func GetSubjectRoles(m police.Manager, cn police.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		tenant, subject := parseParam(c, "tenant"), parseParam(c, "subject")
		acl, err := m.GetSubjectRoles(c.Request().Context(), tenant, subject)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusOK, acl)
	}
}

func GetRoleSubjects(m police.Manager, cn police.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		tenant, policy := parseParam(c, "tenant"), parseParam(c, "policy")
		acl, err := m.GetRoleSubjects(c.Request().Context(), tenant, policy)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusOK, acl)
	}
}

func DeleteSubjectsRole(m police.Manager, cn police.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		tenant, policy, subject := parseParam(c, "tenant"), parseParam(c, "policy"), parseParam(c, "subject")
		err := m.DeleteRole(c.Request().Context(), tenant, subject, policy)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusAccepted, helper.Response{Message: "Deleted"})
	}
}

func FetchRoles(m police.Manager, cn police.Configuration) echo.HandlerFunc {
	return func(c echo.Context) error {
		body := &police.ACL{}
		queries := getQueries(c.QueryParams())
		if err := c.Bind(body); err != nil {
			cn.Logger().Error(err)
			return err
		}
		acls, err := m.FetchRoles(c.Request().Context(), queries...)
		if err != nil {
			cn.Logger().Error(err)
			return responseError(c, err)
		}
		return c.JSON(http.StatusOK, acls)
	}
}
