package api

// API Handler
// Manages HTTP endpoints for controlling the automation (Start/Stop)
// and retrieving status/data for the Frontend dashboard.
// It bridges the gap between the HTTP layer and the BrowserManager/Store.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/sudhan/browser-automation/internal/automation"
	"github.com/sudhan/browser-automation/internal/models"
	"github.com/sudhan/browser-automation/internal/store"
)

type Handler struct {
	BM        *automation.BrowserManager
	Store     *store.Store
	mu        sync.Mutex
	IsRunning bool
	StatusLog []string
}

func NewHandler(bm *automation.BrowserManager, s *store.Store) *Handler {
	return &Handler{
		BM:        bm,
		Store:     s,
		StatusLog: []string{"Ready to start."},
	}
}

func (h *Handler) Log(msg string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	entry := fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), msg)
	h.StatusLog = append(h.StatusLog, entry)
	if len(h.StatusLog) > 100 {
		h.StatusLog = h.StatusLog[1:] // Keep last 100
	}
}

type StartRequest struct {
	Query    string `json:"query"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) StartAutomation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req StartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Default if parsing fails or empty
		req.Query = "Recruiter"
	}
	if req.Query == "" {
		req.Query = "Recruiter"
	}

	h.mu.Lock()
	if h.IsRunning {
		h.mu.Unlock()
		json.NewEncoder(w).Encode(map[string]string{"message": "Already running"})
		return
	}
	h.IsRunning = true
	h.mu.Unlock()

	go h.runAutomationFlow(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Automation started with query: " + req.Query})
}

func (h *Handler) StopAutomation(w http.ResponseWriter, r *http.Request) {
	// Not fully implemented graceful stop
	h.mu.Lock()
	h.IsRunning = false
	h.mu.Unlock()
	h.BM.Stop()
	json.NewEncoder(w).Encode(map[string]string{"message": "Stop signal sent"})
}

func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	resp := map[string]interface{}{
		"running": h.IsRunning,
		"logs":    h.StatusLog,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetData(w http.ResponseWriter, r *http.Request) {
	activities, err := h.Store.GetActivities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}

func (h *Handler) runAutomationFlow(req StartRequest) {
	defer func() {
		h.mu.Lock()
		h.IsRunning = false
		h.mu.Unlock()
	}()

	h.Log("Initializing browser...")
	h.BM.Start()
	// defer h.BM.Stop()

	cfg := models.LoadConfig()

	// Use dynamic credentials if provided, else fallback to env
	email := req.Email
	if email == "" {
		email = cfg.Email
	}
	pass := req.Password
	if pass == "" {
		pass = cfg.Password
	}

	h.Log("Logging in...")
	if err := h.BM.Login(email, pass); err != nil {
		h.Log(fmt.Sprintf("Login failed: %v", err))
		return
	}
	h.Log("Login successful.")

	h.Log(fmt.Sprintf("Searching for: %s", req.Query))
	profiles, err := h.BM.SearchProfiles(req.Query, 5) // Limit 5
	if err != nil {
		h.Log(fmt.Sprintf("Search failed: %v", err))
		return
	}
	h.Log(fmt.Sprintf("Found %d profiles.", len(profiles)))

	// 1. Log all found profiles to DB immediately so UI updates
	for _, p := range profiles {
		h.Store.SaveActivity(models.ProfileActivity{
			ProfileURL: p.URL,
			Action:     "SEARCH_FOUND",
			Metadata:   p.Name,
			Timestamp:  time.Now(),
		})
	}

	// 2. Process Connection Requests (Optional: make this configurable later)
	for _, p := range profiles {
		h.Log(fmt.Sprintf("Processing: %s", p.Name))

		// Simulate action
		err := h.BM.SendConnectionRequest(p.URL)
		if err != nil {
			h.Log(fmt.Sprintf("Connection failed: %v", err))
		} else {
			// Save to DB
			act := models.ProfileActivity{
				ProfileURL: p.URL,
				Action:     "CONNECT",
				Metadata:   p.Name,
				Timestamp:  time.Now(),
			}
			h.Store.SaveActivity(act)
			h.Log(fmt.Sprintf("Connection Sent to %s", p.Name))
		}
	}

	h.Log("Automation flow finished.")
}
