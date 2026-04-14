package handlers

import "net/http"

const themeCookieName = "theme"

func isDarkMode(r *http.Request) bool {
	cookie, err := r.Cookie(themeCookieName)
	if err != nil {
		return false
	}

	return cookie.Value == "dark"
}
