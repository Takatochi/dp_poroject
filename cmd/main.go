package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"project/pkg/handler"
	"project/pkg/server"
	"project/pkg/store"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	router := gin.New()
	g := handler.InitHandler(router)

	srv := new(server.Server)

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	st := store.New(config)

	if err := st.Open(); err != nil {
		log.Fatal(err)
	}
	if err := Run(g, router); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(config, router); err != nil {
		log.Fatal(err)
	}

}
func Run(hadler *handler.Handler, router *gin.Engine) error {

	router.Static("/static", "./static/")

	router.SetFuncMap(template.FuncMap{
		"whole": handler.Whole,
	})

	router.LoadHTMLGlob("templates/*.html")

	result := []map[string]any{}

	err := add(&result)
	if err != nil {
		return err
	}

	r(&hadler.Index, result, router)
	hadler.Contact.Routing(result, "contact", "/contact/", router)

	return nil
}
func r(g handler.Routined, result any, router *gin.Engine) {
	g.Routing(result, "index", "/", router)
}
func add(result *[]map[string]any) error {

	// params := url.Values{}
	// params.Add("X-Token", `uz-gWor_utU_sajBMtbsloKL2DmlxkOElo6eKKy_LhgA`)
	// body := strings.NewReader(params.Encode())

	// req, err := http.NewRequest("GET", "", body)
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

	req, err := http.NewRequest("GET", "", nil)
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
