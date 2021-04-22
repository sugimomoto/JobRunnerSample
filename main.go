package main

import (
	"fmt"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

type JobSample struct {
	samplestring string
}

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 5s", JobSample{
		samplestring: "Hello",
	})

	gin.SetMode(gin.ReleaseMode)

	routes := gin.Default()
	routes.GET("/jobrunner/status", JobJSON)
	// routes.LoadHTMLFiles("../github.com/bamzi/jobrunner/views/Status.html")
	// routes.GET("/jobrunner/html", JobHtml)

	routes.Run(":8881")
}

func JobJSON(c *gin.Context) {
	c.JSON(200, jobrunner.StatusJson())
}

func JobHtml(c *gin.Context) {

}

func (e JobSample) Run() {
	fmt.Println(e.samplestring + " World")
}
