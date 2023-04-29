package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func cleanFile(filename string) {
	time.Sleep(10*time.Minute)
	err := os.Remove(filename)
    if err != nil {
        fmt.Printf("Failed to delete file: %v", err)
        return
    }
 
    fmt.Printf("File %s has been deleted\n", filename)
}

func download(c *gin.Context, filename string){
	file, err := os.Open(filename)
	if err != nil {
		c.String(http.StatusNotFound, "File not found.")
		return
	}
	defer file.Close()

	c.Writer.Header().Set("Content-Description", "File Transfer")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filename)
	c.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Expires", "0")

	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error.")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
	go cleanFile(filename)
}

func getFile(c *gin.Context){
	file, err := c.FormFile("file")
    if err != nil {
        c.JSON(409, gin.H {
			"error" : "couldn't get users file",
		})
    }

	fileID := uuid.New().String()
    fileName := fmt.Sprintf("%s.csv", fileID)
	fileName = "files/" + fileName
	err = c.SaveUploadedFile(file, fileName)
    if err != nil {
        c.JSON(409, gin.H {
			"error" : "couldn't save file",
		})
    }
	cmd := exec.Command("python3", "1.py", fileName)
    err = cmd.Run()

    if err != nil {
        c.JSON(409, gin.H {
			"error" : "couldn't process file",
		})
    }

	download(c, fileName)
}
