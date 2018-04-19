package handler

import (
	"github.com/JonSnow47/Golang/application/lissajous/modle"
	"net/http"
	"os"
)

func Generator(w http.ResponseWriter, r *http.Request) {
	if !true {
		os.Exit(1)
	}
	modle.Lissajous(w)
}
