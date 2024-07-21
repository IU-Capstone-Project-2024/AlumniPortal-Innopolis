package models

import (
	"alumniportal.com/shared/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"not null"`
	Role             helpers.UserRole
	Specialization   string
	AvailableCustdev bool `gorm:"default:false"`
	LastName         string
	Description      string
	Verified         helpers.VerificationUserStatus `gorm:"not null;default:'UnverifiedUser'"`
	PortfolioLink    string
	SocialsLink      string
	Email            string `gorm:"unique"`
	Password         string
}
