package internal

import (
	"time"
)

type User struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:255;not null" json:"name"`
	Status bool   `gorm:"not null;" json:"status"`
}

type Payroll struct {
	ID            uint      `gorm:"primaryKey"`
	Batch         int       `gorm:"not null" json:"batch"`
	UserId        int       `gorm:"not null;" json:"user_id"`
	AccountName   string    `gorm:"not null;" json:"account_name"`
	AccountNumber string    `gorm:"not null;" json:"account_number"`
	Status        bool      `gorm:"not null;" json:"status"`
	Amount        int       `gorm:"not null;" json:"amount"`
	CreatedAt     time.Time `gorm:"not null;" json:"created_at"`
	CreatedBy     time.Time `gorm:"not null;" json:"created_by"`
}

type PayrollLog struct {
	ID           uint      `gorm:"primaryKey"`
	Batch        int       `gorm:"not null" json:"batch"`
	TotalSuccess bool      `gorm:"not null;" json:"total_success"`
	TotalFailed  bool      `gorm:"not null;" json:"total_failed"`
	FileName     string    `gorm:"not null" json:"file_name"`
	CreatedAt    time.Time `gorm:"not null;" json:"created_at"`
	CreatedBy    time.Time `gorm:"not null;" json:"created_by"`
}

type PayrollFailedLog struct {
	ID            uint      `gorm:"primaryKey"`
	Batch         int       `gorm:"not null" json:"batch"`
	UserId        int       `gorm:"not null;" json:"user_id"`
	FileName      string    `gorm:"not null" json:"file_name"`
	AccountName   string    `gorm:"not null;" json:"account_name"`
	AccountNumber string    `gorm:"not null;" json:"account_number"`
	CreatedAt     time.Time `gorm:"not null;" json:"created_at"`
	CreatedBy     time.Time `gorm:"not null;" json:"created_by"`
}

type RequestCsv struct {
	Batch         string `csv:"batch"`
	AccountName   string `csv:"account_name"`
	AccountNumber string `csv:"account_number"`
	UserId        int    `csv:"user_id"`
	Amount        int    `csv:"amount"`
	Status        bool   `csv:"status"`
}
