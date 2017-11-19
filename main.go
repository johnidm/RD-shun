package main

import (
    "net/http"
    "flag"
    "github.com/gin-gonic/gin"
)

func main() {

    router := gin.Default()
    router.LoadHTMLGlob("./templates/*")

    router.GET("/", index)
    router.GET("/detail/:email", detail)

    v1 := router.Group("/api/v1/track")
    {
        v1.GET("/email/:guid", createTrackEmail)        
        v1.GET("/url/:guid", createTrackUrl) 
    }

    port := flag.String("port", "8000", "HTTP Port")
    router.Run(":" + *port)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
    })
}

func detail(c *gin.Context) {

    email := c.Param("email")

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": email,
    })
}

func createTrackEmail(c *gin.Context) { 
    guid := c.Param("guid")
    c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": guid})
}

func createTrackUrl(c *gin.Context) { 
    guid := c.Param("guid")
    c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": guid})
}