package utils

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
)

func CheckRequiredVariables() {
	requiredVars := []string{"PORT"}

	var missingVars bool

	for _, variable := range requiredVars {
		if value, exists := os.LookupEnv(variable); !exists || value == "" {
			log.Errorf("Missing required variable: %s", variable)
			missingVars = true
		}
	}

	if missingVars {
		os.Exit(1)
	}
}
