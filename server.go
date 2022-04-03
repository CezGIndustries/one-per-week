package main

import (
	"html/template"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/ressources/", http.StripPrefix("/ressources/", fileServer))

	http.HandleFunc("/", Error404)
	http.HandleFunc("/home", HomePage)
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

func HomePage(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("./static/homepage/index.html"))
	template.Execute(w, "index.html")
}

func WhyPage(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("./static/why/index.html"))
	template.Execute(w, "why.html")
}
