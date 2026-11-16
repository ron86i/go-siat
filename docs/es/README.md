# Documentación de go-siat v3

Este módulo se conecta a los servicios SOAP SIAT 3.0 del SIN de Bolivia bajo la ruta /sfd. La documentación describe la API v3 actual, con documentos fiscales y solicitudes tipadas.

## Empezar

- [Tutorial de primera factura](tutorial/primera-factura.md)
- [Enviar una factura](how-to/envio-facturas.md)
- [Firmar facturas electrónicas](how-to/firmar-facturas.md)
- [Enviar lotes y paquetes de contingencia](how-to/envio-lotes.md)
- [Configuración](reference/configuracion.md)
- [Referencia de API](reference/api.md)
- [Paquetes de documentos fiscales](reference/documentos-v3.md)
- [Sectores](explanation/sectores.md)
- [Arquitectura](explanation/arquitectura.md)
- [Ejemplo ejecutable de Compra-Venta](../../examples/compraventa/main.go)

## API actual

- Crear el cliente con siat.NewClient y siat.ClientConfig.
- Construir documentos con los structs públicos de documents/&lt;sector&gt;.
- Invocar los servicios SIAT mediante clientes y solicitudes tipadas en protocols.
- Usar pkg/utils para CUF, XML, firma y compresión cuando se necesite una operación de nivel bajo.

Los clientes preparan y serializan los documentos admitidos antes de enviar solicitudes SOAP. El SIN determina la aceptación fiscal: revisar el estado de transacción, el código de recepción y los mensajes devueltos.

[Read this documentation in English](../en/README.md).
