package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LogicLoomsLab/blindops-agent/internal/crypto"
	"github.com/LogicLoomsLab/blindops-agent/internal/model"
	"github.com/LogicLoomsLab/blindops-agent/internal/transport" // Import the new package
)

func main() {
	// CONFIGURATION (Hardcoded for MVP)
	// Pointing to your local Python Backend
	backendURL := "http://127.0.0.1:8000/api/v1/agent/report"
	
	fmt.Println("Starting BlindOps Agent...")
	fmt.Printf("Target Core: %s\n", backendURL)

	// 1. Setup Encryption Key
	key, err := crypto.GenerateRandomKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}

	// 2. Initialize the Transporter
	client := transport.NewClient(backendURL)

	// 3. Simulate a Loop (Sending data)
	// In production, this would loop through AWS resources.
	// For now, we send one dummy report.
	
	fmt.Println("Simulating resource scan...")
	
	rawResourceID := "i-0123456789abcdef0"
	service := "AWS EC2"
	cost := 42.50

	// Encrypt
	encryptedID, err := crypto.EncryptResourceID(rawResourceID, key)
	if err != nil {
		log.Fatalf("Encryption failed: %v", err)
	}

	// Create Payload
	report := model.UsageReport{
		EncryptedID:  encryptedID,
		Service:      service,
		Region:       "us-east-1",
		Cost:         cost,
		UsageDate:    time.Now(),
		IsAnonymized: true,
	}

	// 4. TRANSMIT
	fmt.Printf("Sending encrypted report for [%s]...\n", service)
	
	err = client.SendReport(report)
	if err != nil {
		log.Fatalf("Transmission FAILED: %v", err)
	}

	fmt.Println("------------------------------------------------")
	fmt.Println("SUCCESS: Report delivered to BlindOps Core.")
	fmt.Println("------------------------------------------------")
}