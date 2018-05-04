package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type errorBody struct {
	Code string
	Msg  string
}

var logger *log.Logger

func createLogger() {
	fp, _ := os.OpenFile("./run.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	logger = log.New(fp, "", log.Ldate|log.Ltime|log.Lmicroseconds)
}

func find(w http.ResponseWriter, req *http.Request) {
	logger.Println(req.URL)
	q := req.URL.Query()
	rec, err := Find(q.Get("mobiles"))
	w.Header().Set("content-type", "application/json")
	if err != nil {
		e := errorBody{"fail", err.Error()}
		b, _ := json.Marshal(e)
		w.Write(b)
	} else {
		b, _ := json.Marshal(rec)
		w.Write(b)
	}

}
func main() {
	createLogger()
	http.HandleFunc("/q", find)
	http.ListenAndServe(":8001", nil)
	select {} //阻塞进程
}
