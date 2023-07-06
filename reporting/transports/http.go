package reporting_transports

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	reporting_endpoints "github.com/gpabois/cougnat/reporting/endpoints"
	scaffolding "github.com/gpabois/goservice/scaffolding"
	mods "github.com/gpabois/goservice/scaffolding/modules"
)

func ProvideHttpHandler(logger log.Logger, createReport *reporting_endpoints.CreateReportEndpoint) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").
		Path("/reports").
		Handler(scaffolding.Http(
			mods.Endpoint(createReport),
			mods.Http(
				mods.DeserializeBody(),
				mods.ExtractAuthenticationStrategy(
					"Requester",
					mods.WitHeaderBasedAuthentication(),
				),
			),
			mods.Authentication("Requester",
				mods.EnableSubject(
					mods.EnableSubjectInjection("owner"),
				),
			),
		))

	return r
}
