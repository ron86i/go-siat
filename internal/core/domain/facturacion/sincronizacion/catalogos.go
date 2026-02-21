package sincronizacion

import (
	"encoding/xml"
	"time"
)

// SolicitudSincronizacion representa los parámetros comunes requeridos para las solicitudes de sincronización al SIAT.
type SolicitudSincronizacion struct {
	CodigoAmbiente   int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoPuntoVenta int    `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	CodigoSistema    string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal   int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis             string `xml:"cuis" json:"cuis"`
	NIT              int64  `xml:"nit" json:"nit"`
}

// --- Sincronizar actividades ---

// SincronizarActividades representa la solicitud para obtener el catálogo de actividades económicas del contribuyente.
type SincronizarActividades struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarActividades" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarActividadesResponse representa la respuesta del SIAT con el catálogo de actividades.
type SincronizarActividadesResponse struct {
	XMLName                   xml.Name                  `xml:"sincronizarActividadesResponse" json:"-"`
	XmlnsNs2                  string                    `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaActividades RespuestaListaActividades `xml:"RespuestaListaActividades" json:"respuestaListaActividades"`
}

// RespuestaListaActividades contiene el resultado de la sincronización de actividades.
type RespuestaListaActividades struct {
	Transaccion      bool               `xml:"transaccion" json:"transaccion"`
	ListaActividades []ListaActividades `xml:"listaActividades" json:"listaActividades"`
}

// ListaActividades detalla una actividad económica según el catálogo del SIAT.
type ListaActividades struct {
	CodigoCaeb    string `xml:"codigoCaeb" json:"codigoCaeb"`
	Descripcion   string `xml:"descripcion" json:"descripcion"`
	TipoActividad string `xml:"tipoActividad" json:"tipoActividad"`
}

// --- Sincronizar fecha y hora ---

// SincronizarFechaHora representa la solicitud para obtener la fecha y hora oficial del servidor del SIAT.
type SincronizarFechaHora struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarFechaHora" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarFechaHoraResponse representa la respuesta del SIAT con la fecha y hora oficial.
type SincronizarFechaHoraResponse struct {
	XMLName            xml.Name           `xml:"sincronizarFechaHoraResponse" json:"-"`
	XmlnsNs2           string             `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaFechaHora RespuestaFechaHora `xml:"RespuestaFechaHora" json:"respuestaFechaHora"`
}

// RespuestaFechaHora contiene la fecha y hora sincronizada con el SIAT.
type RespuestaFechaHora struct {
	Transaccion bool      `xml:"transaccion" json:"transaccion"`
	FechaHora   time.Time `xml:"fechaHora" json:"fechaHora"`
}

// --- Sincronizar lista actividades documento sector ---

// SincronizarListaActividadesDocumentoSector representa la solicitud para obtener la relación entre actividades y documentos sector.
type SincronizarListaActividadesDocumentoSector struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarListaActividadesDocumentoSector" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarListaActividadesDocumentoSectorResponse representa la respuesta con la relación actividad-sector.
type SincronizarListaActividadesDocumentoSectorResponse struct {
	XMLName                                  xml.Name                                 `xml:"sincronizarListaActividadesDocumentoSectorResponse" json:"-"`
	XmlnsNs2                                 string                                   `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaActividadesDocumentoSector RespuestaListaActividadesDocumentoSector `xml:"RespuestaListaActividadesDocumentoSector" json:"respuestaListaActividadesDocumentoSector"`
}

// RespuestaListaActividadesDocumentoSector contiene el listado de documentos por actividad.
type RespuestaListaActividadesDocumentoSector struct {
	Transaccion                     bool                              `xml:"transaccion" json:"transaccion"`
	ListaActividadesDocumentoSector []ListaActividadesDocumentoSector `xml:"listaActividadesDocumentoSector" json:"listaActividadesDocumentoSector"`
}

