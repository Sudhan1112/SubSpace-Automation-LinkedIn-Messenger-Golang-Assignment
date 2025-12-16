package main

import (
	"fmt"

	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	fmt.Println("Checking/Downloading Browser...")
	u := launcher.New().
		Headless(false).
		Devtools(true).
		Leakless(false).
		MustLaunch()
	fmt.Printf("Browser launched successfully at: %s\n", u)
}
