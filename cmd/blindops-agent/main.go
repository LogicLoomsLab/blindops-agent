package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LogicLoomsLab/blindops-agent/internal/crypto"
	"github.com/LogicLoomsLab/blindops-agent/internal/model"
)

func main() {
	fmt.Println("Starting BlindOps Agent...")

	// 1. Setup Encryption Key (In production, this comes from Vault/Env)
	key, err := crypto.GenerateRandomKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}
	fmt.Printf("Session Key Generated (32 bytes)\n")

	// 2. Simulate a sensitive AWS Resource
	rawResourceID := "i-0123456789abcdef0" // This should never leave the machine
	service := "AWS EC2"
	cost := 42.50

	// 3. Encrypt the ID
	encryptedID, err := crypto.EncryptResourceID(rawResourceID, key)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}

	// 4. Construct the Payload
	report := model.UsageReport{
		EncryptedID:  encryptedID,
		Service:      service,
		Region:       "us-east-1",
		Cost:         cost,
		UsageDate:    time.Now(),
		IsAnonymized: true,
	}

	// 5. Output Verification
	fmt.Println("------------------------------------------------")
	fmt.Println("PRIVACY CHECK:")
	fmt.Printf("Original ID:  %s\n", rawResourceID)
	fmt.Printf("Encrypted ID: %s\n", report.EncryptedID)
	fmt.Println("------------------------------------------------")
	fmt.Println("Ready to transmit to BlindOps Core.")
}