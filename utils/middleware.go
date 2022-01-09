package utils

import "net/http"

func IsloggedIn(r *http.Request) string {
	c, err := r.Cookie("UserRole")
	if err != nil {
		return "0"
	}

	return c.Value
}
