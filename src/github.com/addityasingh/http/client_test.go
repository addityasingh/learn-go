package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClientTimeout(t *testing.T) {
	reqHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d := map[string]interface{}{
			"id":    "12",
			"scope": "test-scope",
		}
		var log bytes.Buffer
		wr := io.MultiWriter(&log, w)

		w.Write([]byte{">>>>>>>lorem ipsum"})
		time.Sleep(100 * time.Millisecond)
		e := json.NewEncoder(wr)
		t.Log(">>>>>>>>>>>>log is ", log)
		err := e.Encode(&d)
		if err != nil {
			t.Error(err)
		}
		w.WriteHeader(http.StatusOK)
	})
	backend := httptest.NewServer(http.TimeoutHandler(reqHandler, 50*time.Millisecond, "server timeout"))

	url := backend.URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("Request error", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("Response error", err)
		return
	}

	defer resp.Body.Close()

	t.Log(">>>>>>>Response is: ", resp)
}
