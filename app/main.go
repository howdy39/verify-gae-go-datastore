package app

import (
	"net/http"
	"verify-gae-go-datastore/gae"
)

func init() {
	http.HandleFunc("/", gae.IndexHandler)
}