package server

import (
	"log"
	"net/http"
	"text/template"
)

type Server struct {
	_handle    []string
	handle     string
	handlename string
	maineroot  string
	tmp        []string
}

func Init() *Server {
	return &Server{}

}
func (s *Server) Prefix(_handle []string) {
	s._handle = _handle
	for i := 0; i < len(s._handle)-1; i++ {
		s.Handle(s._handle[i])
	}
}

func (s *Server) Handle(handle string) {
	s.handle = handle
	http.Handle(s.handle, http.StripPrefix(s.handle, http.FileServer(http.Dir("."+s.handle))))
}

func (s *Server) RequestTemplate(maineroot string, handlename string, tmp []string) {
	s.handlename = handlename
	s.maineroot = maineroot
	s.tmp = tmp
	http.HandleFunc(s.handlename, s.index)

}
/// 
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	//старт темлейтов для метода index (головна сторінка)

	t, err := template.ParseFiles(s.tmp[0])
	for i := 1; i < len(s.tmp); i++ {
		t.ParseFiles(s.tmp[i])
	}
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
	t.ExecuteTemplate(w, s.maineroot, nil)
}
