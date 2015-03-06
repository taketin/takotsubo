package gh

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestPrepare(t *testing.T) {
	// override global var
	exit = (func (i int) { })

	Token = "hoge"
	if prepare() {
		t.Errorf("got true want to false")
	}

	Host = "fuga"
	if prepare() {
		t.Errorf("got true want to false")
	}

	ApiVersion = "piyo"
	if !prepare() {
		t.Errorf("got false want to true")
	}
}

var handler = http.HandlerFunc(func (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "HTTP test")
})

func TestPerformRequest(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	result := performRequest(req)

	fmt.Print(result, err)
	if string(result) != "HTTP test" {
		t.Errorf("%s", result)
	}
}
