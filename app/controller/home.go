package controller

import (
	"appeals/app/model"
	"appeals/app/view"
	"html/template"
	"net/http"
)

type home struct {
	homeTemplate  *template.Template
	loginTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
	http.HandleFunc("/logout", h.handleLogout)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := view.NewHome()
	w.Header().Add("Content-Type", "text/html")
	_, err := model.Sessions(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := view.NewLogin()
	if r.Method == http.MethodPost {
		model.Authenticate(w, r)
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}

func (h home) handleLogout(w http.ResponseWriter, r *http.Request) {
	vm := view.NewLogin()
	w.Header().Add("Content-Type", "text/html")
	model.Logout(w, r)
	h.loginTemplate.Execute(w, vm)
}
