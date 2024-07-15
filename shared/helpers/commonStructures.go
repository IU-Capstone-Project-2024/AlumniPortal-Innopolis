package helpers

type VerificationStatus string

const (
	Unverified VerificationStatus = "Unverified"
	Accepted   VerificationStatus = "Accepted"
	Declined   VerificationStatus = "Declined"
)
