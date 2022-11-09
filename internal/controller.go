package internal

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controllers interface {
	InsertPayroll(c *gin.Context)
}

type controllers struct {
	db *gorm.DB
}

func Controller(db *gorm.DB) Controllers {
	return &controllers{
		db: db,
	}
}

func (s *controllers) InsertPayroll(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	var userData []*User
	data := GetData(file)
	if err := s.db.Where("status = 'Active'").Find(&userData).Error; err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Data user tidak ditemukan"))
		return
	}

	s.db.CreateInBatches(data, len(data))

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}
