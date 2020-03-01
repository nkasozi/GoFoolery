package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	w := httptest.NewRecorder()

	r, err := setupRouter()

	if err != nil {
		t.Fail()
		return
	}

	body := "{\"ID\":2,\"Name\":\"testUser\"}"

	req, _ := http.NewRequest("POST", "/user", strings.NewReader(body))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
		return
	}

	p, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Fail()
		return
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, p)
	return
}
