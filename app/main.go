package app

import (
	"net/http"
	"verify-gae-go-datastore/gae"
)

func init() {
	http.HandleFunc("/", gae.IndexHandler)
	http.HandleFunc("/goon", gae.GoonHandler)
	http.HandleFunc("/goonget", gae.GoongetHandler)
	http.HandleFunc("/mem", gae.MemHandler)

	http.HandleFunc("/cheat-save", gae.CheatSaveHandler)
}
