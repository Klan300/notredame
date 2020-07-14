package workers

import (
    "net/http"
)

type Security struct {
    Exchange      string
    Description   string `json:"description,omitempty"`
    DisplaySymbol string `json:"displaySymbol,omitempty"`
    Symbol        string `json:"symbol,omitempty"`
}

type Request struct {
    Exchange  string
    Symbol    string
    Document  string
    Statement string
    Frequency string
    HttpReq   *http.Request
}

type Response struct {
    Data    []byte
    Request *Request
}
