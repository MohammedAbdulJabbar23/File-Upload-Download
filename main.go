package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type File struct {
	ID       string
	Filename string
	Path     string
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("postgres", "postgresql://root:password@localhost:5433/file-upload?sslmode=disable")
	if err != nil {
		panic(err)
	}
	createTable := `
  CREATE TABLE IF NOT EXISTS files (
		id UUID PRIMARY KEY,
		filename TEXT NOT NULL,
		path TEXT NOT NULL
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		panic(err)
	}

}

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file");
	if err != nil {
		c.String(http.StatusBadRequest, "Error retrieving file: %s", err.Error());
		return;
	}
	uploadDir := "./uploads";
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755);
	}
	fileID := uuid.New().String();
	fileName := file.Filename;
	fileExt := filepath.Ext(fileName);
	fileNameWithoutExt := fileName[0 : len(fileName)-len(fileExt)];
	newFileName := fmt.Sprintf("%s_%s%s", fileNameWithoutExt, fileID, fileExt);
	filePath := filepath.Join(uploadDir, newFileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.String(http.StatusInternalServerError, "Error saving file: %s", err.Error());
		return;
	}
	insertQuery := `
	INSERT INTO files (id, filename, path)
	VALUES ($1, $2, $3);
	`;
	_, err = db.Exec(insertQuery, fileID, fileName, filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error inserting file into database: %s", err.Error())
		return
	}

	downloadLink := fmt.Sprintf("/download/%s", fileID)
	c.String(http.StatusOK, "File uploaded successfully. Download link: %s", "localhost:8080" + downloadLink)
}
func downloadFile(c *gin.Context) {
	fileID := c.Param("fileID");
	var file File;
	row := db.QueryRow("SELECT id, filename, path FROM files WHERE id=$1", fileID);
	err := row.Scan(&file.ID, &file.Filename, &file.Path)
	if err != nil {
		c.String(http.StatusBadRequest, "file not found");
		return;
	}
	c.FileAttachment(file.Path, file.Filename);
}

func main() {
	initDB();
	defer db.Close();
	r := gin.Default();

	r.Static("/static", "./static");
	r.POST("/upload", uploadFile);
	r.GET("/download/:fileID", downloadFile);

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html");
	});

	r.Run("0.0.0.0:8080");
}
