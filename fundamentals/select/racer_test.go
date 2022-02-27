package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returnin url of fastest one", func(t *testing.T) {
		slowServer := makeDelayServer(time.Millisecond * 20)
		fastServer := makeDelayServer(time.Millisecond * 0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("did not expect an error but got one: %v", err)
		}
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond", func(t *testing.T) {
		server := makeDelayServer(time.Millisecond * 25)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, time.Millisecond*20)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
