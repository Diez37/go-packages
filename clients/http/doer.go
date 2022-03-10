package http

import "net/http"

type Doer interface {
	Do(request *http.Request) (*http.Response, error)
}
