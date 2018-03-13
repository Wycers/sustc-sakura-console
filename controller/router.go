package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/wycers/sustc-sakura-console/log"
	"github.com/wycers/sustc-sakura-console/util"
	"strings"
	"net/http"
	"path/filepath"
	"fmt"
	"os"
)

var logger = log.NewLogger(os.Stdout)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Headers","content-type")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			// c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}

func Routes() *gin.Engine {
	res := gin.New()
	res.Use(gin.Recovery())

	store := sessions.NewCookieStore([]byte(util.Config.SessionSecret))
	store.Options(sessions.Options{
		Path: "/",
		MaxAge:util.Config.SessionMaxAge,
		Secure:strings.HasPrefix(util.Config.Server, "https"),
		HttpOnly:true,
	})
	res.Use(sessions.Sessions("Sakura", store))

	if "dev" == util.Config.RuntimeMode {
		res.Use(Cors())
	}

	res.Static("/static", staticPath("static"))
	res.Static("/_nuxt", staticPath("static/dist/_nuxt"))
	res.StaticFile("/", staticPath("static/dist/index.html"))
	api := res.Group(util.PathAPI)
	api.POST("/login", LoginAction)
	res.GET("/download", DownloadAction)

	//static files

	res.NoRoute()
	return res
}

func staticPath(relativePath string) string {
	return filepath.ToSlash(filepath.Join(util.Config.StaticRoot, relativePath))
}
