package controller

import (
	"appeals/app/model"
	"appeals/app/view"
	"fmt"
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
		err := model.Authenticate(w, r)
		if err == nil {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
		}
		vm.Warning = fmt.Sprint(err)
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}

func (h home) handleLogout(w http.ResponseWriter, r *http.Request) {
	vm := view.NewLogin()
	w.Header().Add("Content-Type", "text/html")
	model.Logout(w, r)
	vm.Warning = ""
	h.loginTemplate.Execute(w, vm)
}

func (h home) handleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	_, err := model.Sessions(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	r.ParseForm()
	id := r.Form.Get("id")
	r, err = http.NewRequest(http.MethodGet, "/appeal/"+id, r.Body)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, fmt.Sprintf("/appeal/%v", id), http.StatusTemporaryRedirect)
}
