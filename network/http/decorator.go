package http

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type ClientFunc func(*http.Request) (*http.Response, error)

func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

type Decorator func(Client) Client

func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})
	}
}

type Request[T any] struct {
	body   T
	method string
	uri    string
	url    string
}

func (r Request[T]) SetBody(b T) Request[T] {
	r.body = b
	return r
}

func (r Request[T]) SetMethod(m string) Request[T] {
	r.method = m
	return r
}
func (r Request[T]) SetPath(path string) Request[T] {
	r.uri = path
	return r
}
func (r Request[T]) SetHost(host string) Request[T] {
	r.url = host
	return r
}

func (r Request[T]) Build(t string) (*http.Request, error) {
	switch t {
	case "json":
		return r.toJson()
	case "form":
	}
	return nil, errors.New("build error")
}

func (r Request[T]) toJson() (*http.Request, error) {
	data, err := json.Marshal(r.body)
	if err != nil {
		return nil, errors.Wrap(err, "marshal error")
	}
	return http.NewRequest(r.method, r.url, bytes.NewBuffer(data))
}

// func (r Request[T]) toForm() (*http.Request, error) {

// 	return http.NewRequest(r.method, r.url, bytes.NewBufferString(body.Encode()))
// }
