package httpUtil

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

var testString = `{"hello":"world"}`

func TestRequest(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(serveHTTP))
	wg := sync.WaitGroup{}
	wg.Add(10)
	for range [10]struct{}{} {
		go func() {
			code, res, err := Request(http.MethodGet, svr.URL, nil, nil)
			if err != nil {
				t.Errorf("Request test error:%s", err)
			}
			if string(res) != testString || code != http.StatusOK {
				t.Errorf("Request test error.code=%d;res=%s", code, string(res))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestFetch(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(serveHTTP))
	request, err := NewRequest(http.MethodGet, svr.URL, "", nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	var response struct {
		Hello string `json:"hello"`
	}
	err = Fetch(request, &response)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v", response)
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 3)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(testString))
}
