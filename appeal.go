package main

import (
	"appeals/app/controller"
	"appeals/app/middleware"
	"appeals/app/model"
	"appeals/app/view"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var config = struct {
	DB struct {
		Host     string
		User     string
		Password string `required:"true"`
		Database string
	}
}{}

// DBConf (Public) -
type DBConf struct {
	Host     string
	User     string
	Password string
	Database string
}

func main() {
	// Main function should only start up everything
	loadConfig()
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Startup(templates)
	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{Next: new(middleware.GzipMiddleware)})
}

func loadConfig() {
	if _, err := os.Stat("./config.yml"); err != nil {
		panic("no configuration file found")
	}
	configor.Load(&config, "config.yml")
}

func connectToDatabase() *gorm.DB {
	dbconf := DBConf{
		Host:     config.DB.Host,
		User:     config.DB.User,
		Password: config.DB.Password,
		Database: config.DB.Database,
	}
	// Step1: Create the Database instance and takes the boolean flag to either populate it with data or not
	db, err := gorm.Open("mssql", fmt.Sprintf("sqlserver://%v:%v@%v:1433?database=%v", dbconf.User, dbconf.Password, dbconf.Host, dbconf.Database))
	if err != nil {
		panic(err.Error())
	}
	model.SetDatabase(db)
	return db
}

// populateTemplates (Private) - Returns a map of templates with filename as
// the key and the content of the template as the value of the map
func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)                            // Creates a basic map
	const basePath = "templates"                                             // The name of the folder for the templates located in the root web server folder
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html")) // Making sure that the layout is loaded first
	// Then loads all the layout dependancies next
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html", basePath+"/_shim.html"))
	// opens the folder with all the content for the website
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error()) // exits webserver imediately when the directory is not found.
	}
	fis, err := dir.Readdir(-1) // Reads all the information contained in the w
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Funcs(template.FuncMap{
			"formatTime": view.FormatTime,
		}).Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}
