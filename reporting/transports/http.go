package reporting_transports

import (
	"log"
	"net/http"

	"github.com/go-kit/kit/transport"
	"github.com/gorilla/mux"
	reporting_endpoints "github.com/gpabois/cougnat/reporting/endpoints"
	"github.com/gpabois/goservice/chain"
	endpoint_modules "github.com/gpabois/goservice/endpoint/modules"
	http_transport "github.com/gpabois/goservice/http_transport"
	http_links "github.com/gpabois/goservice/http_transport/links"
	http_modules "github.com/gpabois/goservice/http_transport/modules"
	httputil "github.com/gpabois/gostd/http"
	"github.com/gpabois/gostd/option"
)

func ProvideHttpHandler(logger log.Logger, createReport *reporting_endpoints.CreateReportEndpoint) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").
		Path("/reports").
		Handler(http_transport.NewHandler(
			chain.NewChain().Install(endpoint_modules.NewEndpointModule(createReport)),
			http_modules.HttpModuleArgs{
				DeserializeBody:      option.Some(http_links.ReflectDeserializeBodyArgs{}),
				EnableAuthentication: option.Some(http_modules.AuthenticationArgs{}),
			},
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
