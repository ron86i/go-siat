# SIAT v3 Declarative Documents

The SIAT 3.0 document API exposes **strongly typed public structs**. Each sector has its dedicated package under `documents/&lt;sector&gt;`.

The shared facturation pipeline currently supports Compra-Venta, Alquileres,
Servicios Básicos, Tasa Cero, and Turismo/Hospedaje. These packages provide
`New(header, items)`, defensive copies, and `XML(modality)`, and implement
`facturacion.Documento`. Direct `xml.Marshal(invoice)` returns an error because
serialization requires an explicit modality.

Other packages in the inventory expose their own public models and XML behavior;
they do not necessarily implement the shared preparation pipeline or finite-number
validation. Package presence is not a claim of official XSD conformance or SIN
acceptance. Check the selected package's API before choosing a submission method.

## Sector Inventory

| Sector | Package | Main Struct | Modalities |
| :--- | :--- | :--- | :--- |
| 1 | `documents/compraventa` | `Factura` | Electronic and computerized |
| 2 | `documents/alquileres` | `Factura` | Electronic and computerized |
| 3 | `documents/comercialexportacion` | `Factura` | Electronic and computerized |
| 4 | `documents/libreconsignacion` | `Factura` | Electronic and computerized |
| 5 | `documents/zonafranca` | `Factura` | Electronic and computerized |
| 6 | `documents/turismohospedaje` | `Factura` | Electronic and computerized |
| 7 | `documents/seguridadalimentaria` | `Factura` | Electronic and computerized |
| 8 | `documents/tasacero` | `Factura` | Electronic and computerized |
| 9 | `documents/monedaextranjera` | `Factura` | Electronic and computerized |
| 10 | `documents/duttyfree` | `Factura` | Electronic and computerized |
| 11 | `documents/sectoreducativo` | `Factura` | Electronic and computerized |
| 13 | `documents/serviciosbasicos` | `Factura` | Electronic and computerized |
| 14 | `documents/alcanzadaice` | `Factura` | Electronic and computerized |
| 15 | `documents/entidadesfinancieras` | `Factura` | Electronic and computerized |
| 16 | `documents/hoteles` | `Factura` | Electronic and computerized |
| 17 | `documents/hospitales` | `Factura` | Electronic and computerized |
| 18 | `documents/juegosazar` | `Factura` | Electronic and computerized |
| 19 | `documents/hidrocarburosiehd` | `Factura` | Electronic and computerized |
| 20 | `documents/exportacionminerales` | `Factura` | Electronic and computerized |
| 21 | `documents/ventaminerales` | `Factura` | Electronic and computerized |
| 22 | `documents/telecomunicaciones` | `Factura` | Electronic and computerized |
| 23 | `documents/prevalorada` | `Factura` | Electronic and computerized |
| 24 | `documents/notacreditodebito` | `NotaCreditoDebito` | Electronic and computerized |
| 28 | `documents/exportacionservicios` | `Factura` | Electronic and computerized |
| 29 | `documents/notaconciliacion` | `NotaCreditoDebito` | Electronic and computerized |
| 30 | `documents/boletoaereo` | `Factura` | Electronic and computerized |
| 31 | `documents/suministroenergia` | `Factura` | Electronic and computerized |
| 34 | `documents/seguros` | `Factura` | Electronic and computerized |
| 35 | `documents/compraventabonificaciones` | `Factura` | Electronic and computerized |
| 36 | `documents/prevaloradasd` | `Factura` | Electronic and computerized |
| 37 | `documents/gnv` | `Factura` | Electronic and computerized |
| 38 | `documents/hidrocarburosnoiehd` | `Factura` | Electronic and computerized |
| 39 | `documents/gnglsp` | `Factura` | Electronic and computerized |
| 40 | `documents/serviciosbasicoszf` | `Factura` | Electronic and computerized |
| 41 | `documents/vuelocharter` | `Factura` | Electronic and computerized |
| 42 | `documents/alquilerzf` | `Factura` | Electronic and computerized |
| 43 | `documents/exportacionhidrocarburos` | `Factura` | Electronic and computerized |
| 44 | `documents/lubricantes` | `Factura` | Electronic and computerized |
| 45 | `documents/exportacionprecioventa` | `Factura` | Electronic and computerized |
| 46 | `documents/sectoreducativozf` | `Factura` | Electronic and computerized |
| 48 | `documents/notacreditodebitoice` | `NotaCreditoDebito` | Electronic and computerized |
| 49 | `documents/telecomunicacioneszf` | `Factura` | Electronic and computerized |
| 50 | `documents/hospitaleszf` | `Factura` | Electronic and computerized |
| 52 | `documents/ventamineralbcb` | `Factura` | Electronic only |
| 53 | `documents/lubricantesiehd` | `Factura` | Electronic and computerized |
| 54 | `documents/biodiesel` | `Factura` | Electronic and computerized |
| 55 | `documents/combustiblesinsubvencion` | `Factura` | Electronic and computerized |
