package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)
func fileServer(c *gin.Context){
	path:="./"
	fileName:=path+c.Query("name")
	c.File(fileName)
}


func upload(c *gin.Context, dest string) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]

	for _, file := range files {
		basename := filepath.Base(file.Filename)
		filename := filepath.Join(".", dest, basename)
		if _, err := os.Stat(filename); err == nil {
			os.Remove(filename)
		}
		if err := c.SaveUploadedFile(file, filename); err != nil {
			fmt.Println(err.Error())
			c.HTML(http.StatusOK, "upload_failed.html", gin.H{})
		}
		c.HTML(http.StatusOK, "upload_result.html", gin.H{})
		return
	}

}

func staticUpload(c *gin.Context) {
	upload(c, "static")
}
func tempUpload(c *gin.Context) {
	upload(c, "tmp")
}
