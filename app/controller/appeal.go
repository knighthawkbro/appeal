package controller

import (
	"appeals/app/model"
	"appeals/app/view"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type appeal struct {
	appealTemplate *template.Template
}

func (a appeal) registerRoutes() {
	http.HandleFunc("/appeal/", a.handleAppeal)
}

func (a appeal) handleAppeal(w http.ResponseWriter, r *http.Request) {
	appealPattern, _ := regexp.Compile(`/appeal/(\d+)`)
	match := appealPattern.FindStringSubmatch(r.URL.Path)
	appeal, _ := strconv.Atoi(match[1])
	a.handleAppellant(w, r, appeal)
}

func (a appeal) handleAppellant(w http.ResponseWriter, r *http.Request, appealId int) {
	appeal, err := model.GetAppealByID(appealId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Appeal not found", http.StatusNotFound)
	}
	vm := view.NewAppeal(appeal)
	w.Header().Add("Content-Type", "text/html")
	a.appealTemplate.Execute(w, vm)
}
