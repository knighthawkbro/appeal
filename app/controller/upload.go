package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type upload struct{}

func (u upload) registerRoutes() {
	http.HandleFunc("/upload", u.handleUpload)
}

func (u upload) handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, fmt.Sprint("Cannot handle GET Requests"), http.StatusNotAcceptable)
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	//fmt.Fprintf(w, "%v", handler.Header)
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		log.Println(err)
		return
	}
	cell := xlsx.GetCellValue("Sheet1", "A1")
	log.Println(cell)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
