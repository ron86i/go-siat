package sincronizacion

import "encoding/xml"

// Mensaje representa un mensaje informativo o de error emitido por el SIN.
type Mensaje struct {
	Codigo      int    `xml:"codigo"`
	Descripcion string `xml:"descripcion"`
}

// Estado contiene los mensajes y el resultado de la transacción.
type Estado struct {
	Mensajes    []Mensaje `xml:"mensajesList"`
	Transaccion *bool     `xml:"transaccion"`
}

// Parametrica representa una entrada de un catálogo clasificador.
type Parametrica struct {
	CodigoClasificador int    `xml:"codigoClasificador"`
	Descripcion        string `xml:"descripcion"`
}

// ResultadoParametricas contiene el listado de códigos de un catálogo clasificador.
type ResultadoParametricas struct {
	XMLName xml.Name `xml:"RespuestaListaParametricas"`
	Estado
	Codigos []Parametrica `xml:"listaCodigos"`
}

// Actividad representa una actividad económica del contribuyente según CAEB.
type Actividad struct {
	CodigoCaeb    string `xml:"codigoCaeb"`
	Descripcion   string `xml:"descripcion"`
	TipoActividad string `xml:"tipoActividad"`
}

// ResultadoActividades contiene las actividades económicas sincronizadas.
type ResultadoActividades struct {
	XMLName xml.Name `xml:"RespuestaListaActividades"`
	Estado
	Actividades []Actividad `xml:"listaActividades"`
}

// ActividadSector relaciona una actividad económica con un tipo de documento sector.
type ActividadSector struct {
	CodigoActividad       string `xml:"codigoActividad"`
	CodigoDocumentoSector int    `xml:"codigoDocumentoSector"`
	TipoDocumentoSector   string `xml:"tipoDocumentoSector"`
}

// ResultadoActividadesSector contiene las relaciones de actividad y documento sector.
type ResultadoActividadesSector struct {
	XMLName xml.Name `xml:"RespuestaListaActividadesDocumentoSector"`
	Estado
	Actividades []ActividadSector `xml:"listaActividadesDocumentoSector"`
}

// Leyenda contiene el texto legal obligatorio asociado a una actividad.
type Leyenda struct {
	CodigoActividad    string `xml:"codigoActividad"`
	DescripcionLeyenda string `xml:"descripcionLeyenda"`
}

// ResultadoLeyendas contiene las leyendas de factura por actividad.
type ResultadoLeyendas struct {
	XMLName xml.Name `xml:"RespuestaListaParametricasLeyendas"`
	Estado
	Leyendas []Leyenda `xml:"listaLeyendas"`
}

// Producto contiene un producto o servicio homologado por el SIN.
type Producto struct {
	CodigoActividad     string `xml:"codigoActividad"`
	CodigoProducto      int64  `xml:"codigoProducto"`
	DescripcionProducto string `xml:"descripcionProducto"`
	Nandina             string `xml:"nandina"`
}

// ResultadoProductos contiene los productos y servicios homologados.
type ResultadoProductos struct {
	XMLName xml.Name `xml:"RespuestaListaProductos"`
	Estado
	Productos []Producto `xml:"listaCodigos"`
}

// ResultadoFechaHora conserva la fecha y hora oficial del servidor del SIN.
type ResultadoFechaHora struct {
	XMLName xml.Name `xml:"RespuestaFechaHora"`
	Estado
	FechaHora string `xml:"fechaHora"`
}

// RespuestaComunicacion conserva el resultado de verificar comunicación con el endpoint.
type RespuestaComunicacion struct {
	Mensajes    []Mensaje `xml:"mensajesList"`
	Transaccion *bool     `xml:"transaccion"`
}
