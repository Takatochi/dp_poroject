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

	//провірка директорії
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
}
func run() error {
	tmp := []string{"templates/index.html", "templates/header.html", "templates/footer.html"}
	tmp1 := []string{"templates/contact.html", "templates/header.html", "templates/footer.html"}
	prf := []string{"/static/", "/content/"}
	start := server.Server{}
	start.Prefix(prf)
	// start.Handle("/static/")
	start.RequestTemplate("index", "/", tmp)
	start.RequestTemplate("contact", "/contact", tmp1)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
