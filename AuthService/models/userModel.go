package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name             string `gorm:"not null"`
	Surname          string
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
}
