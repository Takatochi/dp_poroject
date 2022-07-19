package handler

import (
	"log"
	"net/http"
	"text/template"
)

type Handler struct {
	_handle    []string
	handlename string
	maineroot  string
	tmp        []string
	post       any
	MUX        http.ServeMux
}

func InitHandler() *Handler {
	return &Handler{
		MUX: *http.NewServeMux(),
	}

}

func (s *Handler) Prefix(_handle ...string) {
	s._handle = _handle

	for i := 0; i < len(s._handle); i++ {
		s.prehandle(s._handle[i])
	}
}

func (s *Handler) prehandle(handle string) {
	s.MUX.Handle(handle, http.StripPrefix(handle, http.FileServer(http.Dir("."+handle))))
}

func (s *Handler) RequestTemplate(post any, maineroot string, handlename string, tmp ...string) {

	s.post = &post
	s.handlename = handlename
	s.maineroot = maineroot
	s.tmp = tmp
	s.MUX.HandleFunc(s.handlename, s.index)

}

func (s Handler) index(w http.ResponseWriter, r *http.Request) {
	// id := r.FormValue("id")
	s.GET(w, r)
	t, err := template.ParseFiles(s.tmp...)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
	t.ExecuteTemplate(w, s.maineroot, s.post)

}
func (s *Handler) GET(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		r.ParseForm()
		// они все тут
		params := r.Form
		// s.post.Get = params
		log.Println(params)
		// fmt.Fprintln(w, params["id"])
	}

}
