package models

import (
	"alumniportal.com/shared/helpers"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	FounderID   uint                       `gorm:"not null;index"`
	User        User                       `gorm:"foreignKey:FounderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string                     `gorm:"not null"`
	Description string                     `gorm:"not null"`
	Finished    bool                       `gorm:"not null;default:false"`
	Date        datatypes.Date             `gorm:"not null"`
	Duration    int                        `gorm:"not null"`
	Status      helpers.VerificationStatus `gorm:"not null;default:'Unverified'"`
}
