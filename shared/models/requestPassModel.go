package models

import (
	"alumniportal.com/shared/helpers"
	"gorm.io/gorm"
	"time"
)

type PassType string

const (
	Dormitory  PassType = "Dormitory"
	University PassType = "University"
)

type PassRequest struct {
	gorm.Model
	UserID             uint      `gorm:"not null;index"`
	User               User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PassStartDate      time.Time `gorm:"not null"`
	PassExpirationDate time.Time `gorm:"not null"`
	Message            string
	PassType           PassType                   `gorm:"not null"`
	Status             helpers.VerificationStatus `gorm:"not null;default:'Unverified'"`
}
