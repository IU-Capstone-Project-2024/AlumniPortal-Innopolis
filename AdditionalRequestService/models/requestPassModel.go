package models

import (
	"gorm.io/gorm"
	"shared/models"
	"time"
)

type PassType string

const (
	Dormitory  PassType = "Dormitory"
	University PassType = "University"
)

type PassRequestStatus string

const (
	Unverified PassRequestStatus = "Unverified"
	Accepted   PassRequestStatus = "Accepted"
	Declined   PassRequestStatus = "Declined"
)

type PassRequest struct {
	gorm.Model
	UserID             uint        `gorm:"not null"`
	User               models.User `gorm:"foreignKey:UserID"`
	PassStartDate      time.Time   `gorm:"not null"`
	PassExpirationDate time.Time   `gorm:"not null"`
	Message            string
	PassType           PassType          `gorm:"type:enum('Dormitory', 'University');not null"`
	Status             PassRequestStatus `gorm:"type:enum('Unverified', 'Accepted', 'Declined');not null;default:'Unverified'"`
}
