package controller

import (
	"appeals/app/model"
	"net/http"
	"strconv"
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
	appeal := r.URL.Query().Get("appeal")
	if appeal != "" {
		if appealID, err := strconv.Atoi(appeal); err == nil {
			app, err := model.GetAppealByID(appealID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			notice := model.NewNotice(app)
			notice.Document.Output(w)
		}
	}
}
