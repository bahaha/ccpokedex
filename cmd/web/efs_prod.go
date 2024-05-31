//go:build !dev
// +build !dev

package web

import (
	"embed"
	"net/http"
)

//go:embed "assets"
var Files embed.FS

func Public() http.Handler {
	return http.FileServer(http.FS(Files))
}
