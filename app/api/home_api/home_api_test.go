package home_api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func serve(t *testing.T, path string, handler gin.HandlerFunc) *httptest.ResponseRecorder {
	t.Helper()

	r := gin.New()
	r.GET(path, handler)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("GET %s: expected status 200, got %d", path, w.Code)
	}
	if ct := w.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
		t.Fatalf("GET %s: expected a text/html content type, got %q", path, ct)
	}
	return w
}

func assertBodyContains(t *testing.T, w *httptest.ResponseRecorder, wants ...string) {
	t.Helper()
	body := w.Body.String()
	for _, want := range wants {
		if !strings.Contains(body, want) {
			t.Fatalf("expected body to contain %q", want)
		}
	}
}

func TestIndexActionRendersHomePage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := serve(t, "/", IndexAction)
	assertBodyContains(t, w,
		"<title>Emo — 语法简洁 · 显式定义 · 符合直觉的国产系统级编程语言</title>",
		"今天你 <span class=\"emo\">Emo</span> 了没?",
		"一门语言，五个目标",
		"Actor 并发",
		"EmoUI",
		"EmoOS",
	)
}

func TestFeaturesActionRendersFeaturesPage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := serve(t, "/features", FeaturesAction)
	assertBodyContains(t, w,
		"<title>语言特性 — Emo</title>",
		"目录树即模块树",
		"渐进式",
	)
}

func TestTourActionRendersTourPage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := serve(t, "/tour", TourAction)
	assertBodyContains(t, w,
		"<title>语言漫游 — Emo</title>",
		"绑定与字符串",
		"进程：一条消息链",
	)
}

func TestQuickstartActionRendersQuickstartPage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := serve(t, "/quickstart", QuickstartAction)
	assertBodyContains(t, w,
		"<title>快速上手 — Emo</title>",
		"emo run hello.emo",
		"emo build hello.emo -o hello",
	)
}
