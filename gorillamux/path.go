package gorillamux

import (
	"net/http"

	"github.com/gorilla/mux"
)

type PathValue struct {
}

func (p PathValue) GetRequestPath(r *http.Request, name string) (bool, any, error) {
	vars := mux.Vars(r)
	if v, found := vars[name]; found {
		return true, v, nil
	}
	return false, nil, nil
}

func NewPathValue() *PathValue {
	return &PathValue{}
}
