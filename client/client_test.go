package client

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
	 "gotest.tools/v3/assert"
)

func TestGetStatusCode(t *testing.T) {

	url := "https://github.com"
	code , err := GetStatusCode(url)
	assert.Equal(t , code, 200)
	assert.NilError(t, err)

	url = "https://github.con"
	code , err = GetStatusCode(url)
	assert.Equal(t , code, 0)
	assert.ErrorContains(t, err, "no such host")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()
	code , err = GetStatusCode(ts.URL)
	assert.Equal(t , code, 200)
	assert.NilError(t, err)
}