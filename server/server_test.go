package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Weather
	expectedError error
}

func TestGetWeather(t *testing.T) {

	tests := []Tests{
		{
			name: "basic_request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{ "city": "Denver, CO", "weather": "sunny"}`))
			})),
			response: &Weather{
				City:     "Denver, CO",
				Forecast: "sunny",
			},
			expectedError: nil,
		},
	}
}
