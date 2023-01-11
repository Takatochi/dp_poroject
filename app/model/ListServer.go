package model

type Server struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Port int64  `json:"port"`
}
