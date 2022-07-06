// Шаблон для роботи с go та http запросами( веб сайти і тд)
package main

import (
	"log"
	"net/http"
	"www/server"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

}
func run() error {

	start := server.Server{}

	// візміть за увагу що шлях до стилів то що повний шлях в писаний в лінці як: /static/css/main.css
	start.Prefix("/static/", "/content/")

	//метод RequestTemplate примаэ 3 параметра попарятку це назва темплйету, роутінг для запита сторінки, прямий шлях до hmtl шаблонів які ви використовуєте в темлейті
	start.RequestTemplate("index", "/", "templates/index.html", "templates/header.html", "templates/footer.html")
	start.RequestTemplate("contact", "/contact/", "templates/contact.html", "templates/header_contact.html", "templates/footer.html")

	if err := http.ListenAndServe(":8088", nil); err != nil {
		return err
	}

	return nil
}
