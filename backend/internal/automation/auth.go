package automation

import (
	"fmt"
	"strings"
	"time"
)

func (bm *BrowserManager) Login(email, password string) error {
	loginURL := "https://www.linkedin.com/login"
	
	fmt.Println("Navigating to login page...")
	bm.Page.MustNavigate(loginURL)
	bm.Page.MustWaitLoad()

	// Check if already logged in or if captcha is present
	// Simple check: Look for "Feed" or "Me" icon
	if exists, _, _ := bm.Page.Has(".global-nav__me"); exists {
		fmt.Println("Already logged in.")
		return nil
	}

	fmt.Println("Entering credentials...")
	
	// Email
	emailInput := bm.Page.MustElement("#username")
	emailInput.MustInput(email)
	bm.RandomSleep(500, 1000)

	// Password
	passInput := bm.Page.MustElement("#password")
	passInput.MustInput(password)
	bm.RandomSleep(600, 1200)

	// Click Sign In
	signInBtn := bm.Page.MustElement(".btn__primary--large")
	signInBtn.MustClick()
	
	// Wait for navigation
	bm.Page.MustWaitNavigation() // This waits for page load after click

	// Check for CAPTCHA / Security Challenge
	// Note: We are just DETECTING it as per requirements.
	
	// Give it a moment to render whatever comes next
	time.Sleep(2 * time.Second)

	if heading, err := bm.Page.Element("h1"); err == nil {
		text := heading.MustText()
		if strings.Contains(strings.ToLower(text), "verification") || 
		   strings.Contains(strings.ToLower(text), "security") ||
		   strings.Contains(strings.ToLower(text), "challenge") {
			return fmt.Errorf("security check detected: %s", text)
		}
	}

	// Verify login success
	// We look for the main feed element or nav bar
	// LinkedIn feed is usually "#voyager-feed" or global nav
	if exists, _, _ := bm.Page.Has(".global-nav__me"); exists {
		fmt.Println("Login successful!")
		return nil
	}

	return fmt.Errorf("login failed or unknown state")
}
