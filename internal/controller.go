package internal

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controllers interface {
	InsertPayroll(c *gin.Context)
}

type controllers struct {
	db   *gorm.DB
	data []*RequestCsv
}

func NewController(db *gorm.DB, data []*RequestCsv) Controllers {
	return &controllers{
		db:   db,
		data: data,
	}
}

func (s *controllers) InsertPayroll(c *gin.Context) {

}
