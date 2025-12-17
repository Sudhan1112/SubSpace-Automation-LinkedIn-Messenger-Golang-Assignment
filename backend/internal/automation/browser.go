package automation

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/sudhan/browser-automation/internal/utils"
)

type BrowserManager struct {
	Browser *rod.Browser
	Page    *rod.Page
	LastX   float64
	LastY   float64
}

func NewBrowserManager(headless bool) *BrowserManager {
	// STEALTH: Use Incognito mode for Context Isolation
	launch := launcher.New().
		Headless(headless).
		Devtools(true).
		Leakless(false).         // Disable leakless to avoid AV issues
		Set("incognito", "true") // Explicitly enable incognito

	u := launch.MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	return &BrowserManager{
		Browser: browser,
		LastX:   0,
		LastY:   0,
	}
}

func (bm *BrowserManager) Start() {
	// STEALTH: Initialize stealth page to mask navigator properties
	page := stealth.MustPage(bm.Browser)
	bm.Page = page

	// STEALTH: Randomize viewport slightly to avoid identical fingerprints
	width := 1920 + rand.Intn(100) - 50
	height := 1080 + rand.Intn(100) - 50
	bm.Page.MustSetViewport(width, height, 1.0, false)
}

func (bm *BrowserManager) Stop() {
	if bm.Browser != nil {
		bm.Browser.MustClose()
	}
}

// RandomSleep sleeps for a random duration between min and max milliseconds
func (bm *BrowserManager) RandomSleep(min, max int) {
	utils.RandomSleep(min, max)
}

// HumanMove simulates human-like mouse movement using stepwise smoothing
func (bm *BrowserManager) HumanMove(selector string) {
	el := bm.Page.MustElement(selector)
	box := el.MustShape().Box()

	// Target point (center of element with random jitter)
	targetX := box.X + box.Width/2 + (rand.Float64()*10 - 5)
	targetY := box.Y + box.Height/2 + (rand.Float64()*10 - 5)

	// We interpolate from bm.LastX, bm.LastY to targetX, targetY
	steps := 20
	startX := bm.LastX
	startY := bm.LastY

	// If it's the first move (0,0), we might just warp or assume mouse was there.
	// In a real browser, mouse starts at 0,0 usually.

	for i := 1; i <= steps; i++ {
		t := float64(i) / float64(steps)
		// Linear interpolation
		x := startX + (targetX-startX)*t
		y := startY + (targetY-startY)*t

		// Add slight random jitter to the path
		jitterX := (rand.Float64() - 0.5) * 2
		jitterY := (rand.Float64() - 0.5) * 2

		// Use MustMoveTo instead of Move
		bm.Page.Mouse.MustMoveTo(x+jitterX, y+jitterY)

		// Sleep tiny amount between steps
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)+2))
	}

	// Update last position
	bm.LastX = targetX
	bm.LastY = targetY

	// Final hover to ensure events trigger and snap to element
	el.MustHover()
	bm.RandomSleep(200, 500)
}

// TypeHuman simulates human typing with variable delays
func (bm *BrowserManager) TypeHuman(selector, text string) {
	el := bm.Page.MustElement(selector)
	el.MustClick()

	bm.RandomSleep(200, 500)

	for _, char := range text {
		bm.Page.Keyboard.Type(input.Key(char))
		// STEALTH: Random delay between keystrokes (50ms - 150ms)
		bm.RandomSleep(50, 150)
	}
}
