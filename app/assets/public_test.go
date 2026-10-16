package assets

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// The landing page fetches its precompiled demo modules from /public/demo,
// so the embedded public tree must serve them with the WebAssembly MIME type.
func TestPublicHandlerServesDemoModules(t *testing.T) {
	handler := PublicHandler()

	get := func(path string) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		return rec
	}

	rec := get("/public/demo/hero.wasm")
	if rec.Code != http.StatusOK {
		t.Fatalf("hero.wasm status %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "application/wasm") {
		t.Errorf("content type %q, want application/wasm", ct)
	}
	if rec.Body.Len() == 0 {
		t.Error("hero.wasm is empty")
	}

	if rec := get("/public/logo.png"); rec.Code != http.StatusOK {
		t.Errorf("logo status %d", rec.Code)
	}

	if rec := get("/public/..%2fgo.mod"); rec.Code != http.StatusNotFound {
		t.Errorf("traversal attempt status %d, want 404", rec.Code)
	}
}
