package transports

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/log"

	"github.com/go-kit/kit/transport"
	http_transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gpabois/cougnat/core/serde"
	"github.com/gpabois/cougnat/reporting/endpoints"
	httputil "github.com/gpabois/gostd/http"
)

func decodeReportRequest(ctx context.Context, req *http.Request) (request any, err error) {
	contentType := req.Header.Get("Content-Type")
	decodeResult := serde.UnMarshalStream[endpoints.CreateReportRequest](req.Body, contentType)
	return decodeResult.Unwrap()
}

func encodeReportResponse(ctx context.Context, w http.ResponseWriter, response any) error {
	contentType := "application/json"
	encodeResult := serde.MarshalStream(w, response, contentType)
	return encodeResult.UnwrapError()
}

func decodeDeleteReportRequest(ctx context.Context, req *http.Request) (request any, err error) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("missing :id")
	}
	return endpoints.DeleteReportRequest{ReportID: id}, nil
}

func encodeDeleteReportRequest(ctx context.Context, w http.ResponseWriter, response any) error {
	contentType := w.Header.Get("Content-Type")
	httputil.EncodeResult()
	return encodeResult.UnwrapError()
}

type HttpResult[T any] struct {
	value T
	err   string
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func ProvideHttpHandler(e endpoints.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").
		Path("/reports").
		Handler(http_transport.NewServer(
			e.Report,
			decodeReportRequest,
			encodeReportResponse,
			http_transport.ServerErrorEncoder(httputil.EncodeError[endpoints.CreateReportResponse]),
			http_transport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		))

	r.Methods("DELETE").
		Path("/reports/{id}").
		Handler(http_transport.NewServer(
			e.DeleteReport,
			decodeDeleteReportRequest,
			encodeDeleteReportRequest,
			http_transport.ServerErrorEncoder(httputil.EncodeError[endpoints.DeleteReportRequest]),
			http_transport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		))

	return r
}
