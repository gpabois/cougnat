package tests

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-kit/log"

	"github.com/gpabois/cougnat/core/httputil"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/rand"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/serde"
	"github.com/gpabois/cougnat/reporting/endpoints"
	"github.com/gpabois/cougnat/reporting/models/fixtures"
	"github.com/gpabois/cougnat/reporting/services"
	svc_mocks "github.com/gpabois/cougnat/reporting/services/mocks"
	"github.com/gpabois/cougnat/reporting/transports"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func Test_Report_Success(t *testing.T) {
	container := dig.New()

	// Setup the test
	mockedReportService := svc_mocks.NewIReportService(t)
	container.Provide(func() log.Logger { return log.NewLogfmtLogger(os.Stderr) })
	container.Provide(func() services.IReportService { return mockedReportService })
	container.Provide(func() *svc_mocks.IReportService { return mockedReportService })
	container.Provide(endpoints.ProvideEndpoints)
	container.Provide(transports.ProvideHttpHandler)

	err := container.Invoke(func(h http.Handler, reportSvc *svc_mocks.IReportService) {
		expectedReport := fixtures.RandomAnonymousReport()
		reportID, _ := rand.RandomHex(20)
		expectedReport.ID = option.Some(reportID)
		newReport := fixtures.AsNewReport(expectedReport)

		// Mock the service function
		reportSvc.EXPECT().Report(context.TODO(), newReport).Return(result.Success(expectedReport))

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/reports", bytes.NewBuffer(serde.Marshal(expectedReport, "application/json").Expect()))
		req.Header.Set("content-type", "application/json")
		h.ServeHTTP(rec, req)

		body, _ := io.ReadAll(rec.Result().Body)
		res := serde.UnMarshal[httputil.HttpResult[endpoints.CreateReportResponse]](body, rec.Result().Header.Get("content-type"))

		assert.True(t, res.IsSuccess(), res.UnwrapError())
		assert.Equal(t, expectedReport, res.Expect())
	})
	assert.Nil(t, err, err)
}
