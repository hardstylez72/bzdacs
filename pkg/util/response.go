package util

import (
	"encoding/json"
	"net/http"
)

type jsonRes struct {
	w                http.ResponseWriter
	err              error
	errorDescription string
	body             interface{}
	statusCode       int
}

func NewResp(w http.ResponseWriter) *jsonRes {
	return &jsonRes{
		w: w,
	}
}

func (r *jsonRes) Error(err error) *jsonRes {
	r.err = err
	return r
}

func (r *jsonRes) Status(statusCode int) *jsonRes {
	r.statusCode = statusCode
	return r
}

func (r *jsonRes) Json(body interface{}) *jsonRes {
	r.body = body
	return r
}

func (r *jsonRes) Send() {
	if r.statusCode != 0 {
		r.w.WriteHeader(r.statusCode)
	}

	r.w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.body != nil {
		body, err := json.Marshal(r.body)
		if err != nil {
			r.w.WriteHeader(http.StatusInternalServerError)
			_, _ = r.w.Write([]byte(err.Error()))
			return
		}
		_, err = r.w.Write(body)
		if err != nil {
			r.w.WriteHeader(http.StatusInternalServerError)
			_, _ = r.w.Write([]byte(err.Error()))
			return
		}
	}
}
