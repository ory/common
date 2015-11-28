package pkg

import (
	"github.com/go-errors/errors"
	"net/http"
)

var (
	ErrNotFound = errors.New("not found.")
)

func HttpErrorHandler(w http.ResponseWriter, err error) {
	if err == ErrNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
