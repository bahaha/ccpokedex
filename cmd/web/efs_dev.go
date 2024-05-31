//go:build dev
// +build dev

package web

import (
	"log/slog"
	"net/http"
)

func Public() http.Handler {
	slog.Info("building public assets for development")
	return http.StripPrefix("/assets/", http.FileServer(http.Dir("./cmd/web/assets")))
}
