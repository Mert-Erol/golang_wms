package utils

import "net/http"

// Checking the user logged in
func IsloggedIn(r *http.Request) string {
	c, err := r.Cookie("UserRole")
	if err != nil {
		return "0"
	}

	return c.Value
}
