package main

import "net/http"

func isloggedIn(r *http.Request) bool {
	c, err := r.Cookie("user")
	if err != nil {
		return false
	}
	username := c.Value
	_, ok := databaseUsers[username]
	return ok
}

func getUser(r *http.Request) User {
	var user User
	c, err := r.Cookie("user")
	if err != nil {
		return user
	}
	if uid, ok := databaseSessions[c.Value]; ok {
		user = databaseUsers[uid]
	}
	return user
}
