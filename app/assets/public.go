package assets

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"io/fs"
	"net/http"
	"path"
	"strings"
	"time"
)

//go:embed all:public
var publicFS embed.FS

// PublicHandler serves the committed static files under app/assets/public
// (logo, images) at /public/. Responses carry an ETag and revalidate.
func PublicHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(path.Clean(r.URL.Path), "/public/")
		if name == "" || name == "." || strings.Contains(name, "..") {
			http.NotFound(w, r)
			return
		}
		data, err := fs.ReadFile(publicFS, "public/"+name)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		sum := sha256.Sum256(data)
		etag := `"` + hex.EncodeToString(sum[:])[:16] + `"`
		w.Header().Set("ETag", etag)
		w.Header().Set("Cache-Control", "no-cache")
		if r.Header.Get("If-None-Match") == etag {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		w.Header().Set("Content-Type", publicContentType(name))
		http.ServeContent(w, r, name, time.Time{}, strings.NewReader(string(data)))
	})
}

func publicContentType(name string) string {
	switch path.Ext(name) {
	case ".png":
		return "image/png"
	case ".svg":
		return "image/svg+xml"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".webp":
		return "image/webp"
	case ".ico":
		return "image/x-icon"
	default:
		return "application/octet-stream"
	}
}
