package cmd

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/gpabois/cougnat/core/cfg"
	"github.com/gpabois/cougnat/core/common"
	"github.com/gpabois/cougnat/reporting/repositories"
	"github.com/gpabois/cougnat/reporting/services"
	"github.com/gpabois/cougnat/reporting/transports"
	"go.uber.org/dig"
)

func main() {
	// Get the flags
	httpAddr := flag.String("http.addr", ":8080", "HTTP listen address")

	container := dig.New()
	errChannel := make(chan error)

	// Setup the cfg map
	config := cfg.ConfigMap{}
	container.Provide(config.Provide)

	// Setup the logger
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)

	// Configure the deps
	container.Invoke(common.ConfigureCommon)
	container.Invoke(repositories.ConfigureRepositories)
	container.Invoke(services.ConfigureServices)

	// Setup a handler for OS signals
	go func() {
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
		errChannel <- fmt.Errorf("%s", <-osSignals)
	}()

	// Launch the HTTP servers
	go func() {
		// Create a scoped container
		httpScope := container.Scope("HTTP")
		httpScope.Provide(func(log.Logger) {
			log.With(logger, "component", "HTTP")
		})

		// Provide the http handlers.
		container.Provide(transports.ProvideHttpHandler)

		// Start the http server
		container.Invoke(func(h http.Handler) {
			errChannel <- http.ListenAndServe(*httpAddr, h)
		})

	}()

}
