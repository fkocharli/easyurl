package handlers

import (
	"io/ioutil"
	"net/http"
)

type RouteHandler struct {
	Route   string
	Handler http.Handler
}

type url map[string]string

var urls = map[string]string{"https://example.com/asjakdjsakdj/sjdi9wjweijdiewd234234/4234324": "https://example.com/a/b/c"}

func NewRouteHandler() []RouteHandler {
	return []RouteHandler{
		{
			Route:   "/",
			Handler: http.HandlerFunc(shorten()),
		},
	}
}

func shorten() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			var data []byte
			data, err := ioutil.ReadAll(r.Body)
			r.Body.Read(data)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			v, ok := urls[string(data)]
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Write([]byte(v))
			w.WriteHeader(http.StatusCreated)
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			v, ok := urls[id]
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Add("Location", v)
			w.WriteHeader(http.StatusTemporaryRedirect)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return

		}

	}
}
