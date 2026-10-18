package services

import (
	"strings"
	"testing"

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
