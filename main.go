// Шаблон для роботи с go та http запросами( веб сайти і тд)
package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"project/handler"
	"project/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	g := handler.InitHandler(router)

	srv := new(server.Server)
	Run(*g, router)

	if err := srv.Run("8088", router); err != nil {
		log.Fatal(err)
	}

}
func Run(hadler handler.Handler, router *gin.Engine) {

	router.Static("/static", "./static/")

	router.SetFuncMap(template.FuncMap{
		"whole": hadler.Index.Whole,
	})

	router.LoadHTMLGlob("templates/*.html")
	result := []map[string]any{}

	err := add(&result)
	if err != nil {
		log.Fatal(err)
	}
	r(&hadler.Index, result, router)
	hadler.Contact.Routing(result, "contact", "/contact/", router)
}
func r(g handler.Routined, result any, router *gin.Engine) {
	g.Routing(result, "index", "/", router)
}
func add(result *[]map[string]any) error {

	// params := url.Values{}
	// params.Add("X-Token", `uz-gWor_utU_sajBMtbsloKL2DmlxkOElo6eKKy_LhgA`)
	// body := strings.NewReader(params.Encode())

	// req, err := http.NewRequest("GET", "https://api.monobank.ua/personal/statement/{0}/{1546304461}/{to}", body)
	// if err != nil {
	// 	// handle err
	// }
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	// handle err
	// }

	// defer resp.Body.Close()

	// log.Println(&resp.Body)

	req, err := http.NewRequest("GET", "https://api.monobank.ua/personal/statement/huGbpnagwBu09tUnio8zXA/1658086424", nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Token", "uz-gWor_utU_sajBMtbsloKL2DmlxkOElo6eKKy_LhgA")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//Convert the body to type string

	json.Unmarshal(body, &result)

	// var money float32
	// for _, element := range result {
	// 	if element.Amount < 0 {
	// 		money = float32(element.Amount) / 100
	// 		fmt.Println(element.Description, money, float32(element.Balance)/100, time.Unix(int64(element.Time), 0))
	// 	}
	// }
	defer resp.Body.Close()

	return nil
}
