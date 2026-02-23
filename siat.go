package siat

import (
	"github.com/gofiber/fiber/v3/client"
	"github.com/ron86i/go-siat/internal/adapter/service"
)

// sdk es privado (minúscula), el usuario no puede hacer: var s siat.sdk
type sdk struct {
	Operaciones    *service.SiatOperacionesService
	Sincronizacion *service.SiatSincronizacionService
	Codigos        *service.SiatCodigosService
}

// New retorna un puntero al struct privado.
// Go permite esto y el usuario podrá usar los campos públicos.
func New(URL string, httpClient *client.Client) (*sdk, error) {
	operaciones, err := service.NewSiatOperacionesService(URL, httpClient)
	if err != nil {
		return nil, err
	}
	sincronizacion, err := service.NewSiatSincronizacionService(URL, httpClient)
	if err != nil {
		return nil, err
	}
	codigos, err := service.NewSiatCodigosService(URL, httpClient)
	if err != nil {
		return nil, err
	}

	return &sdk{
		Operaciones:    operaciones,
		Sincronizacion: sincronizacion,
		Codigos:        codigos,
	}, nil
}
