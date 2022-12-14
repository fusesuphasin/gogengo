package gotemplate

import "html/template"

var ACLTemplate = template.Must(template.New("").Parse(
`package middleware

import (
	"errors"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

// TODO : Refactor This
func CheckPermission(enforcer *casbin.Enforcer, page string) fiber.Handler {
	
	return func(c *fiber.Ctx) error {
		/* auth := session.GetAuth()

		role := "test" */

		ok, err := enforcer.Enforce("test", page, c.Method("POST,GET,PUT"))
		okManage, _ := enforcer.Enforce("test", page, "manage")
		if err != nil {
			c.Status(500)
			return c.JSON(response.ErrorResponse{
				Success: false,
				Message: err.Error(),
				// Error:   err,
			})
		}

		if okManage {
			return c.Next()
		}

		if !ok {
			errorForbidden := errors.New("unauthorized access")
			c.Status(403)
			return c.JSON(response.ErrorResponse{
				Success: false,
				Message: "Forbidden access",
				Error: errorForbidden,
			})
		}
		return c.Next()
	}
}
`))

