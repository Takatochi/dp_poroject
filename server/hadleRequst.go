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
	post       any
	reader     http.Request
	Get        any
	MUX        http.ServeMux
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

func (s *Server) RequestTemplate(post any, maineroot string, handlename string, tmp ...string) {

	s.post = post
	s.handlename = handlename
	s.maineroot = maineroot
	s.tmp = tmp

	s.MUX.HandleFunc(s.handlename, s.index)

}
func (s Server) index(w http.ResponseWriter, r *http.Request) {
	s.reader = *r
	// id := r.FormValue("id")

	s.GET()
	t, err := template.ParseFiles(s.tmp...)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
	t.ExecuteTemplate(w, s.maineroot, s.post)

}
func (s Server) GET() {

	if s.reader.Method == "GET" {
		s.reader.ParseForm()
		// они все тут
		params := s.reader.Form

		s.Get = params
		// fmt.Printf("GET: id=%s\n", s.Get)
	}

}
