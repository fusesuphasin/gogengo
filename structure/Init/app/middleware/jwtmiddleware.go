package middleware

import (
	"gogengotest/app/service"
	"gogengotest/app/utils/jwt"
	"gogengotest/app/utils/session"

	"github.com/gofiber/fiber/v2"
)

// TODO : Refactor This
func JWTProtected(svc service.RedisService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// store := session.New()
		// sess, _ := store.Get(c)
		//token1, _ := c.Locals("my_token").(string)
		token, errExtract := jwt.ExtractTokenMetadata(c)

		if errExtract != nil {
			c.Status(401)
			
			return c.JSON(fiber.Map{"error": "Unauthorized access"})
		}

	    userName, user, _ := jwt.FetchAuth(svc, token)
		if userName == "" {
			c.Status(401)
			return c.JSON(fiber.Map{"error": "Unauthorized accesst"})
		}
		session.InitSession(c, &svc, user) 

		//sess.Set("username", userName)
		
		return c.Next()

	}
}