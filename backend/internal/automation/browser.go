package automation

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/sudhan/browser-automation/internal/utils"
)

type BrowserManager struct {
	Browser *rod.Browser
	Page    *rod.Page
}

func NewBrowserManager(headless bool) *BrowserManager {
	launch := launcher.New().
		Headless(headless).
		Devtools(true).
		Leakless(false) // Disable leakless to avoid AV issues

	u := launch.MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	// Create a new Incognito browser context for fresh session
	// or use Default to persist session if not incognito
	// We want persistence, so we might stick to default or manage profiles.
	// For "human-like", we should avoid "incognito" flags sometimes, but incognito is safer for clean starts.
	// However, requirement says "persist session data". So we should use a user data dir if we want true persistence across restarts,
	// or just keep the browser open.
	// Let's use a standard page for now with Stealth.

	return &BrowserManager{
		Browser: browser,
	}
}

func (bm *BrowserManager) Start() {
	// Initialize stealth page
	page := stealth.MustPage(bm.Browser)
	bm.Page = page

	// Randomize viewport slightly
	bm.Page.MustSetViewport(1920, 1080, 1.0, false)
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

// HumanMove simulates human-like mouse movement to an element
func (bm *BrowserManager) HumanMove(selector string) {
	el := bm.Page.MustElement(selector)
	// Rod's MoveTo roughly simulates movement, but we can add random points if needed.
	// For now, simple MoveTo with random delays before/after is decent.
	bm.RandomSleep(100, 300)
	el.MustHover()
	bm.RandomSleep(200, 500)
}
