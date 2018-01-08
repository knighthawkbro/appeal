package controller

import (
	"appeals/app/view"
	"html/template"
	"net/http"
)

type appeal struct {
	appealTemplate *template.Template
}

func (a appeal) registerRoutes() {
	http.HandleFunc("/appeal", a.handleAppeal)
}

func (a appeal) handleAppeal(w http.ResponseWriter, r *http.Request) {
	vm := view.NewAppeal()
}
