package router

import (
	"github.com/JonSnow47/Golang/app/lissajous/handler"
	"net/http"
)

func Router() {
	http.HandleFunc("/lissajous", handler.Generator)
}
