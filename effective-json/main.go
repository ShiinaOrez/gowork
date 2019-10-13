package main

import (
	"encoding/json"
	"fmt"
	jsongen "github.com/darjun/json-gen"
	"net/http"
	"time"
)

type JSONify interface {
	Map() *jsongen.Map
}

type Response struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    JSONify `json:"data"`
}

func (response Response) Map() *jsongen.Map {
	m := jsongen.NewMap()                  // m type *jsongen.Map
	m.PutInt("code", int64(response.Code)) // response.Code's type decide by CPU(runtime.GOARCH).
	m.PutString("message", response.Message)
	m.PutMap("data", response.Data.Map())
	return m
}

type ResponseData struct {
	Name    string   `json:"name"`
	Score   float64  `json:"score"`
	Friends []string `json:"friends"`
}

func (data ResponseData) Map() *jsongen.Map {
	m := jsongen.NewMap()
	m.PutString("name", data.Name)
	m.PutFloat("score", data.Score)
	m.PutStringArray("friends", data.Friends)
	return m
}

func (response Response) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

func (response Response) Render(w http.ResponseWriter) error {
	writeContentType(w, jsonContentType)

	m := response.Map()
	w.Write(m.Serialize(nil))

	return nil
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = jsonContentType
	}
	return
}

var (
	_ JSONify = Response{}
	_ JSONify = ResponseData{}

	jsonContentType = []string{"application/json; charset=utf-8"}
)

func main() {
	responseData := ResponseData{
		Name:    "Bob",
		Score:   1.2223,
		Friends: []string{"Alice", "Cherry"},
	}
	response := Response{
		Code:    0,
		Message: "OK",
		Data:    responseData,
	}

	time1 := time.Now()
	m := response.Map()
	bytes := m.Serialize(nil)
	fmt.Println(time.Now().Sub(time1))

	time2 := time.Now()
	bytes, _ = json.Marshal(response)
	fmt.Println(time.Now().Sub(time2))
	fmt.Println(string(bytes))
}
