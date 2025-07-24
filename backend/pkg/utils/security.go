// pkg/utils/security.go
package utils 

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt" 
)

// Pre-compile the regex once for better performance
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

// ValidateEmail checks if the provided email string has a valid format.
// It returns an error if the email is empty or has an invalid format.
func ValidateEmail(email string) error { 
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

// HashedPassword generates a bcrypt hash of the provided password.
// It returns an error if the password is too short or if hashing fails.
func HashedPassword(password string) (string, error) { 
	if len(password) < 8 {
		return "", fmt.Errorf("password must be at least 8 characters long")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		
		return "", fmt.Errorf("failed to generate bcrypt hash: %w", err)
	}
	return string(bytes), nil
}