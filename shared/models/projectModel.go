package models

import (
	"gorm.io/gorm"
)

type ProjectVerificationStatus string

const (
	UnverifiedProject ProjectVerificationStatus = "Unverified"
	AcceptedProject   ProjectVerificationStatus = "Accepted"
	DeclinedProject   ProjectVerificationStatus = "Declined"
)

type Project struct {
	gorm.Model
	FounderID   uint                      `gorm:"not null;index"`
	User        User                      `gorm:"foreignKey:FounderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string                    `gorm:"not null"`
	Description string                    `gorm:"not null"`
	Finished    bool                      `gorm:"not null;default:false"`
	Status      ProjectVerificationStatus `gorm:"not null;default:'UnverifiedProject'"`
}

type ProjectParticipant struct {
	gorm.Model
	ProjectID uint                      `gorm:"not null;index"`
	Project   Project                   `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID    uint                      `gorm:"not null;index"`
	User      User                      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status    ProjectVerificationStatus `gorm:"not null;default:'UnverifiedProject'"`
}
