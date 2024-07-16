package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PaymentMethod string

const (
	CreditCard   PaymentMethod = "Credit Card"
	PayPal       PaymentMethod = "PayPal"
	BankTransfer PaymentMethod = "Bank Transfer"
)

type RecurringDonation string

const (
	Monthly   RecurringDonation = "Monthly"
	Quarterly RecurringDonation = "Quarterly"
	Yearly    RecurringDonation = "Yearly"
)

type Donation struct {
	gorm.Model
	DonatorID         uint              `gorm:"not null; index"`
	User              User              `gorm:"foreignKey:DonatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProjectID         uint              `gorm:"not null; index"`
	Project           Project           `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount            float32           `gorm:"not null"`
	Date              datatypes.Date    `gorm:"not null"`
	PaymentMethod     PaymentMethod     `gorm:"not nul"`
	RecurringDonation RecurringDonation `gorm:"not null"`
}
