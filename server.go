package main

import (
	"html/template"
	"net/http"
)

func main() {
	Pages := EveryPages()

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/ressources/", http.StripPrefix("/ressources/", fileServer))

	http.HandleFunc("/", Error404)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		HomePage(w, r, Pages)
	})
	http.HandleFunc("/why", WhyPage)

	// log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
	// (
	// openssl genrsa -out server.key 2048
	// openssl ecparam -genkey -name secp384r1 -out server.key
	// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
	// )
	http.ListenAndServe(":8080", nil)
}

func Error404(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("./static/error404/index.html"))
	template.Execute(w, "index.html")
}

func HomePage(w http.ResponseWriter, r *http.Request, Pages *[]HTMLPage) {
	var template = template.Must(template.ParseFiles("./static/homepage/index.html"))
	template.ExecuteTemplate(w, "index.html", Pages)
}

func WhyPage(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("./static/why/index.html"))
	template.Execute(w, "why.html")
}

type HTMLPage struct {
	Image       string
	Name        string
	Href        string
	Languages   []string
	Description string
	Link        []string
	Learned     string
}

func EveryPages() *[]HTMLPage {
	HTMLPage_Home := HTMLPage{
		Image: "/ressources/img/default.png",
		Name:  "Home",
		Href:  "home",
		Languages: []string{
			"HTML/CSS",
			"Golang",
		},
		Description: "Default description",
		Link:        []string{},
		Learned:     "",
	}

	HTMLPage_Why := HTMLPage{
		Image: "/ressources/img/default.png",
		Name:  "Why?",
		Href:  "why",
		Languages: []string{
			"HTML/CSS",
		},
		Description: "Default description",
		Link:        []string{},
		Learned:     "",
	}

	Pages := &[]HTMLPage{
		HTMLPage_Home,
		HTMLPage_Why,
	}

	return Pages
}
