package main

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func processFile(c *gin.Context, file *multipart.FileHeader) (error) {
	err := c.SaveUploadedFile(file, "temp_file.csv")
    if err != nil {
        return err
    }

	return nil
}