// ListaActividadesDocumentoSector detalla la relación entre una actividad y un tipo de documento sector.
type ListaActividadesDocumentoSector struct {
	CodigoActividad       string `xml:"codigoActividad" json:"codigoActividad"`
	CodigoDocumentoSector uint64 `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
	TipoDocumentoSector   string `xml:"tipoDocumentoSector" json:"tipoDocumentoSector"`
}

// --- Sincronizar lista leyendas factura ---

// SincronizarListaLeyendasFactura representa la solicitud para obtener el catálogo de leyendas para las facturas.
type SincronizarListaLeyendasFactura struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarListaLeyendasFactura" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarListaLeyendasFacturaResponse representa la respuesta con el catálogo de leyendas.
type SincronizarListaLeyendasFacturaResponse struct {
	XMLName                            xml.Name                           `xml:"sincronizarListaLeyendasFacturaResponse" json:"-"`
	XmlnsNs2                           string                             `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricasLeyendas RespuestaListaParametricasLeyendas `xml:"RespuestaListaParametricasLeyendas" json:"respuestaListaParametricasLeyendas"`
}

// RespuestaListaParametricasLeyendas contiene las leyendas obtenidas para las actividades.
type RespuestaListaParametricasLeyendas struct {
	Transaccion   bool             `xml:"transaccion" json:"transaccion"`
	ListaLeyendas *[]ListaLeyendas `xml:"listaLeyendas" json:"listaLeyendas"`
}

// ListaLeyendas detalla una leyenda asociada a una actividad.
type ListaLeyendas struct {
	CodigoActividad    string `xml:"codigoActividad" json:"codigoActividad"`
	DescripcionLeyenda string `xml:"descripcionLeyenda" json:"descripcionLeyenda"`
}

// --- Sincronizar lista productos servicios ---

