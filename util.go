package main

import (
	"chitchat/data"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/goravel/framework/log/logger"
	"gorm.io/gorm/logger"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Statis       string
}

var config Configuration
var logging *log.Logger

func session(w http.ResponseWriter, r *http.Request) (sess *data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = &data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, f := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", f))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range files {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(template.ParseFiles(files...))

	return
}

// logging
func info(args ...interface{}) {
	logging.SetPrefix("INFO")
	logging.Println(args...)
}

func danger(args ...interface{}) {
	logging.SetPrefix("ERROR")
	logging.Println(args...)
}

func warning(args ...interface{}) {
	logging.SetPrefix("WARNING")
	logging.Println(args...)
}

// version
func version() string {
	return "0.1"
}

func error_message(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}
