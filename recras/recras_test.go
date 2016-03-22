package recras

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestIsValidUser_(t *testing.T) {
	var calledURL *url.URL
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calledURL = r.URL
		if user, pass, ok := r.BasicAuth(); ok {
			if user == "recras" && pass == "test" {
				w.WriteHeader(200)
			} else {
				w.WriteHeader(401)
			}
		} else {
			w.WriteHeader(401)
		}
		fmt.Fprintln(w, `{}`)
	}))
	defer ts.Close()

	err := isValidUser(ts.URL, "recras", "test")
	if p := calledURL.Path; p != "/api2/personeel/me" {
		t.Errorf("Expected path to be `/api2/personeel/me`, got %#v", p)
	}
	if err != nil {
		t.Errorf("Expected valid credentials to mark user as invalid, got error %s", err)
	}

	err = isValidUser(ts.URL, "recras", "test2")
	if err != ErrInvalidCredentials {
		t.Errorf("Expected invalid credentials error, got %#v", err)
	}
}
