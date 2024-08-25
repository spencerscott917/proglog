package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Log *Log
}

type ProduceRequest struct {
	Record Record
}

type ProduceResponse struct {
	Offset uint64
}

type ConsumeRequest struct {
	Offset uint64
}

type ConsumeResponse struct {
	Record Record
}

func newHttpServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

func NewHttpServer(addr string) *http.Server {
	httpServer := *newHttpServer()
	router := mux.NewRouter()
	router.HandleFunc("/", httpServer.handleProduce).Methods("POST")
	router.HandleFunc("/", httpServer.handleConsume).Methods("GET")
	return &http.Server{
		Addr: addr,
		Handler: router,
	}
}

// Unmarshal json into struct, append body to log, write response
func (srv *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	var req ProduceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	offset, err := srv.Log.Append(req.Record)
	if err := nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := ProduceResponse{Offset: offset}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (srv *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
	return
}

