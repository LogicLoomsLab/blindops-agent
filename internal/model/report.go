package model

import "time"

// UsageReport represents a single line item of cloud usage.
type UsageReport struct {
	EncryptedID  string    `json:"encrypted_id"`
	Service      string    `json:"service"`
	Region       string    `json:"region"`
	Cost         float64   `json:"cost"`
	UsageDate    time.Time `json:"usage_date"`
	IsAnonymized bool      `json:"is_anonymized"` // Flag to prove encryption
}