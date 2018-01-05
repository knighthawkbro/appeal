package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController   home
	noticeController notice
	uploadController upload
)

func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	homeController.loginTemplate = templates["login.html"]
	homeController.registerRoutes()
	noticeController.registerRoutes()
	uploadController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
}
