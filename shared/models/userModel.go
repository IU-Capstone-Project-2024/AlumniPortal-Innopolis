package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name             string `gorm:"not null"`
	Role             string
	Specialization   string
	AvailableCustdev bool `gorm:"default:false"`
	LastName         string
	Description      string
	Verified         bool `gorm:"default:false"`
	PortfolioLink    string
	SocialsLink      string
	Email            string `gorm:"unique"`
	Password         string
	IsAdmin          bool `gorm:"default:false"`
	IsAlumni         bool `gorm:"default:false"`
}
