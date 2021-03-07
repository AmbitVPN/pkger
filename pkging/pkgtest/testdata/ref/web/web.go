package web

import (
	"net/http"

	"github.com/ambitvpn/pkger"
)

func Serve() {
	pkger.Stat("github.com/ambitvpn/pkger:/README.md")
	dir := http.FileServer(pkger.Dir("/public"))
	http.ListenAndServe(":3000", dir)
}
