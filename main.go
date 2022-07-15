// Шаблон для роботи с go та http запросами( веб сайти і тд)
package main

import (
	"log"
	"net/http"
	"project/server"
	"project/view"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}
func run() error {

	start := &server.Server{}
	post := view.Post{
		Age:   15,
		Name:  "Arthur",
		Sname: "Havor",
	}
	// візміть за увагу що шлях до стилів то що повний шлях в писаний в лінці як: /static/css/main.css
	start.Prefix("/static/")

	//метод RequestTemplate 4
	start.RequestTemplate(post, "index", "/", "templates/index.html", "templates/header.html", "templates/footer.html")
	start.RequestTemplate(nil, "contact", "/contact/", "templates/contact.html", "templates/header_contact.html", "templates/footer.html")

	if err := http.ListenAndServe(":8088", &start.MUX); err != nil {
		return err
	}

	return nil
}
