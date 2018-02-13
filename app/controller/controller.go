package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController   home
	noticeController notice
	uploadController upload
	appealController appeal
)

func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	homeController.loginTemplate = templates["login.html"]
	//homeController.searchTemplate = templates["search.html"]
	appealController.appealTemplate = templates["appeal.html"]
	homeController.registerRoutes()
	noticeController.registerRoutes()
	uploadController.registerRoutes()
	appealController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
}
