package helpers

type VerificationStatus string

const (
	Unverified VerificationStatus = "Unverified"
	Accepted   VerificationStatus = "Accepted"
	Declined   VerificationStatus = "Declined"
)

type UserRole string

const (
	Student UserRole = "Student"
	Alumni  UserRole = "Alumni"
	Admin   UserRole = "Admin"
)

type VerificationUserStatus bool

const (
	UnverifiedUser VerificationUserStatus = false
	VerifiedUser   VerificationUserStatus = true
)
