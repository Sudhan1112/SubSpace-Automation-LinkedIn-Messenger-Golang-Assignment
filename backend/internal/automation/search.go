package automation

import (
	"fmt"
	"time"

	"github.com/go-rod/rod/lib/input"
	"github.com/sudhan/browser-automation/internal/models"
)

// SearchProfiles searches for people using the provided query and returns a list of profile URLs.
// It handles pagination to limit the number of results.
func (bm *BrowserManager) SearchProfiles(query string, limit int) ([]models.Profile, error) {
	fmt.Printf("Searching for profiles with query: %s\n", query)

	// Encode query
	searchURL := fmt.Sprintf("https://www.linkedin.com/search/results/people/?keywords=%s", query)

	bm.Page.MustNavigate(searchURL)
	bm.Page.MustWaitLoad()
	bm.RandomSleep(4000, 6000) // Increased wait time

	var profiles []models.Profile
	pageCount := 1

	for len(profiles) < limit {
		fmt.Printf("Scraping page %d...\n", pageCount)

		// Wait for results to appear
		// We accept either the list or the empty state to verify page load
		// The list is usually 'ul.reusable-search__entity-result-list'
		// The items are 'li.reusable-search__result-container'

		listSelector := ".reusable-search__entity-result-list"

		if err := bm.Page.Timeout(10*time.Second).WaitElementsMoreThan(listSelector, 0); err != nil {
			fmt.Println("Warning: Results list not found. Checking for empty state...")
			if exists, _, _ := bm.Page.Has(".artdeco-empty-state__headline"); exists {
				fmt.Println("No results found (Empty State).")
				break
			}
			// One more retry: sometimes the list class varies, try a generous fallback
			if exists, _, _ := bm.Page.Has(".search-results-container"); !exists {
				fmt.Println("Could not find search results container.")
				break
			}
		}

		// Select all result items
		// LinkedIn structure: ul.reusable-search__entity-result-list > li.reusable-search__result-container
		itemSelector := "li.reusable-search__result-container"
		elements, err := bm.Page.Elements(itemSelector)
		if err != nil {
			fmt.Printf("Error finding elements: %v\n", err)
			break
		}

		fmt.Printf("Found %d visible elements on page %d\n", len(elements), pageCount)

		for _, el := range elements {
			if len(profiles) >= limit {
				break
			}

			// Scroll element into view
			el.ScrollIntoView()

			// Extract Name & URL
			var name, url string

			// Try primary selector for Title/Link
			linkEl, err := el.Element(".entity-result__title-text a")
			if err == nil {
				if hrefPtr, _ := linkEl.Attribute("href"); hrefPtr != nil {
					url = *hrefPtr
				}
				// Name is inside a span with aria-hidden="true" usually
				nameEl, err := linkEl.Element("span[aria-hidden='true']")
				if err == nil {
					name, _ = nameEl.Text()
				} else {
					// Fallback to full text of link
					name, _ = linkEl.Text()
				}
			}

			if name != "" && url != "" {
				profiles = append(profiles, models.Profile{
					Name: name,
					URL:  url,
				})
			}
		}

		if len(profiles) >= limit {
			break
		}

		// Handle Pagination
		// Scroll to bottom to ensure 'Next' button is visible
		bm.Page.Keyboard.Press(input.PageDown)
		bm.RandomSleep(1000, 2000)
		bm.Page.Keyboard.Press(input.PageDown)
		bm.RandomSleep(1000, 2000)

		// Click Next
		nextBtn, err := bm.Page.Element("button[aria-label='Next']")
		if err != nil {
			fmt.Println("No next page button found.")
			break
		}

		if disabled, _ := nextBtn.Attribute("disabled"); disabled != nil {
			fmt.Println("Last page reached.")
			break
		}

		bm.HumanMove("button[aria-label='Next']")
		nextBtn.MustClick()
		bm.Page.MustWaitLoad()
		bm.RandomSleep(3000, 5000) // Longer wait between pages
		pageCount++
	}

	return profiles, nil
}
