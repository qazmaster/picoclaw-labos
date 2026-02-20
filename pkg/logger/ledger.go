package logger

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// LedgerEvent represents a single event in the audit trail.
type LedgerEvent struct {
	Timestamp string                 `json:"timestamp"`
	EventType string                 `json:"event_type"`
	Data      map[string]interface{} `json:"data,omitempty"`
	PrevHash  string                 `json:"prev_hash,omitempty"`
	Hash      string                 `json:"hash"`
}

var (
	ledgerMutex sync.Mutex
	lastHash    string
	ledgerPath  string
)

// InitLedger initializes the JSONL ledger file path and reads the last hash if it exists.
func InitLedger(workspacePath string) error {
	ledgerMutex.Lock()
	defer ledgerMutex.Unlock()

	stateDir := filepath.Join(workspacePath, "state")
	if err := os.MkdirAll(stateDir, 0755); err != nil {
		return fmt.Errorf("failed to create state dir for ledger: %w", err)
	}

	ledgerPath = filepath.Join(stateDir, "BOS_Ledger.jsonl")

	// Try to read the last line to get the previous hash
	file, err := os.Open(ledgerPath)
	if err == nil {
		defer file.Close()
		// Basic approach: read entire file and get last line to find lastHash.
		// For huge files, reverse scanning is better, but this is sufficient for now.
		stat, err := file.Stat()
		if err == nil && stat.Size() > 0 {
			buf := make([]byte, stat.Size())
			if _, err := file.Read(buf); err == nil {
				// Find last newline before EOF
				data := string(buf)
				lines := strings.Split(strings.TrimSpace(data), "\n")
				if len(lines) > 0 {
					lastLine := lines[len(lines)-1]
					var lastEvent LedgerEvent
					if err := json.Unmarshal([]byte(lastLine), &lastEvent); err == nil {
						lastHash = lastEvent.Hash
					}
				}
			}
		}
	}

	return nil
}

// LogLedgerEvent appends an event to the BOS Ledger.
func LogLedgerEvent(eventType string, data map[string]interface{}) {
	if ledgerPath == "" {
		// Ledger not initialized, skip or log warning normally
		Warn("Ledger is not initialized, skipping event: " + eventType)
		return
	}

	ledgerMutex.Lock()
	defer ledgerMutex.Unlock()

	event := LedgerEvent{
		Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
		EventType: eventType,
		Data:      data,
		PrevHash:  lastHash,
	}

	// Create hash (Timestamp + EventType + DataJSON + PrevHash)
	dataJSON, _ := json.Marshal(data)
	hashInput := fmt.Sprintf("%s|%s|%s|%s", event.Timestamp, event.EventType, string(dataJSON), event.PrevHash)
	hashBytes := sha256.Sum256([]byte(hashInput))
	event.Hash = hex.EncodeToString(hashBytes[:])

	// Update last hash for next event
	lastHash = event.Hash

	eventJSON, err := json.Marshal(event)
	if err != nil {
		ErrorF("Failed to marshal ledger event", map[string]interface{}{"error": err})
		return
	}

	// Append to file
	file, err := os.OpenFile(ledgerPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		ErrorF("Failed to open ledger file", map[string]interface{}{"error": err})
		return
	}
	defer file.Close()

	if _, err := file.WriteString(string(eventJSON) + "\n"); err != nil {
		ErrorF("Failed to write ledger event", map[string]interface{}{"error": err})
	}
}
