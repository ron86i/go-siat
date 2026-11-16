package turismohospedaje_test

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/turismohospedaje"
	"github.com/ron86i/go-siat/v3/protocols/facturacion"
	"github.com/stretchr/testify/require"
)

var _ facturacion.Documento = (*turismohospedaje.Factura)(nil)

func fixture() (turismohospedaje.Cabecera, []turismohospedaje.Detalle) {
	fecha := time.Date(2026, 10, 31, 10, 20, 30, 123000000, time.FixedZone("America/La_Paz", -4*60*60))
	return turismohospedaje.Cabecera{
			NitEmisor: 123456789, RazonSocialEmisor: "Turismo & Hospedaje", Municipio: "La Paz",
			Telefono: new("2222333"), NumeroFactura: 1, Cuf: "ABC123", Cufd: "CUFD-TEST", CodigoSucursal: 1,
			Direccion: "Av. Principal 123", CodigoPuntoVenta: new(2), FechaEmision: fecha,
			NombreRazonSocial: new("Huésped <uno>"), CodigoTipoDocumentoIdentidad: 1,
			NumeroDocumento: "123456", Complemento: "A", CodigoCliente: "CLI-01",
			RazonSocialOperadorTurismo: new("Operador"), CantidadHuespedes: new(2), CantidadHabitaciones: new(1),
			CantidadMayores: new(1), CantidadMenores: new(1), FechaIngresoHospedaje: new(fecha.Add(-time.Hour)),
			CodigoMetodoPago: 1, NumeroTarjeta: new(int64(1234567812345678)),
			SubTotalNeto: 123.456, TotalPrecioNeto: 123.456, Iva: 0, MontoAPagar: 123.456, CodigoInterno: "TUR-01",
			CodigoMoneda: 1, TipoCambio: 1, MontoTotalMoneda: 123.456,
			MontoGiftCard: new(1.25), DescuentoAdicional: new(0.5), CodigoExcepcion: new(1),
			Leyenda: "Ley 453", Usuario: "operador",
		}, []turismohospedaje.Detalle{{
			ActividadEconomica: "001234", CodigoProductoSin: 63111, CodigoProducto: "HAB-01",
			Descripcion: "Hospedaje & desayuno", CodigoTipoHabitacion: new(1), Cantidad: 1.234567,
			UnidadMedida: 58, PrecioNetoUnitario: 123.456789, MontoDescuento: new(0.123456),
			SubTotal: 123.456789, DetalleHuespedes: new("Huésped uno"),
		}}
}

func TestNewAndXML(t *testing.T) {
	cab, detalles := fixture()
	f, err := turismohospedaje.New(cab, detalles)
	require.NoError(t, err)
	require.Equal(t, 6, f.CodigoDocumentoSector())
	require.Equal(t, 3, f.Protocolo())
	for _, modalidad := range []int{1, 2} {
		data, err := f.XML(modalidad)
		require.NoError(t, err)
		root := "facturaElectronicaServicioTuristicoHospedaje"
		if modalidad == 2 {
			root = "facturaComputarizadaServicioTuristicoHospedaje"
		}
		require.True(t, bytes.HasPrefix(data, []byte("<"+root+" ")))
		require.Contains(t, string(data), `xsi:noNamespaceSchemaLocation="`+root+`.xsd"`)
		require.Contains(t, string(data), "<fechaIngresoHospedaje>2026-10-31T09:20:30.123</fechaIngresoHospedaje>")
		require.Contains(t, string(data), "<actividadEconomica>001234</actividadEconomica>")
		require.Contains(t, string(data), "<descripcion>Hospedaje &amp; desayuno</descripcion>")
		require.Contains(t, string(data), "<subTotalNeto>123.46</subTotalNeto>")
		require.Contains(t, string(data), "<precioNetoUnitario>123.46</precioNetoUnitario>")
		require.Contains(t, string(data), "<codigoInterno>TUR-01</codigoInterno>")
		require.NotContains(t, string(data), "<montoTotal>")
		require.NotContains(t, string(data), "<cafc")
	}
	_, err = f.XML(0)
	require.Error(t, err)
	_, err = xml.Marshal(f)
	require.Error(t, err, "la modalidad se debe indicar explícitamente")
}

func clearPointers(v any) {
	r := reflect.ValueOf(v).Elem()
	for i := range r.NumField() {
		field := r.Field(i)
		if field.Kind() == reflect.Pointer && !field.IsNil() {
			field.Elem().SetZero()
		}
	}
}

