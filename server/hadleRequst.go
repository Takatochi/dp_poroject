package server

import (
	
	"log"
	"net/http"
	"text/template"
)

type Server struct {
	_handle    []string
	handlename string
	maineroot  string
	tmp        []string
	post     interface{}
}

func Init() *Server {
	return &Server{}

}

func (s *Server) Prefix(_handle ...string) {
	s._handle = _handle

	for i := 0; i < len(s._handle); i++ {
		s.prehandle(s._handle[i])
	}
}

func (s *Server) prehandle(handle string) {
	http.Handle(handle, http.StripPrefix(handle, http.FileServer(http.Dir("."+handle))))
}

func (s *Server) RequestTemplate(post interface{}, maineroot string, handlename string, tmp ...string) {
	
	s.post=post
	s.handlename = handlename
	s.maineroot = maineroot
	s.tmp = tmp

	http.HandleFunc(s.handlename, s.index)

}
func (s Server) index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles(s.tmp...)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
	t.ExecuteTemplate(w, s.maineroot, s.post)

}
