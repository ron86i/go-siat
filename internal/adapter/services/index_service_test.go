package services

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	siatErrors "github.com/ron86i/go-siat/v3/internal/core/errors"
	"github.com/ron86i/go-siat/v3/pkg/models"
)

func TestBuildRequestPackageUsesSFDContract(t *testing.T) {
	req := models.NewRecepcionPaqueteFacturaBuilder().Build()
	xmlBody, err := buildRequest(req)
	if err != nil {
		t.Fatalf("buildRequest() error = %v", err)
	}

	serialized := string(xmlBody)
	if !strings.Contains(serialized, `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"`) {
		t.Fatal("el sobre SOAP debe declarar xmlns:xsi")
	}
	if strings.Contains(serialized, `<cafc`) {
		t.Fatal("el contrato /sfd/ no admite cafc en el paquete")
	}
}

func TestParseSoapResponseReportsHTTPStatusForHTML(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusServiceUnavailable,
		Status:     "503 Service Temporarily Unavailable",
		Body:       io.NopCloser(strings.NewReader("<html><h1>503 Service Temporarily Unavailable</h1></html>")),
	}
	_, err := parseSoapResponse[struct{}](resp)
	var siatErr *siatErrors.SiatError
	if !errors.As(err, &siatErr) {
		t.Fatalf("se esperaba un error HTTP del SIAT, recibido: %v", err)
	}
	if siatErr.StatusCode != http.StatusServiceUnavailable || siatErr.IsRetryable {
		t.Fatalf("error HTTP inesperado: %+v", siatErr)
	}
}
