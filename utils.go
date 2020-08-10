package main

import (
	"log"
	"net/http"
	"strings"
)

// CheckError checks error from request
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CheckCode confirms whether the request has succeeded
func CheckCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("status code err: %d %s", res.StatusCode, res.Status)
	}
}

// CleanString removes unnecessary spaces from string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
