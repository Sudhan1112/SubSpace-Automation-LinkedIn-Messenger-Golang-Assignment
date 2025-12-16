package models

import "time"

type Profile struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status string `json:"status"` // e.g., "New", "Request Sent", "Connected"
}

// ProfileActivity tracks actions taken on profiles
type ProfileActivity struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProfileURL string    `json:"profile_url"`
	Action    string    `json:"action"` // "CONNECT", "MESSAGE"
	Metadata  string    `json:"metadata"` // e.g., message content
	Timestamp time.Time `json:"timestamp"`
}
