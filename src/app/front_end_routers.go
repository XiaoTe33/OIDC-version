package app

import (
	"net/http"

	"winter-examination/src/conf"
	"winter-examination/src/utils"

	"github.com/gin-gonic/gin"
)

func PageRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func PageLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func PageMain(c *gin.Context) {
	c.HTML(200, "main.html", nil)
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Next()
			return
		}
		if !utils.IsValidJWT(token) {
			c.Next()
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, conf.IP+conf.FrontEndPort+"/user/main")
		return
	}
}
