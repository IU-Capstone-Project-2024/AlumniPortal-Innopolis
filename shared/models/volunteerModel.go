package models

import (
	"gorm.io/gorm"
)

type VolunteerVerificationStatus string

const (
	UnverifiedVolunteer VolunteerVerificationStatus = "Unverified"
	AcceptedVolunteer   VolunteerVerificationStatus = "Accepted"
	DeclinedVolunteer   VolunteerVerificationStatus = "Declined"
)

type Volunteer struct {
	gorm.Model
	UserID    uint                        `gorm:"not null;index"`
	User      User                        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProjectID uint                        `gorm:"not null;index"`
	Project   Project                     `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Role      string                      `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status    VolunteerVerificationStatus `gorm:"not null;default:'UnverifiedVolunteer'"`
}