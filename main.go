// Шаблон для роботи с go та http запросами( веб сайти і тд)
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"project/handler"
	"project/server"
	"project/view"
)

type AutoGenerated struct {
	ClientID    string `json:"clientId"`
	Name        string `json:"name"`
	WebHookURL  string `json:"webHookUrl"`
	Permissions string `json:"permissions"`
	Accounts    []struct {
		ID           string   `json:"id"`
		SendID       string   `json:"sendId"`
		Balance      int      `json:"balance"`
		CreditLimit  int      `json:"creditLimit"`
		Type         string   `json:"type"`
		CurrencyCode int      `json:"currencyCode"`
		CashbackType string   `json:"cashbackType"`
		MaskedPan    []string `json:"maskedPan"`
		Iban         string   `json:"iban"`
	} `json:"accounts"`
	Jars []struct {
		ID           string `json:"id"`
		SendID       string `json:"sendId"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		CurrencyCode int    `json:"currencyCode"`
		Balance      int    `json:"balance"`
		Goal         int    `json:"goal"`
	} `json:"jars"`
}

var result AutoGenerated

func main() {
	srv := new(server.Server)
	handlers := handler.InitHandler()
	if err := add(); err != nil {
		log.Fatal(err)
	}
	post := view.Post{
		Age:   15,
		Name:  "Arthur",
		Sname: "Havor",
	}
	handlers.Prefix("/static/")

	handlers.RequestTemplate(post, "index", "/", "templates/index.html", "templates/header.html", "templates/footer.html")
	handlers.RequestTemplate(result, "contact", "/contact/", "templates/contact.html", "templates/header_contact.html", "templates/footer.html")
	log.Println(post.Get)
	if err := srv.Run("8080", &handlers.MUX); err != nil {
		log.Fatal(err)
	}

}
func add() error {

	req, err := http.NewRequest("GET", "https://api.monobank.ua/personal/client-info", nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Token", "uz-gWor_utU_sajBMtbsloKL2DmlxkOElo6eKKy_LhgA")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//Convert the body to type string
	// var result map[string]interface{}

	json.Unmarshal(body, &result)
	fmt.Println(result)

	resp.Request.ParseForm()
	params := resp.Request.Form
	log.Println(params)
	defer resp.Body.Close()

	return nil
}
