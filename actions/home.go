package actions

import "net/http"

// Home renders the Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
