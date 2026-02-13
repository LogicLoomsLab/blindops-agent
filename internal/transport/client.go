package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/LogicLoomsLab/blindops-agent/internal/model"
)

// Client handles communication with the BlindOps Core.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new transporter.
func NewClient(url string) *Client {
	return &Client{
		BaseURL: url,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second, // Always set timeouts!
		},
	}
}

// SendReport marshals the report to JSON and POSTs it to the backend.
func (c *Client) SendReport(report model.UsageReport) error {
	// 1. Convert struct to JSON
	jsonData, err := json.Marshal(report)
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	// 2. Create the Request
	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "BlindOps-Agent/v0.1")

	// 3. Send the Request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 4. Check Response Code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned error: %d", resp.StatusCode)
	}

	return nil
}