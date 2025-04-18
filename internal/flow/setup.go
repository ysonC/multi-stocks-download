package flow

import (
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

func SetupDirectories(dirs ...string) {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Error creating directory '%s': %v", dir, err)
		}
	}
}

func SetupPlaywright() *playwright.Playwright {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Failed to start Playwright: %v", err)
	}
	return pw
}
