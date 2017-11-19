package main

import (
    "net/http"

    "time"
    "github.com/gin-gonic/gin"
)

type Token string
type Email string

type EmailTrack struct {
    Email     Email `json:"email" binding:"required"`
}

type UrlTrack struct {
    Title       string `json:"title" binding:"required"`
    Url         string `json:"url" binding:"required"`
    Date        time.Time `json:"date" binding:"required"`
}

type Emails map[Email][]Token
var emails = make(Emails)

type Urls  map[Token][]UrlTrack
var urls = make(Urls)

func main() {

    router := gin.Default()
    router.LoadHTMLGlob("./templates/*")

    router.Use(CORSMiddleware())

    // router.Use(favicon.New("./favicon.ico"))

    router.GET("/", index)
    router.GET("/:token", detail)

    v1 := router.Group("/api/v1/track")
    {
        v1.POST("/email/:guid", createTrackEmail)        
        v1.POST("/url/:guid", createTrackUrl) 
    }

    router.Run()
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"emails": emails,
    })
}

func detail(c *gin.Context) {
    guid := Token(c.Param("token"))

	c.HTML(http.StatusOK, "detail.tmpl", gin.H{
		"urls": urls[guid],
    })
}

func createTrackEmail(c *gin.Context) {
    guid := Token(c.Param("guid"))
    
    var json EmailTrack    
    c.Bind(&json) 
    
    insertTrackEmail(guid, json.Email)

    c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": guid, "resourceId": json.Email})
}

func createTrackUrl(c *gin.Context) { 
    guid := Token(c.Param("guid"))
    
    var json UrlTrack
    c.Bind(&json)

    insertTrackNewUrl(guid, json)

    c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": guid, "resourceId": json})
}

func insertTrackEmail(guid Token, email Email) {
    emails[email] = append(emails[email], guid)
}

func insertTrackNewUrl(guid Token, url UrlTrack) {
    urls[guid] = append(urls[guid], url)
}

func CORSMiddleware() gin.HandlerFunc {
     return func(c *gin.Context) {
         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
         c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding")

         if c.Request.Method == "OPTIONS" {
      
             c.AbortWithStatus(200)
         } else {
             c.Next()
         }
     }
 }