// SincronizarListaProductosServicios representa la solicitud para obtener el catálogo de productos y servicios.
type SincronizarListaProductosServicios struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarListaProductosServicios" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarListaProductosServiciosResponse representa la respuesta con el catálogo de productos y servicios.
type SincronizarListaProductosServiciosResponse struct {
	XMLName                 xml.Name                `xml:"sincronizarListaProductosServiciosResponse" json:"-"`
	XmlnsNs2                string                  `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaProductos RespuestaListaProductos `xml:"RespuestaListaProductos" json:"respuestaListaProductos"`
}

// RespuestaListaProductos contiene el listado de productos homologados.
type RespuestaListaProductos struct {
	Transaccion  bool                    `xml:"transaccion" json:"transaccion"`
	ListaCodigos []ListaCodigosProductos `xml:"listaCodigos" json:"listaCodigos"`
}

// ListaCodigosProductos detalla un producto asociado a una actividad.
type ListaCodigosProductos struct {
	CodigoActividad     uint64 `xml:"codigoActividad" json:"codigoActividad"`
	CodigoProducto      uint64 `xml:"codigoProducto" json:"codigoProducto"`
	DescripcionProducto string `xml:"descripcionProducto" json:"descripcionProducto"`
}

// --- Estructuras genéricas para paramétricas ---

// RespuestaListaParametricas es una estructura genérica para las respuestas de catálogos paramétricos del SIAT.
type RespuestaListaParametricas struct {
	Transaccion  bool           `xml:"transaccion" json:"transaccion"`
	ListaCodigos []ListaCodigos `xml:"listaCodigos" json:"listaCodigos"`
}

// ListaCodigos representa un elemento básico de un catálogo paramétrico (código y descripción).
type ListaCodigos struct {
	CodigoClasificador uint64 `xml:"codigoClasificador" json:"codigoClasificador"`
	Descripcion        string `xml:"descripcion" json:"descripcion"`
}

// --- Sincronizar lista mensajes servicios ---

// SincronizarListaMensajesServicios representa la solicitud para obtener los mensajes de error e informativos del servicio.
type SincronizarListaMensajesServicios struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarListaMensajesServicios" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarListaMensajesServiciosResponse representa la respuesta con los mensajes del servicio.
type SincronizarListaMensajesServiciosResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarListaMensajesServiciosResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica eventos significativos ---

// SincronizarParametricaEventosSignificativos representa la solicitud para obtener el catálogo de eventos significativos.
type SincronizarParametricaEventosSignificativos struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaEventosSignificativos" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaEventosSignificativosResponse representa la respuesta de eventos significativos.
type SincronizarParametricaEventosSignificativosResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaEventosSignificativosResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica motivo anulación ---

// SincronizarParametricaMotivoAnulacion representa la solicitud para obtener los motivos de anulación de facturas.
type SincronizarParametricaMotivoAnulacion struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaMotivoAnulacion" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaMotivoAnulacionResponse representa la respuesta de motivos de anulación.
type SincronizarParametricaMotivoAnulacionResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaMotivoAnulacionResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica país origen ---

// SincronizarParametricaPaisOrigen representa la solicitud para obtener el catálogo de países de origen.
type SincronizarParametricaPaisOrigen struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaPaisOrigen" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaPaisOrigenResponse representa la respuesta de países de origen.
type SincronizarParametricaPaisOrigenResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaPaisOrigenResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica tipo documento identidad ---

// SincronizarParametricaTipoDocumentoIdentidad representa la solicitud para obtener los tipos de documentos de identidad.
type SincronizarParametricaTipoDocumentoIdentidad struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaTipoDocumentoIdentidad" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaTipoDocumentoIdentidadResponse representa la respuesta de tipos de documentos de identidad.
type SincronizarParametricaTipoDocumentoIdentidadResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaTipoDocumentoIdentidadResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica tipo documento sector ---

// SincronizarParametricaTipoDocumentoSector representa la solicitud para obtener los tipos de documentos sector.
type SincronizarParametricaTipoDocumentoSector struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaTipoDocumentoSector" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaTipoDocumentoSectorResponse representa la respuesta de tipos de documentos sector.
type SincronizarParametricaTipoDocumentoSectorResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaTipoDocumentoSectorResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica tipo emisión ---

// SincronizarParametricaTipoEmision representa la solicitud para obtener los tipos de emisión (Online/Offline).
type SincronizarParametricaTipoEmision struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaTipoEmision" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaTipoEmisionResponse representa la respuesta de tipos de emisión.
type SincronizarParametricaTipoEmisionResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaTipoEmisionResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica tipo habitación ---

// SincronizarParametricaTipoHabitacion representa la solicitud para obtener el catálogo de tipos de habitación (Sector Hoteles).
type SincronizarParametricaTipoHabitacion struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaTipoHabitacion" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaTipoHabitacionResponse representa la respuesta de tipos de habitación.
type SincronizarParametricaTipoHabitacionResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaTipoHabitacionResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica tipo método pago ---

// SincronizarParametricaTipoMetodoPago representa la solicitud para obtener los métodos de pago aceptados.
type SincronizarParametricaTipoMetodoPago struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaTipoMetodoPago" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaTipoMetodoPagoResponse representa la respuesta de tipos de métodos de pago.
type SincronizarParametricaTipoMetodoPagoResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaTipoMetodoPagoResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica tipo moneda ---

// SincronizarParametricaTipoMoneda representa la solicitud para obtener el catálogo de monedas.
type SincronizarParametricaTipoMoneda struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaTipoMoneda" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaTipoMonedaResponse representa la respuesta de tipos de moneda.
type SincronizarParametricaTipoMonedaResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaTipoMonedaResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica tipo punto venta ---

// SincronizarParametricaTipoPuntoVenta representa la solicitud para obtener los tipos de puntos de venta.
type SincronizarParametricaTipoPuntoVenta struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaTipoPuntoVenta" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaTipoPuntoVentaResponse representa la respuesta de tipos de puntos de venta.
type SincronizarParametricaTipoPuntoVentaResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaTipoPuntoVentaResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica tipos factura ---

// SincronizarParametricaTiposFactura representa la solicitud para obtener los tipos de facturas soportados.
type SincronizarParametricaTiposFactura struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaTiposFactura" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaTiposFacturaResponse representa la respuesta de tipos de factura.
type SincronizarParametricaTiposFacturaResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaTiposFacturaResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}

// --- Sincronizar paramétrica unidad medida ---

// SincronizarParametricaUnidadMedida representa la solicitud para obtener el catálogo de unidades de medida.
type SincronizarParametricaUnidadMedida struct {
	XMLName                 xml.Name                `xml:"ns:sincronizarParametricaUnidadMedida" json:"-"`
	SolicitudSincronizacion SolicitudSincronizacion `xml:"SolicitudSincronizacion" json:"solicitudSincronizacion"`
}

// SincronizarParametricaUnidadMedidaResponse representa la respuesta de unidades de medida.
type SincronizarParametricaUnidadMedidaResponse struct {
	XMLName                    xml.Name                   `xml:"sincronizarParametricaUnidadMedidaResponse" json:"-"`
	XmlnsNs2                   string                     `xml:"xmlns:ns2,attr" json:"-"`
	RespuestaListaParametricas RespuestaListaParametricas `xml:"RespuestaListaParametricas" json:"respuestaListaParametricas"`
}
