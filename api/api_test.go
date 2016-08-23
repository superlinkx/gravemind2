package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

var defaultOptionsContent = "application/json; charset=utf-8"

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func testHandler(route string, handler func(c *gin.Context), method string, body io.Reader) (w *httptest.ResponseRecorder, r *gin.Engine) {
	w = httptest.NewRecorder()
	r = gin.Default()

	//Should add error detection on default
	switch method {
	case "OPTIONS":
		r.OPTIONS(route, handler)
	case "GET":
		r.GET(route, handler)
	case "POST":
		r.POST(route, handler)
	default:
		r.GET(route, handler)
	}

	req, _ := http.NewRequest(method, route, body)

	r.ServeHTTP(w, req)

	return
}

func testAllowedMethods(methods string, wants []string) bool {
	splitMethods := strings.Split(methods, ", ")

	if len(splitMethods) != len(wants) {
		return false
	}

	for _, method := range wants {
		found := false
		for _, smethod := range splitMethods {
			if method == smethod {
				found = true
			}
		}
		if found == false {
			return false
		}
	}

	return true
}

func TestGravemindOptions(t *testing.T) {
	w, _ := testHandler("/gravemind", GravemindOptions, "OPTIONS", nil)

	allowed := testAllowedMethods(w.Header().Get("Allow"), []string{"OPTIONS"})

	t.Log("Returned: ", w.Body)

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultOptionsContent) {
		t.Fail()
	}
}

func TestV1Options(t *testing.T) {
	w, _ := testHandler("/gravemind/v1", V1Options, "OPTIONS", nil)

	allowed := testAllowedMethods(w.Header().Get("Allow"), []string{"OPTIONS"})

	t.Log("Returned: ", w.Body)

	if (w.Code != http.StatusOK) || !allowed || (w.Header().Get("Content-Type") != defaultOptionsContent) {
		t.Fail()
	}
}
