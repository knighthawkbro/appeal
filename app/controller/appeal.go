package controller

import (
	"appeals/app/model"
	"appeals/app/view"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type appeal struct {
	appealTemplate *template.Template
}

func (a appeal) registerRoutes() {
	http.HandleFunc("/appeal/", a.handleAppeal)
	http.HandleFunc("/search", a.handleSearch)
}

func (a appeal) handleAppeal(w http.ResponseWriter, r *http.Request) {
	_, err := model.Sessions(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	appealPattern, _ := regexp.Compile(`/appeal/(\d+)`)
	match := appealPattern.FindStringSubmatch(r.URL.Path)
	appeal, _ := strconv.Atoi(match[1])
	if r.Method == http.MethodPost {
		r.ParseForm()
		for k, v := range r.Form {
			fmt.Printf("key: %v; value: %v\n", k, strings.Join(v, ""))
		}
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	a.handleAppellant(w, r, appeal)
}

func (a appeal) handleAppellant(w http.ResponseWriter, r *http.Request, appealID int) {
	appeal, err := model.GetAppealByID(appealID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Appeal not found", http.StatusNotFound)
	}
	vm := view.NewAppeal(appeal)
	w.Header().Add("Content-Type", "text/html")
	a.appealTemplate.Execute(w, vm)
}

func (a appeal) handleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	values := r.URL.Query()
	id, err := strconv.Atoi(strings.Join(values["id"], ""))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
	a.handleAppellant(w, r, id)
}
