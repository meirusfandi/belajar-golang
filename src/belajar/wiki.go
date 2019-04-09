package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var defaultName = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func renderTemplate(response http.ResponseWriter, html string, page *Page) {
	err := templates.ExecuteTemplate(response, html+".html", page)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fun func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		m := defaultName.FindStringSubmatch(request.URL.Path)
		if m == nil {
			http.NotFound(response, request)
			return
		}
		fun(response, request, m[2])
	}
}

func viewHandler(response http.ResponseWriter, request *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(response, request, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(response, "view", p)
}

func editHandler(response http.ResponseWriter, request *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(response, "edit", p)
}

func saveHandler(response http.ResponseWriter, request *http.Request, title string) {
	body := request.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, "/view/"+title, http.StatusFound)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
