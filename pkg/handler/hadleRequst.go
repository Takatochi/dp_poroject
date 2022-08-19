package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"project/pkg/logger"
	"project/pkg/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.Engine
	store  store.Store
}
type Index struct {
	Handler *Handler
}

func NewHandler(store store.Store) *Handler {
	return &Handler{
		router: gin.New(),
		store:  store,
	}

}
func (h *Handler) Routing() *gin.Engine {

	return h.router
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Index) Index(ctx *gin.Context) {

	ctx.Request.ParseForm()
	result := []map[string]any{}

	if err := add(&result); err != nil {
		logger.Debug(err)
	}
	get := ctx.Request.Form
	ctx.HTML(http.StatusOK, "index", gin.H{
		"Post": result,
		"Rget": get,
	})

}
func (h *Index) Api(ctx *gin.Context) {
	store, _ := h.Handler.store.Token().Find(1)
	ctx.JSON(http.StatusOK, store)
}
func add(result *[]map[string]any) error {

	// params := url.Values{}
	// params.Add("X-Token", `uz-gWor_utU_sajBMtbsloKL2DmlxkOElo6eKKy_LhgA`)
	// body := strings.NewReader(params.Encode())

	// req, err := http.NewRequest("GET", "", body)
	// if err != nil {
	// 	// handle err
	// if err := Run(g, routers); err != nil {
	// 	log.Fatal(err)
	// }
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	// handle err
	// }

	// defer resp.Body.Close()

	// log.Println(&resp.Body)

	req, err := http.NewRequest("GET", "https://api.monobank.ua/personal/statement/huGbpnagwBu09tUnio8zXA/1659349687", nil)
	if err != nil {
		return err
	}

	req.Header.Add("X-Token", "uz-gWor_utU_sajBMtbsloKL2DmlxkOElo6eKKy_LhgA")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
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
