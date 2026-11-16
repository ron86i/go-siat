# Documentos Fiscales SIAT v3 (Declarativos)

La API documental de SIAT 3.0 expone **structs públicos fuertemente tipados**. Cada sector cuenta con su propio paquete dentro de `documents/&lt;sector&gt;`.

El pipeline compartido de facturación admite actualmente Compra-Venta, Alquileres,
Servicios Básicos, Tasa Cero y Turismo/Hospedaje. Estos paquetes ofrecen
`New(cabecera, detalles)`, copias defensivas y `XML(modalidad)`, e implementan
`facturacion.Documento`. `xml.Marshal(factura)` devuelve un error porque la
serialización requiere una modalidad explícita.

Los demás paquetes del inventario exponen sus propios modelos públicos y reglas XML;
no necesariamente implementan la preparación compartida ni validan números finitos.
La presencia de un paquete no certifica conformidad con XSD oficiales ni aceptación
del SIN. Revisa la API del paquete elegido antes de seleccionar el método de envío.

## Inventario de Sectores

| Sector | Paquete | Struct Principal | Modalidades |
| :--- | :--- | :--- | :--- |
| 1 | `documents/compraventa` | `Factura` | Electrónica y computarizada |
| 2 | `documents/alquileres` | `Factura` | Electrónica y computarizada |
| 3 | `documents/comercialexportacion` | `Factura` | Electrónica y computarizada |
| 4 | `documents/libreconsignacion` | `Factura` | Electrónica y computarizada |
| 5 | `documents/zonafranca` | `Factura` | Electrónica y computarizada |
| 6 | `documents/turismohospedaje` | `Factura` | Electrónica y computarizada |
| 7 | `documents/seguridadalimentaria` | `Factura` | Electrónica y computarizada |
| 8 | `documents/tasacero` | `Factura` | Electrónica y computarizada |
| 9 | `documents/monedaextranjera` | `Factura` | Electrónica y computarizada |
| 10 | `documents/duttyfree` | `Factura` | Electrónica y computarizada |
| 11 | `documents/sectoreducativo` | `Factura` | Electrónica y computarizada |
| 13 | `documents/serviciosbasicos` | `Factura` | Electrónica y computarizada |
| 14 | `documents/alcanzadaice` | `Factura` | Electrónica y computarizada |
| 15 | `documents/entidadesfinancieras` | `Factura` | Electrónica y computarizada |
| 16 | `documents/hoteles` | `Factura` | Electrónica y computarizada |
| 17 | `documents/hospitales` | `Factura` | Electrónica y computarizada |
| 18 | `documents/juegosazar` | `Factura` | Electrónica y computarizada |
| 19 | `documents/hidrocarburosiehd` | `Factura` | Electrónica y computarizada |
| 20 | `documents/exportacionminerales` | `Factura` | Electrónica y computarizada |
| 21 | `documents/ventaminerales` | `Factura` | Electrónica y computarizada |
| 22 | `documents/telecomunicaciones` | `Factura` | Electrónica y computarizada |
| 23 | `documents/prevalorada` | `Factura` | Electrónica y computarizada |
| 24 | `documents/notacreditodebito` | `NotaCreditoDebito` | Electrónica y computarizada |
| 28 | `documents/exportacionservicios` | `Factura` | Electrónica y computarizada |
| 29 | `documents/notaconciliacion` | `NotaCreditoDebito` | Electrónica y computarizada |
| 30 | `documents/boletoaereo` | `Factura` | Electrónica y computarizada |
| 31 | `documents/suministroenergia` | `Factura` | Electrónica y computarizada |
| 34 | `documents/seguros` | `Factura` | Electrónica y computarizada |
| 35 | `documents/compraventabonificaciones` | `Factura` | Electrónica y computarizada |
| 36 | `documents/prevaloradasd` | `Factura` | Electrónica y computarizada |
| 37 | `documents/gnv` | `Factura` | Electrónica y computarizada |
| 38 | `documents/hidrocarburosnoiehd` | `Factura` | Electrónica y computarizada |
| 39 | `documents/gnglsp` | `Factura` | Electrónica y computarizada |
| 40 | `documents/serviciosbasicoszf` | `Factura` | Electrónica y computarizada |
| 41 | `documents/vuelocharter` | `Factura` | Electrónica y computarizada |
| 42 | `documents/alquilerzf` | `Factura` | Electrónica y computarizada |
| 43 | `documents/exportacionhidrocarburos` | `Factura` | Electrónica y computarizada |
| 44 | `documents/lubricantes` | `Factura` | Electrónica y computarizada |
| 45 | `documents/exportacionprecioventa` | `Factura` | Electrónica y computarizada |
| 46 | `documents/sectoreducativozf` | `Factura` | Electrónica y computarizada |
| 48 | `documents/notacreditodebitoice` | `NotaCreditoDebito` | Electrónica y computarizada |
| 49 | `documents/telecomunicacioneszf` | `Factura` | Electrónica y computarizada |
| 50 | `documents/hospitaleszf` | `Factura` | Electrónica y computarizada |
| 52 | `documents/ventamineralbcb` | `Factura` | Sólo electrónica |
| 53 | `documents/lubricantesiehd` | `Factura` | Electrónica y computarizada |
| 54 | `documents/biodiesel` | `Factura` | Electrónica y computarizada |
| 55 | `documents/combustiblesinsubvencion` | `Factura` | Electrónica y computarizada |

## Construcción y preparación

```go
// Cabecera y Detalle son structs públicos de entrada. Los importes los entrega
// la aplicación; el SDK no calcula impuestos.
factura, err := compraventa.New(cabecera, detalles)
if err != nil {
    return err
}
xmlData, err := factura.XML(facturacion.ModalidadComputarizada)
if err != nil {
    return err
}
_ = xmlData
```

El [ejemplo compilable de Compra-Venta](../../../examples/compraventa/main.go)
muestra los campos de entrada y la preparación con el cliente de facturación.

## Turismo y Hospedaje (sector 6)

El paquete `turismohospedaje` ofrece `New(cabecera, detalles)` y
`factura.XML(modalidad)`. El modelo conserva los campos opcionales y hace copias
profundas de la fecha de ingreso, los datos de huéspedes y los descuentos. Los
importes siguen el contrato SIAT 3.0 y las fechas incluyen milisegundos.

Configura `TipoFacturaDocumento` en la solicitud con el mismo valor usado al
generar el CUF. La preparación completa con `1` si se omite; Turismo sin derecho
a crédito fiscal requiere `2`.

```go
solicitud := facturacion.SolicitudRecepcionFactura{
    CodigoModalidad:      facturacion.ModalidadComputarizada,
    TipoFacturaDocumento: 2,
    CodigoSucursal:       factura.CodigoSucursal(),
    CodigoPuntoVenta:     factura.CodigoPuntoVenta(),
    Cufd:                 factura.Cufd(),
    Cuis:                 cuis,
}
preparada, err := solicitud.PrepararFactura(factura, nil)
```

Las pruebas locales cubren construcción, errores estructurales, valores no finitos,
copias, XML, precisión, preparación individual y de paquetes. No sustituyen la
validación contra XSD oficiales ni la aceptación del SIN.
