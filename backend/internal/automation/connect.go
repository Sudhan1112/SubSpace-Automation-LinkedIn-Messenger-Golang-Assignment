package automation

import (
	"fmt"
	"strings"
)

func (bm *BrowserManager) SendConnectionRequest(profileURL string) error {
	fmt.Printf("Visiting profile: %s\n", profileURL)
	bm.Page.MustNavigate(profileURL)
	bm.Page.MustWaitLoad()
	bm.RandomSleep(2000, 5000)

	// Check for "Pending" first
	if exists, _, _ := bm.Page.Has("button[disabled]"); exists {
		// Need better check for "Pending" text
	}

	// Try finding the Connect button
	buttons := bm.Page.MustElements("main button.artdeco-button--primary")
	clicked := false

	for _, btn := range buttons {
		text, err := btn.Text()
		if err == nil && strings.Contains(strings.ToLower(text), "connect") {
			fmt.Println("Found Connect button")
			btn.MustClick()
			clicked = true
			break
		}
	}

	if !clicked {
		fmt.Println("Connect button not found, checking 'More' menu...")
		moreBtn, err := bm.Page.Element("button[aria-label*='More actions']")
		if err == nil {
			moreBtn.MustClick()
			bm.RandomSleep(500, 1000)

			dropdownItems := bm.Page.MustElements(".artdeco-dropdown__item")
			for _, item := range dropdownItems {
				if strings.Contains(strings.ToLower(item.MustText()), "connect") {
					item.MustClick()
					clicked = true
					break
				}
			}
		}
	}

	if !clicked {
		return fmt.Errorf("could not find connect button")
	}

	bm.RandomSleep(1000, 2000)

	// Handle "Add a note" modal
	sendBtn, err := bm.Page.Element("button[aria-label*='Send without note']")
	if err == nil {
		fmt.Println("Clicking Send without note")
		sendBtn.MustClick()
	} else {
		if modal, err := bm.Page.Element(".artdeco-modal"); err == nil {
			if btn, err := modal.Element("button.artdeco-button--primary"); err == nil {
				btn.MustClick()
			}
		}
	}

	fmt.Println("Connection request sent!")
	return nil
}