func TestCopiasDefensivas(t *testing.T) {
	cab, detalles := fixture()
	f, err := turismohospedaje.New(cab, detalles)
	require.NoError(t, err)
	original, err := f.XML(2)
	require.NoError(t, err)
	clearPointers(&cab)
	clearPointers(&detalles[0])
	detalles[0].Descripcion = "modificado"
	copiaCab := f.Cabecera()
	copiaDet := f.Detalles()
	clearPointers(&copiaCab)
	clearPointers(&copiaDet[0])
	*f.CodigoPuntoVenta() = 999
	actual, err := f.XML(2)
	require.NoError(t, err)
	require.Equal(t, original, actual)
}

func TestValidacion(t *testing.T) {
	for _, tc := range []struct {
		nombre    string
		modificar func(*turismohospedaje.Cabecera, *[]turismohospedaje.Detalle)
	}{
		{"sector", func(c *turismohospedaje.Cabecera, _ *[]turismohospedaje.Detalle) { c.CodigoDocumentoSector = 1 }},
		{"nit", func(c *turismohospedaje.Cabecera, _ *[]turismohospedaje.Detalle) { c.NitEmisor = 0 }},
		{"punto", func(c *turismohospedaje.Cabecera, _ *[]turismohospedaje.Detalle) { c.CodigoPuntoVenta = new(-1) }},
		{"fecha", func(c *turismohospedaje.Cabecera, _ *[]turismohospedaje.Detalle) { c.FechaEmision = time.Time{} }},
		{"fecha ingreso", func(c *turismohospedaje.Cabecera, _ *[]turismohospedaje.Detalle) {
			c.FechaIngresoHospedaje = new(time.Time{})
		}},
		{"huespedes", func(c *turismohospedaje.Cabecera, _ *[]turismohospedaje.Detalle) { c.CantidadHuespedes = new(-1) }},
		{"detalle", func(_ *turismohospedaje.Cabecera, d *[]turismohospedaje.Detalle) { *d = nil }},
		{"cantidad", func(_ *turismohospedaje.Cabecera, d *[]turismohospedaje.Detalle) { (*d)[0].Cantidad = 0 }},
	} {
		t.Run(tc.nombre, func(t *testing.T) {
			cab, detalles := fixture()
			tc.modificar(&cab, &detalles)
			_, err := turismohospedaje.New(cab, detalles)
			require.Error(t, err)
		})
	}
	_, err := turismohospedaje.New(turismohospedaje.Cabecera{}, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "NIT")
	require.Contains(t, err.Error(), "detalle")
}

func TestDecimalesNoFinitos(t *testing.T) {
	for _, valor := range []float64{math.NaN(), math.Inf(1), math.Inf(-1)} {
		cab, det := fixture()
		for _, target := range []any{&cab, &det[0]} {
			fields := reflect.ValueOf(target).Elem()
			for i := range fields.NumField() {
				field := fields.Field(i)
				if field.Kind() != reflect.Float64 && (field.Kind() != reflect.Pointer || field.Type().Elem().Kind() != reflect.Float64) {
					continue
				}
				name := fields.Type().Field(i).Name
				t.Run(fmt.Sprintf("%T/%s/%v", target, name, valor), func(t *testing.T) {
					old := reflect.New(field.Type()).Elem()
					old.Set(field)
					if field.Kind() == reflect.Float64 {
						field.SetFloat(valor)
					} else {
						field.Set(reflect.ValueOf(new(valor)))
					}
					_, err := turismohospedaje.New(cab, det)
					field.Set(old)
					require.Error(t, err)
				})
			}
		}
	}
}

func TestPrepararFactura(t *testing.T) {
	cab, detalles := fixture()
	f, err := turismohospedaje.New(cab, detalles)
	require.NoError(t, err)
	req, err := (facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2, TipoFacturaDocumento: 2}).PrepararFactura(f, nil)
	require.NoError(t, err)
	require.Equal(t, 6, req.CodigoDocumentoSector)
	require.Equal(t, 2, req.TipoFacturaDocumento)
	require.Equal(t, cab.Cufd, req.Cufd)
	require.Equal(t, cab.CodigoSucursal, req.CodigoSucursal)
	require.Equal(t, cab.CodigoPuntoVenta, req.CodigoPuntoVenta)
	require.Equal(t, fmt.Sprintf("%x", sha256.Sum256(req.Archivo)), req.HashArchivo)
	gz, err := gzip.NewReader(bytes.NewReader(req.Archivo))
	require.NoError(t, err)
	data, err := io.ReadAll(gz)
	require.NoError(t, err)
	require.NoError(t, gz.Close())
	expected, err := f.XML(2)
	require.NoError(t, err)
	require.Equal(t, expected, data)
	_, err = (facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2}}).PrepararPaquete(t.Context(), []facturacion.Documento{f}, nil)
	require.NoError(t, err)
}

