package automation

import (
	"fmt"
)

func (bm *BrowserManager) SendMessage(profileURL, message string) error {
	fmt.Printf("Sending message to: %s\n", profileURL)
	bm.Page.MustNavigate(profileURL)
	bm.Page.MustWaitLoad()
	bm.RandomSleep(2000, 4000)

	// Look for "Message" button
	msgBtn, err := bm.Page.Element("button[aria-label^='Message']")
	if err != nil {
		// Try finding by text
		buttons := bm.Page.MustElements("main button")
		for _, btn := range buttons {
			if txt, err := btn.Text(); err == nil && txt == "Message" {
				msgBtn = btn
				break
			}
		}
	}

	if msgBtn == nil {
		return fmt.Errorf("message button not found")
	}

	msgBtn.MustClick()
	bm.RandomSleep(1000, 2000)

	// Messaging overlay should appear
	inputEl, err := bm.Page.Element(".msg-form__contenteditable")
	if err != nil {
		return fmt.Errorf("message input not found")
	}

	inputEl.MustInput(message)
	bm.RandomSleep(1000, 2000)

	// Click Send
	sendBtn, err := bm.Page.Element(".msg-form__send-button")
	if err != nil {
		// Fallback
		sendBtn, err = bm.Page.Element("button[type='submit']")
	}

	if err != nil {
		return fmt.Errorf("send button not found")
	}

	// Double check if enabled
	if disabled, _ := sendBtn.Attribute("disabled"); disabled != nil {
		return fmt.Errorf("send button is disabled (empty message?)")
	}

	sendBtn.MustClick()
	fmt.Println("Message sent!")

	return nil
}
