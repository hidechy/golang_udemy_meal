package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"todo/config"
	"todo/models"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {

	http.HandleFunc("/signup", signup)

	http.HandleFunc("/", login)

	http.HandleFunc("/authenticate", authenticate)

	http.HandleFunc("/meals", index)

	http.HandleFunc("/logout", logout)

	http.HandleFunc("/create", create)

	http.HandleFunc("/save", save)

	http.HandleFunc("/edit/", parseURL(edit))

	http.HandleFunc("/update/", parseURL(update))

	http.HandleFunc("/delete/", parseURL(delete))

	return http.ListenAndServe(":"+config.Config.Port, nil)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

//var validPath = regexp.MustCompile("^/(edit|update)/([0-9]+)$")
var validPath = regexp.MustCompile("^/(edit|update|delete)/([0-9]+)$")
