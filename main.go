package main

import (
	"github.com/DonCuervoC/banking_go_api_hex/app" // Import app package
	"github.com/DonCuervoC/banking_go_api_hex/logger"
)

func main() {
	// Call the Start function of the app package
	logger.Info("Starting the application...")
	app.Start()
}
