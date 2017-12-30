package controller

import (
	"appeals/app/model"
	"net/http"
)

type notice struct{}

func (n notice) registerRoutes() {
	http.HandleFunc("/notice", n.handleNotice)
}

func (n notice) handleNotice(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/pdf")
	_, err := model.Sessions(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	notice := model.NewNotice()
	notice.Document.Output(w)
}
