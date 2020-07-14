package utils

var Endpoints = map[string]string {
    "symbol"    : "%sstock/symbol?exchange=%s&token=%s",
    "profile"   : "%sstock/profile?symbol=%s&token=%s", 
    "financials": "%sstock/financials?symbol=%s&statement=%s&freq=%s&token=%s",
}

var Statements = []string {"bs", "ic", "cf"}
var Frequency  = []string {"annual", "quarterly", "ttm", "ytd"}

var HttpTooManyRequests = 429