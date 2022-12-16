package server

import (
	"net/http"

	"github.com/gorilla/mux"
)


 type server struct{
   http *http.Server
 }


func New() *server {
  return &server{} 
}

func (s *server) Start() (err error){
 
  r := mux.NewRouter()


  s.http = &http.Server{
    Handler: r,
  }
  return
}

func (s *server) Stop() (err error) {
  return 
}
