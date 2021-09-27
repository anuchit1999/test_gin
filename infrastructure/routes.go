package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//GinRouter -> Gin Router
type GinRouter struct {
    Gin *gin.Engine
}

//NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

    httpRouter := gin.Default()
    httpRouter.Use(static.Serve("/", static.LocalFile("./templates", true)))
    httpRouter.LoadHTMLGlob("templates/*")
	httpRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", 
        gin.H{
			"title": "Main website",
		})
	})
    return GinRouter{
        Gin: httpRouter,
    }

}
