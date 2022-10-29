package session

import (
	"gogengotest/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var sessionGet *session.Session

var dataAuth *SessionData

type SessionData struct {
	/* Auth *domain.User
	Role *domain.Role */
}

func InitSession(c *fiber.Ctx, user *service.RedisService, username string) {
	store := session.New()
	
	sess, err := store.Get(c)

	if err != nil {
		panic(err)
	}
	/* auth := user.CurrentUser((username)) */
	
	dataAuth = &SessionData{
		/* Auth: auth,
		Role: &auth.Role, */
	}

	sess.Set("username", username)

	sessionGet = sess

}

func GetSession() *session.Session {
	return sessionGet
}

func GetAuth() *SessionData {
	return dataAuth
}