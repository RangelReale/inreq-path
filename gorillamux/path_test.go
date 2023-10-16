package gorillamux_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rrgmc/inreq"
	"github.com/rrgmc/inreq-path/gorillamux"
	"github.com/stretchr/testify/require"
)

func TestPathValue(t *testing.T) {
	type ReqParams struct {
		DeviceID string `inreq:"path,name=device_id"`
	}

	dec := inreq.NewDecoder(inreq.WithPathValue(gorillamux.NewPathValue()))

	r := mux.NewRouter()
	r.HandleFunc("/device/{device_id}", func(w http.ResponseWriter, r *http.Request) {
		var reqParams ReqParams
		err := dec.Decode(r, &reqParams)
		require.NoError(t, err)
		require.Equal(t, "12", reqParams.DeviceID)
	}).Name("func1")

	req := httptest.NewRequest("GET", "http://localhost/device/12", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
}
