package models

import (
	"alumniportal.com/shared/helpers"
	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	ProjectID uint                       `gorm:"nullable;index"`
	Project   Project                    `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EventID   uint                       `gorm:"index"`
	Event     Event                      `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsProject bool                       `gorm:"default:true"`
	UserID    uint                       `gorm:"not null;index"`
	User      User                       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status    helpers.VerificationStatus `gorm:"not null;default:'Unverified'"`
}
