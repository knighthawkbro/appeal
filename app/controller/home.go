package controller

import (
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
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := view.NewHome()
	w.Header().Add("Content-Type", "text/html")
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := view.NewLogin()
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}
