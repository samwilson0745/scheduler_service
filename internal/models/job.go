package models

import (
	"time"

	"gorm.io/gorm"
)

type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusPaused    JobStatus = "paused"
)

type Job struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100); uniqueIndex; not null" json:"name"`
	Cron        string    `gorm:"type:varchar(50); not null" json:"cron"`
	Status      JobStatus `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	Description string    `gorm:"type:text" json:"description"`

	// Execution details
	LastRun   time.Time `json:"last_run"` // No default needed, nullable by default
	NextRun   time.Time `json:"next_run"` // No default needed, nullable by default
	RunCount  int       `gorm:"default:0" json:"run_count"`
	LastError string    `gorm:"type:text" json:"last_error"`

	// Control fields
	IsActive bool `gorm:"default:true" json:"is_active"`
	Timeout  int  `gorm:"default:3600" json:"timeout"`

	// Audit fields
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Optional: Parameters for the job
	Parameters string `gorm:"type:text" json:"parameters"`
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&Job{})
}
