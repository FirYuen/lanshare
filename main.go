package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)


var port int16 = 3335

const (
	Folder_temp   = "./tmp"
	Folder_static = "./static"
	Folder_buffer = "./buff"
)

func init() {
	ensureDir(Folder_static)
	ensureDir(Folder_temp)
	ensureDir(Folder_buffer)
}
//go:embed public
var FS embed.FS

var shareStr string
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//t,_:=template.ParseFS(fs,"public/*.html")
	templ := template.Must(template.New("").ParseFS(FS, "public/*.html"))
	r.SetHTMLTemplate(templ)
	//r.LoadHTMLGlob("public/*")
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFS("/tmp", http.Dir("./tmp"))


	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{})
	})
	r.POST("tempuploads", tempUpload)

	r.GET("/tempupload", func(c *gin.Context) {
		c.HTML(http.StatusOK,"upload.html",gin.H{
			"uploads":"tempuploads",
		})
	})

	r.GET("/staticupload", func(c *gin.Context) {
		 c.HTML(http.StatusOK,"upload.html",gin.H{
			"uploads":"staticuploads",
		})
	})
	r.POST("staticuploads", staticUpload)
	r.GET("/clip", func(c *gin.Context) {
		c.HTML(http.StatusOK,"clip.html",gin.H{"clipContent":shareStr})
	})
	r.POST("/clip", func(c *gin.Context) {
		shareStr= c.PostForm("clipContent")
		c.HTML(http.StatusOK,"clip.html",gin.H{"clipContent":shareStr})
	})
	fmt.Printf("Starting server on port %d", port)
	r.Run(fmt.Sprintf(":%d", port))

}
