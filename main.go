package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func uploadFile(c *gin.Context) {
  file, err := c.FormFile("file");
  if err != nil {
    c.String(http.StatusBadRequest, "error retrieving file: %s",err.Error());
    return;
  }
  uploadDir := "./uploads";
  if _, err :=os.Stat(uploadDir); os.IsNotExist(err) {
    os.Mkdir(uploadDir, os.ModePerm);
  }
  filePath := filepath.Join(uploadDir, file.Filename);
  if err := c.SaveUploadedFile(file, filePath); err != nil {
    c.String(http.StatusInternalServerError, "error saving file: %s",err.Error());
    return;
  }
  c.String(http.StatusOK, "File uploaded successfully: %s", file.Filename);
}

func main() {
  r := gin.Default();
  r.Static("/static","./static");
  r.POST("/upload", uploadFile);

  r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	});
  r.Run(":8080");
}
