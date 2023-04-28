package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func admin(c *gin.Context){
	c.String(http.StatusOK,"Hello from admin page of trainhop")
}

func root(c *gin.Context){
	data, err := ioutil.ReadFile("frontend/index.html")
    if err != nil {
      c.Status(http.StatusNotFound)
      return
    }
    c.Writer.Header().Set("Content-Type", "text/html")
    c.Writer.Write(data)
}

func download(c *gin.Context){
	data, err := ioutil.ReadFile("frontend/download.html")
    if err != nil {
      c.Status(http.StatusNotFound)
      return
    }
    c.Writer.Header().Set("Content-Type", "text/html")
    c.Writer.Write(data)
}

func getFile(c *gin.Context){
	file, err := c.FormFile("file")
    if err != nil {
        c.JSON(409, gin.H {
			"error" : "couldn't get users file",
		})
    }

	err = c.SaveUploadedFile(file, "temp_file.csv")
    if err != nil {
        c.JSON(409, gin.H {
			"error" : "couldn't save file",
		})
    }
	cmd := exec.Command("python3", "1.py")
    err = cmd.Run()

    if err != nil {
        c.JSON(409, gin.H {
			"error" : "couldn't process file",
		})
    }
	download(c)
	// c.JSON(200, gin.H {
	// 	"success" : "file processed and saved",
	// })
}

func giveFile(c *gin.Context) {
	filename := "temp_file.csv"
	file, err := os.Open(filename)
  	if err != nil {
    	c.String(http.StatusNotFound, "File not found")
    	return
  	}
  	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error retrieving file info")
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+filename)
  	c.Header("Content-Type", "application/octet-stream")
  	c.Header("Content-Length", fmt.Sprint(fileInfo.Size()))
  	http.ServeContent(c.Writer, c.Request, filename, fileInfo.ModTime(), file)
}