func TestNilFactura(t *testing.T) {
	var f *turismohospedaje.Factura
	require.Error(t, f.Validate())
	_, err := f.XML(2)
	require.Error(t, err)
	require.Zero(t, f.Protocolo())
	require.Zero(t, f.CodigoDocumentoSector())
	require.Empty(t, f.Cufd())
	require.Nil(t, f.CodigoPuntoVenta())
	require.Empty(t, f.Cabecera())
	require.Nil(t, f.Detalles())
}

func TestOpcionalesNulos(t *testing.T) {
	cab, det := fixture()
	for _, target := range []any{&cab, &det[0]} {
		fields := reflect.ValueOf(target).Elem()
		for i := range fields.NumField() {
			if fields.Field(i).Kind() == reflect.Pointer {
				fields.Field(i).SetZero()
			}
		}
	}
	f, err := turismohospedaje.New(cab, det)
	require.NoError(t, err)
	data, err := f.XML(2)
	require.NoError(t, err)
	require.True(t, strings.Contains(string(data), `<fechaIngresoHospedaje xsi:nil="true"></fechaIngresoHospedaje>`))
}

func TestCompatibilidadSolicitud(t *testing.T) {
	cab, detalles := fixture()
	f, err := turismohospedaje.New(cab, detalles)
	require.NoError(t, err)
	for _, req := range []facturacion.SolicitudRecepcionFactura{
		{CodigoModalidad: 0},
		{CodigoModalidad: 1},
		{CodigoModalidad: 2, CodigoDocumentoSector: 1},
		{CodigoModalidad: 2, Cufd: "OTRO-CUFD"},
		{CodigoModalidad: 2, CodigoSucursal: 99},
		{CodigoModalidad: 2, CodigoPuntoVenta: new(99)},
	} {
		_, err := req.PrepararFactura(f, nil)
		require.Error(t, err)
	}
	req := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{
		CodigoModalidad: 2, TipoFacturaDocumento: 2, CodigoSucursal: cab.CodigoSucursal, CodigoPuntoVenta: cab.CodigoPuntoVenta,
	}}
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	_, err = req.PrepararPaquete(ctx, []facturacion.Documento{f}, nil)
	require.ErrorIs(t, err, context.Canceled)
}

func TestSecuenciaXML(t *testing.T) {
	cab, detalles := fixture()
	f, err := turismohospedaje.New(cab, detalles)
	require.NoError(t, err)
	for _, modalidad := range []int{1, 2} {
		data, err := f.XML(modalidad)
		require.NoError(t, err)
		var root struct {
			Cabecera struct {
				Campos []struct{ XMLName xml.Name } `xml:",any"`
			} `xml:"cabecera"`
			Detalle struct {
				Campos []struct{ XMLName xml.Name } `xml:",any"`
			} `xml:"detalle"`
		}
		require.NoError(t, xml.Unmarshal(data, &root))
		var cabecera, detalle []string
		for _, campo := range root.Cabecera.Campos {
			cabecera = append(cabecera, campo.XMLName.Local)
		}
		for _, campo := range root.Detalle.Campos {
			detalle = append(detalle, campo.XMLName.Local)
		}
		require.Equal(t, strings.Fields("nitEmisor razonSocialEmisor municipio telefono numeroFactura cuf cufd codigoSucursal direccion codigoPuntoVenta fechaEmision nombreRazonSocial codigoTipoDocumentoIdentidad numeroDocumento complemento codigoCliente razonSocialOperadorTurismo cantidadHuespedes cantidadHabitaciones cantidadMayores cantidadMenores fechaIngresoHospedaje codigoMetodoPago numeroTarjeta subTotalNeto totalPrecioNeto iva montoaPagar codigoMoneda tipoCambio montoTotalMoneda montoGiftCard descuentoAdicional codigoExcepcion leyenda usuario codigoInterno codigoDocumentoSector"), cabecera)
		require.Equal(t, strings.Fields("actividadEconomica codigoProductoSin codigoProducto descripcion codigoTipoHabitacion cantidad unidadMedida precioNetoUnitario montoDescuento subTotal detalleHuespedes"), detalle)
	}
}
